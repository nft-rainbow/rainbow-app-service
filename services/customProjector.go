package services

import (
	"errors"
	"sync/atomic"

	"github.com/nft-rainbow/rainbow-app-service/middlewares"
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/utils"
	openapiclient "github.com/nft-rainbow/rainbow-sdk-go"
	"github.com/spf13/viper"
)

func BindDiscordProjectConfig(config *models.DiscordCustomProjectConfig, id uint) error {
	info, err := GetDiscordGuildInfo(config.GuildId)
	if err != nil {
		return err
	}

	config.GuildName = info.Name
	config.RainbowUserId = int32(id)

	res := models.GetDB().Create(&config)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func BindDoDoProjectConfig(config *models.DoDoCustomProjectConfig, id uint) error {
	info, err := GetDoDoIslandInfo(config.IslandId)
	if err != nil {
		return err
	}

	config.IslandName = info.IslandName
	config.RainbowUserId = int32(id)

	res := models.GetDB().Create(&config)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func DiscordCustomActivityConfig(config *models.DiscordCustomActivityConfig, id uint) error {
	token, err := middlewares.GenDiscordOpenJWTByRainbowUserId(id)
	if err != nil {
		return err
	}
	info, err := GetContractInfo(config.ContractID, token)
	if err != nil {
		return err
	}
	config.ContractType = *info.Type
	config.Chain = *info.ChainType
	config.AppId = *info.AppId
	config.ContractAddress = *info.Address

	res := models.GetDB().Create(&config)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func DoDoCustomActivityConfig(config *models.DoDoCustomActivityConfig, id uint) error {
	token, err := middlewares.GenDoDoOpenJWTByRainbowUserId(id)
	if err != nil {
		return err
	}
	info, err := GetContractInfo(config.ContractID, token)
	if err != nil {
		return err
	}
	config.ContractType = *info.Type
	config.Chain = *info.ChainType
	config.AppId = *info.AppId
	config.ContractAddress = *info.Address

	res := models.GetDB().Create(&config)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func HandleCustomMint(userId, channelId, platform string) (*openapiclient.ModelsMintTask, string, int32, error) {
	req := models.CustomMintReq{
		UserID:    userId,
		ChannelID: channelId,
	}
	if platform == "dodo" {
		resp, token, contractId, err := dodoCustomMint(&req)
		return resp, token, contractId, err
	} else if platform == "discord" {
		resp, token, contractId, err := discordCustomMint(&req)
		return resp, token, contractId, err
	} else {
		return nil, "", 0, nil
	}
}

func dodoCustomMint(req *models.CustomMintReq) (*openapiclient.ModelsMintTask, string, int32, error) {
	config, err := models.FindDoDoCustomActivityConfigByChannelId(req.ChannelID)
	if err != nil {
		return nil, "", 0, err
	}

	ok, err := models.CheckDoDoCustomCount(req.UserID, req.ChannelID, config.MaxMintCount)
	if err != nil {
		return nil, "", 0, err
	}
	if !ok {
		return nil, "", 0, errors.New("This number of the NFTs the account minted has reached the maximum")
	}

	cfxAddress, err := GetDoDoBindCFXAddress(req.UserID)
	if err != nil {
		return nil, "", 0, err
	}

	token, _ := middlewares.GenerateDoDoOpenJWT(req.ChannelID)
	chain, err := utils.ChainById(uint(config.Chain))
	if err != nil {
		return nil, "", 0, err
	}

	resp, err := sendCustomMintRequest(middlewares.PrefixToken(token), openapiclient.ServicesCustomMintDto{
		Chain:           chain,
		ContractAddress: config.ContractAddress,
		MintToAddress:   cfxAddress,
		MetadataUri:     &config.MetadataURI,
	})
	if err != nil {
		return nil, "", 0, err
	}

	return resp, token, config.ContractID, err
}

func discordCustomMint(req *models.CustomMintReq) (*openapiclient.ModelsMintTask, string, int32, error) {
	config, err := models.FindDiscordCustomActivityConfigByChannelId(req.ChannelID)
	if err != nil {
		return nil, "", 0, err
	}

	ok, err := models.CheckDiscordCustomCount(req.UserID, req.ChannelID, config.MaxMintCount)
	if err != nil {
		return nil, "", 0, err
	}
	if !ok {
		return nil, "", 0, errors.New("This number of the NFTs the account minted has reached the maximum")
	}

	cfxAddress, err := GetDiscordBindCFXAddress(req.UserID)
	if err != nil {
		return nil, "", 0, err
	}

	token, _ := middlewares.GenerateDiscordOpenJWT(req.ChannelID)

	chain, err := utils.ChainById(uint(config.Chain))
	if err != nil {
		return nil, "", 0, err
	}

	resp, err := sendCustomMintRequest(middlewares.PrefixToken(token), openapiclient.ServicesCustomMintDto{
		Chain:           chain,
		ContractAddress: config.ContractAddress,
		MintToAddress:   cfxAddress,
		MetadataUri:     &config.MetadataURI,
	})
	if err != nil {
		return nil, "", 0, err
	}

	return resp, token, config.ContractID, err
}

// ChangeAnDao NFT counter
var changAnDaoNum uint64

func InitChangAnDaoNum() {
	var count int64
	models.GetDB().Model(&models.POAPResult{}).Where("activity_id = ? and status = ?", viper.GetString("changAnDao.activityId"), models.STATUS_SUCCESS).Count(&count)
	atomic.StoreUint64(&changAnDaoNum, uint64(count))
}

func IncreaseChangAnDaoNum() {
	atomic.AddUint64(&changAnDaoNum, 1)
}

func GetChangAnDaoNum() uint64 {
	return atomic.LoadUint64(&changAnDaoNum)
}
