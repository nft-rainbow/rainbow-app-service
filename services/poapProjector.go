package services

import (
	"context"
	"fmt"
	"github.com/nft-rainbow/rainbow-app-service/middlewares"
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/utils"
	openapiclient "github.com/nft-rainbow/rainbow-sdk-go"
	"strconv"
	"time"
)

type POAPRequest struct {
	ActivityID int32 `json:"activity_id" binding:"required"`
	UserAddress string `json:"user_address"`
	Command string `json:"command"`
}

func POAPActivityConfig(config *models.POAPActivityConfig, id uint) (*models.POAPActivityConfig, error) {
	config.RainbowUserId = int32(id)
	token, err := middlewares.GenPOAPOpenJWTByRainbowUserId(*config)
	if err != nil {
		return nil, err
	}

	if !config.IsCommandNeeded && config.Command == ""{
		return nil, fmt.Errorf("The corresponding command is needed.")
	}

	info, err := GetContractInfo(config.ContractID, "Bearer " + token)
	if err != nil {
		return nil, err
	}
	config.ContractType = *info.Type
	config.Chain = *info.ChainType
	config.AppId = *info.AppId
	config.ContractAddress = *info.Address

	res := models.GetDB().Create(&config)
	if res.Error != nil {
		return nil, res.Error
	}

	return config, nil
}

func POAPH5Config(config *models.H5Config, id uint) (*models.H5Config, error) {
	config.RainbowUserId = int32(id)

	res := models.GetDB().Create(&config)
	if res.Error != nil {
		return nil, res.Error
	}

	return config, nil
}

func HandlePOAPCSVMint(records [][]string, req *POAPRequest) ([]openapiclient.ModelsMintTask, error){
	config, err := models.FindPOAPActivityConfigById(int(req.ActivityID))
	if err != nil {
		return nil, err
	}

	token, err := middlewares.GeneratePOAPOpenJWT(config.Name, config.ContractID)
	if err != nil {
		return nil, err
	}


	if config.StartedTime != -1 && config.EndedTime != -1 && (time.Now().Unix() < config.StartedTime  || time.Now().Unix() > config.EndedTime) {
		return nil, fmt.Errorf("The activity has already expired or has not been started")
	}

	mintItems := make([]openapiclient.ServicesMintItemDto,0)
	for _, row := range records {
		var address string
		var count int32
		for i := 0; i < 2; i ++ {
			if i == 0 {
				address = row[0]
				if address[0] != 'c' {
					address = address[3:]
				}
			}else {
				tmp, _ := strconv.Atoi(row[1])
				count = int32(tmp)
			}
		}
		tmp := &openapiclient.ServicesMintItemDto{
			Amount: &count,
			MetadataUri: &config.MetadataURI,
			MintToAddress: address,
		}

		mintItems = append(mintItems, *tmp)
	}

	chainType, err := utils.ChainTypeByTypeId(uint(config.Chain))
	if err != nil {
		return nil, err
	}
	contractType, err := utils.ContractTypeByTypeId(uint(config.ContractType))
	if err != nil {
		return nil, err
	}

	dto := &openapiclient.ServicesCustomMintBatchDto{
		Chain: chainType,
		ContractAddress: config.ContractAddress,
		ContractType: contractType,
		MintItems: mintItems,
	}

	resp, _, err := newClient().MintsApi.BatchCustomMint(context.Background()).Authorization("Bearer " + token).CustomMintBatchDto(*dto).Execute()
	if err != nil {
		return nil, err
	}

	for _, task := range resp {
		err = models.StorePOAPResult(models.POAPResult{
			ActivityID: int32(config.ID),
			Address: *task.MintTo,
			ContractID: config.ContractID,
			TxID: *task.Id,
		})
		if err != nil {
			return nil, err
		}
	}

	go SyncNFTMintTaskStatus(token, int32(config.ID))

	return resp, nil
}

func HandlePOAPH5Mint(req *POAPRequest) (*openapiclient.ModelsMintTask, error){
	config, err := models.FindPOAPActivityConfigById(int(req.ActivityID))
	if err != nil {
		return nil, err
	}

	token, err := middlewares.GeneratePOAPOpenJWT(config.Name, config.ContractID)
	if err != nil {
		return nil, err
	}

	if config.IsCommandNeeded && req.Command != config.Command{
		return nil, fmt.Errorf("The command is worng")
	}

	if config.StartedTime != -1 && config.EndedTime != -1 && (time.Now().Unix() < config.StartedTime  || time.Now().Unix() > config.EndedTime) {

		return nil, fmt.Errorf("The activity has already expired or has not been started")
	}

	chainType, err := utils.ChainTypeByTypeId(uint(config.Chain))
	if err != nil {
		return nil, err
	}
	contractType, err := utils.ContractTypeByTypeId(uint(config.ContractType))
	if err != nil {
		return nil, err
	}

	resp, err := sendCustomMintRequest("Bearer " + token, openapiclient.ServicesCustomMintDto{
		Chain: chainType,
		ContractType: contractType,
		ContractAddress: config.ContractAddress,
		MintToAddress: req.UserAddress,
		MetadataUri: &config.MetadataURI,
	})
	if err != nil {
		return nil, err
	}

	err = models.StorePOAPResult(models.POAPResult{
		ActivityID: int32(config.ID),
		Address: req.UserAddress,
		ContractID: config.ContractID,
		TxID: *resp.Id,
	})
	if err != nil {
		return nil, err
	}

	go SyncNFTMintTaskStatus(token, int32(config.ID))

	return resp, nil
}
