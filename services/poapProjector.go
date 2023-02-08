package services

import (
	"fmt"
	"github.com/nft-rainbow/rainbow-app-service/middlewares"
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/utils"
	openapiclient "github.com/nft-rainbow/rainbow-sdk-go"
	"strings"
	"time"
)

type POAPRequest struct {
	ActivityID  string `json:"activity_id" binding:"required"`
	UserAddress string `json:"user_address" binding:"required"`
	Command     string `json:"command"`
}

func POAPActivityConfig(config *models.POAPActivityConfig, id uint) (*models.POAPActivityConfig, error) {
	config.RainbowUserId = int32(id)

	poapId, err := getPoAPId(config.ContractAddress, config.Name)
	if err != nil {
		return nil, err
	}

	config.ActivityID = poapId

	res := models.GetDB().Create(&config)
	if res.Error != nil {
		return nil, res.Error
	}

	return config, nil
}

func POAPH5Config(config *models.H5Config) (*models.H5Config, error) {
	res := models.GetDB().Create(&config)
	if res.Error != nil {
		return nil, res.Error
	}

	return config, nil
}

func UpdatePOAPActivityConfig(config *models.POAPActivityConfig, activityId string) (*models.POAPActivityConfig, error) {
	oldConfig, err := models.FindPOAPActivityConfigById(activityId)
	if err != nil {
		return nil, err
	}
	if oldConfig.ContractID == 0 || config.ContractID != oldConfig.ContractID {
		token, err := middlewares.GenPOAPOpenJWTByRainbowUserId(*oldConfig)
		if err != nil {
			fmt.Println(555555)
			return nil, err
		}
		info, err := GetContractInfo(config.ContractID, "Bearer "+token)
		if err != nil {
			return nil, err
		}
		oldConfig.ContractType = *info.Type
		oldConfig.ChainId = *info.ChainId
		oldConfig.ChainType = *info.ChainType
		oldConfig.AppId = *info.AppId
		oldConfig.ContractAddress = *info.Address
		oldConfig.ContractID = config.ContractID
	}

	oldConfig.NFTConfigs = config.NFTConfigs
	oldConfig.ActivityType = config.ActivityType
	oldConfig.Command = config.Command
	oldConfig.StartedTime = config.StartedTime
	oldConfig.EndedTime = config.EndedTime
	oldConfig.Amount = config.Amount
	oldConfig.ActivityPictureURL = config.ActivityPictureURL
	oldConfig.Name = config.Name
	oldConfig.Description = config.Description
	oldConfig.WhiteListInfos = config.WhiteListInfos

	if oldConfig.NFTConfigs != nil {
		for _, v := range config.NFTConfigs {
			tmp := strings.Split(v.Name, "/")
			err = AddLogoAndUpload(v.ImageURL, tmp[len(tmp)-1], oldConfig.ActivityID)
			if err != nil {
				return nil, err
			}
		}
	}

	models.GetDB().Save(&oldConfig)

	return oldConfig, nil
}

func HandlePOAPCSVMint(req *POAPRequest) (*models.POAPResult, error) {
	config, err := models.FindPOAPActivityConfigById(req.ActivityID)
	if err != nil {
		return nil, err
	}

	if config.WhiteListInfos == nil && len(config.WhiteListInfos) == 0 {
		return nil, fmt.Errorf("The activity has not opened the white list")
	}

	token, err := middlewares.GeneratePOAPOpenJWT(config.Name, config.ContractID)
	if err != nil {
		return nil, err
	}

	err = commonCheck(config, req)
	if err != nil {
		return nil, err
	}

	if len(config.WhiteListInfos) == 0 || !checkWhiteList(config.WhiteListInfos, req.UserAddress) {
		return nil, fmt.Errorf("The address is not listed in the white list")
	}

	err = checkWhiteListLimit(config, req.UserAddress)
	if err != nil {
		return nil, err
	}
	chain, err := utils.ChainById(uint(config.ChainId))
	if err != nil {
		return nil, err
	}

	var metadataURI *string
	if config.ActivityType == "single" {
		metadataURI, err = createMetadata(config, token, 0)
		if err != nil {
			return nil, err
		}
	} else if config.ActivityType == "blind_box" {
		var index int
		probabilities := make([]float32, 0)
		for i := 0; i < len(config.NFTConfigs); i++ {
			probabilities = append(probabilities, config.NFTConfigs[i].Probability)
		}
		index = weightedRandomIndex(probabilities)
		metadataURI, err = createMetadata(config, token, index)
		if err != nil {
			return nil, err
		}
	}

	resp, err := sendCustomMintRequest("Bearer "+token, openapiclient.ServicesCustomMintDto{
		Chain:           chain,
		ContractAddress: config.ContractAddress,
		MintToAddress:   req.UserAddress,
		MetadataUri:     metadataURI,
	})
	if err != nil {
		return nil, err
	}

	item := &models.POAPResult{
		ConfigID:   int32(config.ID),
		Address:    req.UserAddress,
		ContractID: config.ContractID,
		TxID:       *resp.Id,
		ActivityID: config.ActivityID,
	}

	res := models.GetDB().Create(&item)

	go SyncNFTMintTaskStatus(token, item)

	return item, res.Error
}

func HandlePOAPH5Mint(req *POAPRequest) (*models.POAPResult, error) {
	config, err := models.FindPOAPActivityConfigById(req.ActivityID)
	if err != nil {
		return nil, err
	}
	if config.WhiteListInfos == nil && len(config.WhiteListInfos) == 0 {
		return nil, fmt.Errorf("The activity has opened the white list")
	}

	token, err := middlewares.GeneratePOAPOpenJWT(config.Name, config.ContractID)
	if err != nil {
		return nil, err
	}

	err = commonCheck(config, req)
	if err != nil {
		return nil, err
	}

	err = checkLimitAmount(config, req.UserAddress)
	if err != nil {
		return nil, err
	}

	chain, err := utils.ChainById(uint(config.ChainId))
	if err != nil {
		return nil, err
	}

	var metadataURI *string
	var index int
	if config.ActivityType == "single" {
		metadataURI, err = createMetadata(config, token, 0)
		if err != nil {
			return nil, err
		}
	} else if config.ActivityType == "blind_box" {
		probabilities := make([]float32, 0)
		for i := 0; i < len(config.NFTConfigs); i++ {
			probabilities = append(probabilities, config.NFTConfigs[i].Probability)
		}
		index = weightedRandomIndex(probabilities)
		metadataURI, err = createMetadata(config, token, index)
		if err != nil {
			return nil, err
		}
	}

	resp, err := sendCustomMintRequest("Bearer "+token, openapiclient.ServicesCustomMintDto{
		Chain:           chain,
		ContractAddress: config.ContractAddress,
		MintToAddress:   req.UserAddress,
		MetadataUri:     metadataURI,
	})
	if err != nil {
		return nil, err
	}

	item := &models.POAPResult{
		ConfigID:   int32(config.ID),
		Address:    req.UserAddress,
		ContractID: config.ContractID,
		TxID:       *resp.Id,
		ActivityID: config.ActivityID,
		FileURL:    config.NFTConfigs[index].ImageURL,
	}

	res := models.GetDB().Create(&item)

	go SyncNFTMintTaskStatus(token, item)

	return item, res.Error
}

func GetMintCount(activityID, address string) (*MintCountResponse, error) {
	config, err := models.FindPOAPActivityConfigById(activityID)
	if err != nil {
		return nil, err
	}
	resp, err := models.CountPOAPResultByAddress(address, activityID)
	if err != nil {
		return nil, err
	}
	var count int32
	remainedMinted := int32(int64(config.MaxMintCount) - resp)

	if config.Amount == -1 {
		count = remainedMinted
	} else {
		res, err := models.CountPOAPResult(address)
		if err != nil {
			return nil, err
		}
		if config.Amount-int32(res) < remainedMinted {
			count = config.Amount - int32(res)
		} else {
			count = remainedMinted
		}
	}
	return &MintCountResponse{
		Address:    address,
		ActivityID: activityID,
		Count:      count,
	}, nil
}

func commonCheck(config *models.POAPActivityConfig, req *POAPRequest) error {
	if req.Command != config.Command {
		return fmt.Errorf("The command is worng")
	}

	if config.StartedTime != -1 &&
		config.EndedTime != -1 &&
		(time.Now().Unix() < config.StartedTime || time.Now().Unix() > config.EndedTime) {
		return fmt.Errorf("The activity has already expired or has not been started")
	}

	err := checkAmount(config)
	if err != nil {
		return err
	}
	return nil
}

func checkWhiteList(whiteList []models.WhiteListInfo, address string) bool {
	for _, v := range whiteList {
		if address == v.User {
			return true
		}
	}
	return false
}

func checkAmount(config *models.POAPActivityConfig) error {
	if config.Amount != -1 {
		resp, err := models.CountPOAPResult(config.ActivityID)
		if err != nil {
			return err
		}
		if int32(resp) >= config.Amount {
			return fmt.Errorf("The mint amount has exceeded the limit")
		}
	}
	return nil
}

func checkLimitAmount(config *models.POAPActivityConfig, address string) error {
	resp, err := models.CountPOAPResultByAddress(address, config.ActivityID)
	if err != nil {
		return err
	}

	if config.MaxMintCount == -1 {
		return nil
	}
	if resp >= int64(config.MaxMintCount) {
		return fmt.Errorf("The mint amount has exceeded the mint limit")
	}
	return nil
}

func checkWhiteListLimit(config *models.POAPActivityConfig, address string) error {
	resp, err := models.CountPOAPResultByAddress(address, config.ActivityID)
	if err != nil {
		return err
	}
	for _, v := range config.WhiteListInfos {
		if v.User == address && resp >= int64(v.Count) {
			return fmt.Errorf("The NFT minted by the account has exceeded the mint limit")
		}
	}
	return nil
}

func createMetadata(config *models.POAPActivityConfig, token string, index int) (*string, error) {
	attributes := make([]openapiclient.ModelsExposedMetadataAttribute, 0)
	for _, v := range config.NFTConfigs[index].MetadataAttributes {
		attributes = append(attributes, openapiclient.ModelsExposedMetadataAttribute{
			AttributeName: &v.Name,
			DisplayType:   &v.DisplayType,
			TraitType:     &v.TraitType,
			Value:         &v.Value,
		})
		now := time.Now().Format("2006-01-02 15:04:05 MST Mon")
		name := "mint_time"
		trait := "time"
		display := "date"
		attributes = append(attributes, openapiclient.ModelsExposedMetadataAttribute{
			AttributeName: &name,
			Value:         &now,
			TraitType:     &trait,
			DisplayType:   &display,
		})
	}

	resp, err := sendCreateMetadataRequest("Bearer "+token, openapiclient.ServicesMetadataDto{
		Description: config.Description,
		Image:       config.NFTConfigs[index].ImageURL,
		Name:        config.NFTConfigs[index].Name,
		Attributes:  attributes,
	})
	if err != nil {
		return nil, err
	}

	return resp.Uri, nil
}
