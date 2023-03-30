package services

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/nft-rainbow/rainbow-app-service/middlewares"
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/utils"
	"github.com/nft-rainbow/rainbow-app-service/utils/rand"
	openapiclient "github.com/nft-rainbow/rainbow-sdk-go"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BotActivityService struct {
	authcodes        sync.Map
	socialToolClient SocialToolBot
}

func NewBotActivityService(socialToolType models.SocialToolType) (*BotActivityService, error) {
	_socialToolClient, err := getSocialToolBot(socialToolType)
	if err != nil {
		return nil, err
	}
	return &BotActivityService{
		socialToolClient: _socialToolClient,
	}, nil
}

func (d *BotActivityService) VerifyUser(user models.SocialToolUser) (*VerifyUserResponse, error) {
	v, loaded := d.authcodes.LoadOrStore(user, rand.NumString(6))
	if !loaded {
		go func() {
			<-time.After(time.Minute * 5)
			d.authcodes.Delete(user)
		}()
	}

	msg := fmt.Sprintf("You are setting Rainbow-Bot, your Rainbow auth code is %v, please fill back to Rainbow to complete authentication.", v)
	if err := d.socialToolClient.SendDirectMessage(context.Background(), user.UserSocialId, msg); err != nil {
		return nil, err
	}

	return &VerifyUserResponse{AuthCode: v.(string)}, nil
}

func (d *BotActivityService) InsertProjectManager(userId uint, socialTool models.SocialToolType, req InsertProjectorReq) error {
	socialUser := models.SocialToolUser{
		SocialTool:   socialTool,
		UserSocialId: req.UserSocialId,
	}
	if socialTool == models.SOCIAL_TOOL_DODO {
		code, ok := d.authcodes.Load(socialUser)
		if !ok || code.(string) != *req.AuthCode {
			return errors.New("auth code not match")
		}
	} else {
		return errors.New("unsupported social tool")
	}

	_, err := models.FindSocialToolProjectManager(userId, socialTool)
	if err != gorm.ErrRecordNotFound {
		return errors.New("already exists")
	}

	var p models.SocialToolProjectManager
	p.RainbowUserId = userId
	p.SocialTool = socialTool
	p.SocialToolUser = socialUser

	if err := models.GetDB().Save(&p).Error; err != nil {
		return err
	}
	return nil
}

func (d *BotActivityService) GetProjectManager(userId uint, socialType models.SocialToolType) (*models.SocialToolProjectManager, error) {
	return models.FindSocialToolProjectManager(userId, socialType)
}

func BindDiscordProjectConfig(config *models.SocialToolProjectManager, id uint) error {
	// info, err := GetDiscordGuildInfo(config.GuildId)
	// if err != nil {
	// 	return err
	// }

	// config.GuildName = info.Name
	// config.RainbowUserId = int32(id)

	// res := models.GetDB().Create(&config)
	// if res.Error != nil {
	// 	return res.Error
	// }
	return nil
}

func BindDoDoProjectConfig(config *models.SocialToolProjectManager, id uint) error {
	// info, err := GetDoDoIslandInfo(config.IslandId)
	// if err != nil {
	// 	return err
	// }

	// config.IslandName = info.IslandName
	// config.RainbowUserId = int32(id)

	// res := models.GetDB().Create(&config)
	// if res.Error != nil {
	// 	return res.Error
	// }
	return nil
}

func DiscordCustomActivityConfig(config *models.CustomActivityConfig, userId uint) error {
	token, err := middlewares.GenDiscordOpenJWTByRainbowUserId(userId, uint(config.AppId))
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

func DoDoCustomActivityConfig(config *models.CustomActivityConfig, userId uint) error {
	token, err := middlewares.GenDoDoOpenJWTByRainbowUserId(userId, uint(config.AppId))
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
