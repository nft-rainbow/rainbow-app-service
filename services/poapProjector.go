package services

import (
	"fmt"
	"github.com/nft-rainbow/rainbow-app-service/middlewares"
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/utils"
	openapiclient "github.com/nft-rainbow/rainbow-sdk-go"
	"time"
)

type POAPRequest struct {
	ActivityID  string `json:"activity_id" binding:"required"`
	UserAddress string `json:"user_address" binding:"required"`
	Command     string `json:"command"`
}

func POAPActivityConfig(config *models.POAPActivityConfig, id uint) (*models.POAPActivityConfig, error) {
	config.RainbowUserId = int32(id)
	token, err := middlewares.GenPOAPOpenJWTByRainbowUserId(*config)
	if err != nil {
		return nil, err
	}

	info, err := GetContractInfo(config.ContractID, "Bearer "+token)
	if err != nil {
		return nil, err
	}
	config.ContractType = *info.Type
	config.Chain = *info.ChainType
	config.AppId = *info.AppId
	config.ContractAddress = *info.Address

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

func HandlePOAPCSVMint(req *POAPRequest) (*models.POAPResult, error) {
	config, err := models.FindPOAPActivityConfigById(req.ActivityID)
	if err != nil {
		return nil, err
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
	chainType, err := utils.ChainTypeByTypeId(uint(config.Chain))
	if err != nil {
		return nil, err
	}

	_, err = models.CountPOAPResult(req.ActivityID)
	if err != nil {
		return nil, err
	}

	resp, err := sendCustomMintRequest("Bearer "+token, openapiclient.ServicesCustomMintDto{
		Chain:           chainType,
		ContractAddress: config.ContractAddress,
		MintToAddress:   req.UserAddress,
		MetadataUri:     &config.MetadataURI,
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
	cache := models.Cache[config.ActivityID]
	cache.Lock()
	cache.Count += 1
	cache.Unlock()

	go SyncNFTMintTaskStatus(token, item)

	return item, res.Error
}

func HandlePOAPH5Mint(req *POAPRequest) (*models.POAPResult, error) {
	config, err := models.FindPOAPActivityConfigById(req.ActivityID)
	if err != nil {
		return nil, err
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

	_, err = models.CountPOAPResult(req.ActivityID)
	if err != nil {
		return nil, err
	}

	chainType, err := utils.ChainTypeByTypeId(uint(config.Chain))
	if err != nil {
		return nil, err
	}

	resp, err := sendCustomMintRequest("Bearer "+token, openapiclient.ServicesCustomMintDto{
		Chain:           chainType,
		ContractAddress: config.ContractAddress,
		MintToAddress:   req.UserAddress,
		MetadataUri:     &config.MetadataURI,
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
	cache := models.Cache[config.ActivityID]
	cache.Lock()
	cache.Count += 1
	cache.Unlock()

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
