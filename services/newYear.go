package services

import (
	"fmt"
	"github.com/nft-rainbow/rainbow-app-service/middlewares"
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/utils"
	openapiclient "github.com/nft-rainbow/rainbow-sdk-go"
	"math/rand"
	"time"
)

type ShareRequest struct {
	Sharer string `json:"sharer"`
	Receiver string `json:"receiver"`
	ActivityId int32 `json:"activity_id"`
}

var everydayNFTMintCache = make(map[string]int64)
var mintAddressCache = make(map[string][]string)
const newYearActivityId = 11

func SetNewYearConfig(config *models.NewYearConfig, id uint) (*models.NewYearConfig, error) {
	config.RainbowUserId = int32(id)
	token, err := middlewares.GenNewYearOpenJWTByRainbowUserId(config.RainbowUserId, config.AppId)
	if err != nil {
		return nil, err
	}

	for i := range config.ContractInfos {
		info, err := GetContractInfo(config.ContractInfos[i].ContractID, "Bearer " + token)
		if err != nil {
			return nil, err
		}
		config.ContractInfos[i].ContractType = *info.Type
		config.ContractInfos[i].ContractAddress = *info.Address

		config.Chain = *info.ChainType
		config.AppId = *info.AppId
	}

	res := models.GetDB().Create(&config)
	if res.Error != nil {
		return nil, res.Error
	}

	return config, nil
}

func HandleSpecialNFTMint(commonActivityID int, req *POAPRequest)(*openapiclient.ModelsMintTask, error){
	specialConfig, err := models.FindNewYearConfigById(int(req.ActivityID))
	if err != nil {
		return nil, err
	}
	commonConfig, err := models.FindNewYearConfigById(commonActivityID)
	if err != nil {
		return nil, err
	}

	token, err := middlewares.GenNewYearOpenJWTByRainbowUserId(specialConfig.RainbowUserId, specialConfig.AppId)
	if err != nil {
		return nil, err
	}

	err = newYearCommonCheck(specialConfig.StartedTime, specialConfig.EndedTime, int(specialConfig.ID), int32(specialConfig.Amount))
	if err != nil {
		return nil, err
	}

	chainType, err := utils.ChainTypeByTypeId(uint(specialConfig.Chain))
	if err != nil {
		return nil, err
	}

	err = burnNFTs(commonConfig, req.UserAddress, token, chainType)
	if err != nil {
		return nil, err
	}

	resp, index, err := randomMint(specialConfig, token, req.UserAddress, chainType)
	if err != nil {
		return nil, err
	}

	err = models.StorePOAPResult(models.POAPResult{
		ActivityID: int32(specialConfig.ID),
		Address: req.UserAddress,
		ContractID: specialConfig.ContractInfos[index].ContractID,
		TxID: *resp.Id,
	})
	if err != nil {
		return nil, err
	}

	go SyncNFTMintTaskStatus(token, int32(specialConfig.ID))

	return resp, nil
}

func HandleCommonNFTMint(req *POAPRequest)(*openapiclient.ModelsMintTask, error) {
	err := checkMintCount(req.ActivityID, req.UserAddress)
	if err != nil {
		return nil, err
	}

	config, err := models.FindNewYearConfigById(int(req.ActivityID))
	if err != nil {
		return nil, err
	}

	token, err := middlewares.GenNewYearOpenJWTByRainbowUserId(config.RainbowUserId, config.AppId)
	if err != nil {
		return nil, err
	}
	err = newYearCommonCheck(config.StartedTime, config.EndedTime, int(config.ID), config.Amount)
	if err != nil {
		return nil, err
	}

	chainType, err := utils.ChainTypeByTypeId(uint(config.Chain))
	if err != nil {
		return nil, err
	}

	resp, index, err := randomMint(config, token, req.UserAddress, chainType)
	if err != nil {
		return nil, err
	}

	err = models.StorePOAPResult(models.POAPResult{
		ActivityID: int32(config.ID),
		Address: req.UserAddress,
		ContractID: config.ContractInfos[index].ContractID,
		TxID: *resp.Id,
	})
	if err != nil {
		return nil, err
	}

	go SyncNFTMintTaskStatus(token, int32(config.ID))
	res, _ := models.FindMintCount(req.UserAddress, req.ActivityID)
	if res.Count == 0 {
		everydayNFTMintCache[req.UserAddress] = 3
		_, err = models.UpdateMintCount(req.UserAddress, req.ActivityID, 3)
		if err != nil {
			return nil, err
		}
	}
	_, err = models.UpdateMintCount(req.UserAddress, req.ActivityID, -1)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func burnNFTs(config *models.NewYearConfig, address, token, chainType string) error{
	err := checkEnough(config, address)
	if err != nil {
		return err
	}
	for i, v := range config.ContractInfos {
		tmp, _:= models.FindAndCountPOAPResultByAddresses(int(config.ID), int(v.ContractID), 0, 10, address)
		contractType, err := utils.ContractTypeByTypeId(uint(v.ContractType))
		if err != nil {
			return err
		}
		dto := &openapiclient.ServicesBurnDto{
			Chain: chainType,
			ContractAddress: config.ContractInfos[i].ContractAddress,
			ContractType: contractType,
			User: &address,
			TokenId: tmp.Items[0].TokenID,
		}
		//dto := *openapiclient.NewServicesBurnDto(
		//	chainType,
		//	config.ContractInfos[i].ContractAddress,
		//	contractType,
		//	tmp.Items[0].TokenID,
		//)

		_, err = sendBurnNFTRequest("Bearer " + token, *dto)
		if err != nil {
			return err
		}

		record, err := models.FindPOAPResultById(int(config.ID), int(tmp.Items[0].ID))
		if err != nil {
			return err
		}
		res := models.GetDB().Delete(record)
		if res.Error != nil {
			return res.Error
		}
	}
	return nil
}

func UpdateBySharing(req ShareRequest)error {
	for _, v := range mintAddressCache[req.Receiver] {
		if v == req.Sharer {
			return nil
		}
	}

	mintAddressCache[req.Receiver] = append(mintAddressCache[req.Receiver], req.Sharer)

	if everydayNFTMintCache[req.Sharer] > 0 {
		_, err := models.UpdateMintCount(req.Sharer, req.ActivityId, 1)
		if err != nil {
			return err
		}
		everydayNFTMintCache[req.Sharer] -= 1
	}

	resp, err := models.FindMintCount(req.Receiver, req.ActivityId)
	if err != nil {
		return err
	}
	if resp.Count == 0 {
		everydayNFTMintCache[req.Receiver] = 3
		mintAddressCache[req.Receiver] = []string{}
		_, err := models.UpdateMintCount(req.Sharer, req.ActivityId, 1)
		if err != nil {
			return nil
		}
	}

	return nil
}

func GetSpecialMintRemain(activityId int, address string)(int, int, error){
	config, err := models.FindNewYearConfigById(activityId)
	if err != nil {
		return 0, 0, err
	}
	res := 0
	for _, v := range config.ContractInfos {
		resp, err := models.FindAndCountPOAPResultByAddresses(int(config.ID), int(v.ContractID), 0, 10, address)
		if err != nil {
			return 0, 0, err
		}
		if resp.Count != 0 {
			res ++
		}
	}
	return res, int(config.MaxMintCount), nil
}

func UpdateEveryday() {
	c := time.Tick(24 * time.Hour)
	go func() {
		for {
			<- c
			for key := range everydayNFTMintCache {
				everydayNFTMintCache[key] = 3
				_, _ = models.UpdateMintCount(key, int32(newYearActivityId), 1)
				mintAddressCache[key] = []string{}
			}
		}
	}()
}

func checkEnough(config *models.NewYearConfig, address string)error{
	for i := range config.ContractInfos {
		resp, err := models.FindAndCountPOAPResultByAddresses(int(config.ID), int(config.ContractInfos[i].ContractID), 0, 10, address)
		if err != nil {
			return err
		}
		if resp.Count == 0 {
			return fmt.Errorf("The common NFTs are not enough")
		}
	}
	return nil
}

func randomMint(config *models.NewYearConfig, token, address, chain string)(*openapiclient.ModelsMintTask, int, error) {
	probabilities := make([]float32, 0)
	for _, v := range config.ContractInfos {
		probabilities = append(probabilities, v.Probability)
	}

	index := weightedRandomIndex(probabilities)
	resp, err := sendCustomMintRequest("Bearer " + token, openapiclient.ServicesCustomMintDto{
		Chain: chain,
		ContractAddress: config.ContractInfos[index].ContractAddress,
		MintToAddress: address,
		MetadataUri: &(config.ContractInfos[index].MetadataURI),
	})
	if err != nil {
		return nil, 0, err
	}

	return resp, index, nil

}

func weightedRandomIndex(weights []float32) int {
	if len(weights) == 1 {
		return 0
	}
	var sum float32 = 0.0
	for _, w := range weights {
		sum += w
	}
	r := rand.Float32() * sum
	var t float32 = 0.0
	for i, w := range weights {
		t += w
		if t > r {
			return i
		}
	}
	return len(weights) - 1
}

func newYearCommonCheck(startTime, endTime int64, configID int, amount int32)error{
	if startTime != -1 && time.Now().Unix() < startTime {
		return fmt.Errorf("The activity has not been started")
	}

	if endTime != -1 && time.Now().Unix() > endTime {
		return fmt.Errorf("The activity has been expired")
	}

	err := checkNewYearAmount(configID, amount)
	if err != nil {
		return err
	}
	return nil
}

func checkNewYearAmount(configID int, amount int32) error {
	if amount != -1 {
		resp, err := models.FindAndCountPOAPResult(configID, 0, 10)
		if err != nil {
			return err
		}
		if int32(resp.Count) >= amount{
			return fmt.Errorf("The mint amount has exceeded the limit")
		}
	}
	return nil
}

func checkMintCount(activityId int32, address string) error{
	if _, ok := everydayNFTMintCache[address]; !ok {
		return nil
	}
	resp, err := models.FindMintCount(address, activityId)
	if err != nil {
		return err
	}

	if resp.Count <= 0 {
		return fmt.Errorf("The mint count is not enough")
	}
	return nil
}