package services

import (
	cryptoRand "crypto/rand"
	"fmt"
	"github.com/nft-rainbow/rainbow-app-service/middlewares"
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/utils"
	openapiclient "github.com/nft-rainbow/rainbow-sdk-go"
	"github.com/spf13/viper"
	"math"
	"math/big"
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

func SetNewYearConfig(config *models.NewYearConfig, id uint) (*models.NewYearConfig, error) {
	config.RainbowUserId = int32(id)
	token, err := middlewares.GenNewYearOpenJWTByRainbowUserId(config.RainbowUserId, config.AppId)
	if err != nil {
		return nil, err
	}
	info, err := GetContractInfo(config.ContractID, "Bearer " + token)
	if err != nil {
		return nil, err
	}
	config.ContractType = *info.Type
	config.ContractAddress = *info.Address
	config.Chain = *info.ChainType
	config.AppId = *info.AppId

	res := models.GetDB().Create(&config)
	if res.Error != nil {
		return nil, res.Error
	}

	return config, nil
}

func HandleSpecialNFTMint(req *POAPRequest)(*openapiclient.ModelsMintTask, error){
	config, err := models.FindNewYearConfigById(int(req.ActivityID))
	if err != nil {
		return nil, err
	}

	commonConfig, err := models.FindNewYearConfigById(viper.GetInt("newYearEvent.newYearCommonId"))
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

	err = burnNFTs(commonConfig, req.UserAddress, token, chainType)
	if err != nil {
		return nil, err
	}

	resp, index, err := randomMint(config, token, req.UserAddress, chainType)
	if err != nil {
		return nil, err
	}

	res := &models.POAPResult{
		ActivityID: int32(config.ID),
		Address: req.UserAddress,
		ContractID: config.ContractID,
		TxID: *resp.Id,
		TokenID: config.ContractInfos[index].TokenID,
	}

	err = models.StorePOAPResult(*res)
	if err != nil {
		return nil, err
	}

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

	item := &models.POAPResult{
		ActivityID: int32(config.ID),
		Address: req.UserAddress,
		ContractID: config.ContractID,
		TxID: *resp.Id,
		TokenID: config.ContractInfos[index].TokenID,
	}

	err = models.StorePOAPResult(*item)
	if err != nil {
		return nil, err
	}

	res, err := models.FindMintCount(req.UserAddress, req.ActivityID)
	if err != nil {
		return nil, err
	}
	if res.Count == 0 {
		everydayNFTMintCache[req.UserAddress] = 3
		_, err = models.UpdateMintCount(req.UserAddress, req.ActivityID, viper.GetInt32("everyDaySharerLimit"))
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
	var amount = int32(1)
	for i := 0; i < len(config.ContractInfos); i++ {
		tmp, _:= models.FindAndCountPOAPResultByTokenId(
			int(config.ID),
			int(config.ContractID),
			0, 10,
			config.ContractInfos[i].TokenID,
			address,
			)
		contractType, err := utils.ContractTypeByTypeId(uint(config.ContractType))
		if err != nil {
			return err
		}
		result, _ := cryptoRand.Int(cryptoRand.Reader, big.NewInt(tmp.Count))

		dto := &openapiclient.ServicesBurnDto{
			Chain: chainType,
			ContractAddress: config.ContractAddress,
			ContractType: contractType,
			User: &address,
			TokenId: tmp.Items[result.Int64()].TokenID,
			Amount: &amount,
		}

		_, err = sendBurnNFTRequest("Bearer " + token, *dto)
		if err != nil {
			return err
		}

		record, err := models.FindPOAPResultById(int(config.ID), int(tmp.Items[result.Int64()].ID))
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
			return fmt.Errorf("The sharer has shared the link to receiver")
		}
	}

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
		everydayNFTMintCache[req.Receiver] = viper.GetInt64("everyDaySharerLimit")
		mintAddressCache[req.Receiver] = []string{}
		_, err := models.UpdateMintCount(req.Sharer, req.ActivityId, 1)
		if err != nil {
			return nil
		}
	}
	mintAddressCache[req.Receiver] = append(mintAddressCache[req.Receiver], req.Sharer)

	return nil
}

func GetSpecialMintCount(activityId int, address string)(int64, error){
	config, err := models.FindNewYearConfigById(activityId)
	if err != nil {
		return 0, err
	}
	res := int64(math.MaxInt64)
	for i := 0; i < viper.GetInt("newYearEvent.commonMintLimit"); i++ {
		resp, err := models.FindAndCountPOAPResultByTokenId(int(config.ID), int(config.ContractID), 0, 10, config.ContractInfos[i].TokenID,address)
		if err != nil {
			return 0, err
		}
		if resp.Count < res {
			res = resp.Count
		}
	}
	return res, nil
}

func UpdateEveryday() {
	c := time.Tick(24 * time.Hour)
	go func() {
		for {
			<- c
			for key := range everydayNFTMintCache {
				everydayNFTMintCache[key] = viper.GetInt64("everyDaySharerLimit")
				_, _ = models.UpdateMintCount(key, viper.GetInt32("newYearActivityId"), 1)
				mintAddressCache[key] = []string{}
			}
		}
	}()
}

func checkEnough(config *models.NewYearConfig, address string)error{
	for i := range config.ContractInfos {
		resp, err := models.FindAndCountPOAPResultByTokenId(int(config.ID), int(config.ContractID), 0, 10, config.ContractInfos[i].TokenID,address)
		if err != nil {
			return err
		}
		if resp.Count <= 0 {
			return fmt.Errorf("The common NFTs are not enough")
		}
	}
	return nil
}

func randomMint(config *models.NewYearConfig, token, address, chain string)(*openapiclient.ModelsMintTask, int32, error) {
	var index int
	probabilities := make([]float32, 0)
	for i := 0; i < len(config.ContractInfos); i ++ {
		probabilities = append(probabilities, config.ContractInfos[i].Probability)
	}
	index = weightedRandomIndex(probabilities)

	resp, err := sendCustomMintRequest("Bearer " + token, openapiclient.ServicesCustomMintDto{
		Chain: chain,
		ContractAddress: config.ContractAddress,
		MintToAddress: address,
		MetadataUri: &(config.ContractInfos[index].MetadataURI),
		TokenId: &(config.ContractInfos[index].TokenID),
	})
	if err != nil {
		return nil, 0, err
	}

	return resp, int32(index), nil

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