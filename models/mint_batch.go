package models

import (
	"time"

	"github.com/nft-rainbow/rainbow-app-service/models/enums"
	"github.com/nft-rainbow/rainbow-app-service/utils"
	"gorm.io/datatypes"
)

type BatchMintTask struct {
	BaseModel
	UserId          uint                      `json:"user_id"`
	AppId           uint                      `json:"app_id"`
	SourceType      enums.SourceType          `json:"source_type" swaggertype:"string"`
	RequestFilePath string                    `json:"-"`
	Status          enums.BatchMintStatus     `json:"status" swaggertype:"string"`
	Error           string                    `json:"error"`
	Message         string                    `json:"message,omitempty"`
	MintTaskIds     datatypes.JSONSlice[uint] `json:"mint_task_ids"`
	// SourceAddresses map[string]string     `gorm:"-" json:"sourceAddresses"`
}

func (t *BatchMintTask) IsFinalized() bool {
	return t.IsFail() || t.IsSuceess()
}

func (t *BatchMintTask) IsFail() bool {
	return t.Status == enums.BATCH_MINT_STATUS_FAIL

}

func (t *BatchMintTask) IsSuceess() bool {
	return t.Status == enums.BATCH_MINT_STATUS_MINT
}

func (t *BatchMintTask) Save() error {
	return utils.Retry(3, time.Second, func() error {
		return GetDB().Save(t).Error
	})
}

func (t *BatchMintTask) SetStatus(status enums.BatchMintStatus) {
	t.Status = status
	t.Save()
}

func (t *BatchMintTask) SetError(err error) {
	t.Error = err.Error()
	t.SetStatus(enums.BATCH_MINT_STATUS_FAIL)
}

func (t *BatchMintTask) SetMessage(msg string) {
	t.Message = msg
	t.Save()
}

func FindUnfinalizedBatchMintTasks() ([]*BatchMintTask, error) {
	var tasks []*BatchMintTask
	err := GetDB().Model(&BatchMintTask{}).
		Where("status not in (?)", []enums.BatchMintStatus{enums.BATCH_MINT_STATUS_MINT, enums.BATCH_MINT_STATUS_FAIL}).
		Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func FindBatchMintTaskById(id uint) (*BatchMintTask, error) {
	var task BatchMintTask
	err := GetDB().Model(&BatchMintTask{}).
		Where("id = ?", id).
		First(&task).Error
	if err != nil {
		return nil, err
	}
	return &task, nil
}
