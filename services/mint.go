package services

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path"
	"time"

	"github.com/nft-rainbow/rainbow-app-service/clients/rainbow"
	"github.com/nft-rainbow/rainbow-app-service/config"
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/models/enums"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type MintItemDto struct {
	MintTo      string `form:"mint_to" json:"mint_to" binding:"required"`
	TokenId     string `form:"token_id" json:"token_id"`
	Amount      *uint  `form:"amount" json:"amount"`
	MetadataUri string `form:"metadata_uri" json:"metadata_uri"`
}

type MintBatchDto struct {
	AppId      uint             `json:"app_id" binding:"required"`
	SourceType enums.SourceType `swaggertype:"string" json:"source_type" binding:"required"`
	rainbow.ContractInfoDtoWithoutType
	MintItems []*MintItemDto `json:"mint_items" binding:"required,dive"`
}

func (m *MintBatchDto) ToCustomMintBatchDto(source2Addr map[string]string) (*rainbow.CustomMintBatchDto, error) {
	var result rainbow.CustomMintBatchDto
	result.ContractInfoDtoWithoutType = m.ContractInfoDtoWithoutType

	for _, item := range m.MintItems {
		v, ok := source2Addr[item.MintTo]
		if !ok {
			return nil, errors.Errorf("missing address of %v", item.MintTo)
		}

		mintItem := rainbow.MintItemDto{
			MintToAddress: v,
			TokenId:       item.TokenId,
			Amount:        item.Amount,
			MetadataUri:   item.MetadataUri,
		}
		result.MintItems = append(result.MintItems, &mintItem)
	}
	return &result, nil
}

type MintService struct {
	cellar Cellar
}

// 1. map source to address
// 2. call rainbow-api
func (s *MintService) MintBatchByMetaUri(userId uint, req *MintBatchDto) (*models.BatchMintTask, error) {
	task := models.BatchMintTask{
		UserId:     userId,
		AppId:      req.AppId,
		SourceType: req.SourceType,
		Status:     enums.BATCH_MINT_STATUS_INIT,
	}
	if req.SourceType != enums.SOURCE_TYPE_ADDRESS {
		reqJ, _ := json.Marshal(req)
		filePath := path.Join(config.GetConfig().Storage.Base, config.GetConfig().Storage.BatchMintRequests, fmt.Sprintf("%d", time.Now().UnixNano()))
		if err := os.WriteFile(filePath, reqJ, fs.ModePerm); err != nil {
			return nil, errors.WithMessage(err, "failed to save request")
		}
		task.RequestFilePath = filePath
	}

	if err := task.Save(); err != nil {
		return nil, err
	}

	go s.runBatchMintTask(&task, req)
	return &task, nil
}

func (s *MintService) runBatchMintTask(task *models.BatchMintTask, req *MintBatchDto) {
	if task.IsFinalized() {
		return
	}

	// read req if req is nil
	if req == nil {
		b, err := ioutil.ReadFile(task.RequestFilePath)
		if err != nil {
			task.SetError(errors.WithMessage(err, "failed to read request from cached file"))
			return
		}
		if err = json.Unmarshal(b, &req); err != nil {
			task.SetError(errors.WithMessage(err, "failed to json unmarshal to request"))
			return
		}
	}

	var sources []string
	for _, item := range req.MintItems {
		sources = append(sources, item.MintTo)
	}

	exists, unexist, err := (&AddressFinder{req.SourceType}).Find(req.Chain, sources)
	if err != nil {
		task.SetError(errors.WithMessage(err, "failed to find address"))
		return
	}

	if len(unexist) == 0 {
		taskIds, err := s.MintBatchViaRainbowApi(task.UserId, task.AppId, req, exists)
		if err != nil {
			task.SetError(errors.WithMessage(err, "failed to batch mint via rainbow-api"))
			return
		}

		task.Status = enums.BATCH_MINT_STATUS_MINT
		task.MintTaskIds = taskIds
		if err := task.Save(); err != nil {
			task.SetError(errors.WithMessage(err, "failed to save task"))
		}

		return
	}

	s.createWalletAccounts(task, req.Chain, unexist)
	if !task.IsFinalized() {
		if task.Status == enums.BATCH_MINT_STATUS_CREATE_WALLET_DONE {
			s.runBatchMintTask(task, req)
			return
		}
		panic(fmt.Sprintf("unexpected task status %v", task.Status))
	}

}

func (s *MintService) createWalletAccounts(task *models.BatchMintTask, chain enums.Chain, unexists []string) {
	logrus.WithField("phones", unexists).WithField("chain", chain).Info("create wallet by phones")
	switch task.SourceType {
	case enums.SOURCE_TYPE_PHONE:
		task.SetStatus(enums.BATCH_MINT_STATUS_CREATING_WALLET)
		for i, p := range unexists {
			if _, err := s.cellar.GetOrCreateAccount(chain, p); err != nil {
				task.SetError(errors.WithMessage(err, "failed to create wallet"))
				return
			}

			task.SetMessage(fmt.Sprintf("total %d account need create, the %dth created", len(unexists), i))
		}
		task.SetStatus(enums.BATCH_MINT_STATUS_CREATE_WALLET_DONE)
		return

	default:
		task.SetError(errors.New("not support source type"))
		return
	}
}

func (s *MintService) MintBatchFromCerti() {
	panic("unimplemented")
}

func (s *MintService) MintBatchViaRainbowApi(userId, appId uint, req *MintBatchDto, sourceMap map[string]string) ([]uint, error) {
	dto, err := req.ToCustomMintBatchDto(sourceMap)
	if err != nil {
		return nil, err
	}

	taskIds, err := rainbow.NewRainbowApiDashboardClient().MintBatchByMetauri(userId, appId, dto)
	if err != nil {
		return nil, err
	}
	return taskIds, nil
}

func (s *MintService) GetBatchMintTask(userId uint, taskId uint) (*models.BatchMintTask, error) {
	task, err := models.FindBatchMintTaskById(taskId)
	if err != nil {
		return nil, err
	}
	if task.UserId != userId {
		return nil, errors.New("no permission to access this task")
	}
	return task, nil
}

func RunBatchMintTaskOnInit() {
	tasks, err := models.FindUnfinalizedBatchMintTasks()
	if err != nil {
		panic(err)
	}

	logrus.WithField("batch ming tasks", models.GetIds(tasks)).Info("find unfinalized batch mint tasks")

	go func() {
		for _, task := range tasks {
			(&MintService{}).runBatchMintTask(task, nil)
			logrus.WithField("task id", task.ID).Info("batch mint task completed")
		}
	}()
}
