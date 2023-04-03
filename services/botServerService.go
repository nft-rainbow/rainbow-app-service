package services

import (
	"context"
	"strings"

	"fmt"
	"sync"
	"time"

	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/utils/rand"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type BotServerService struct {
	authcodes sync.Map
	bots      map[models.SocialToolType]Bot
	botsLock  sync.Mutex
}

func NewBotServerService() (*BotServerService, error) {
	dodo, err := getSocialToolBot(models.SOCIAL_TOOL_DODO)
	if err != nil {
		return nil, err
	}
	return &BotServerService{
		bots: map[models.SocialToolType]Bot{
			models.SOCIAL_TOOL_DODO: dodo,
		},
	}, nil
}

func (d *BotServerService) getBot(socialTool models.SocialToolType) (Bot, error) {
	d.botsLock.Lock()
	defer d.botsLock.Unlock()
	if bot, ok := d.bots[socialTool]; !ok {
		return nil, errors.Errorf("not support social tool %v", socialTool)
	} else {
		return bot, nil
	}
}

func (d *BotServerService) mustGetBot(socialTool models.SocialToolType) Bot {
	b, err := d.getBot(socialTool)
	if err != nil {
		panic(err)
	}
	return b
}

func (d *BotServerService) VerifyBotServer(socialTool models.SocialToolType, serverId string) error {
	authcodeKey := d.GetServerAuthCodeKey(socialTool, serverId)
	v, loaded := d.authcodes.LoadOrStore(authcodeKey, rand.NumString(6))
	if !loaded {
		go func() {
			<-time.After(time.Minute * 5)
			d.authcodes.Delete(authcodeKey)
		}()
	}

	// get owner id from group info
	bot := d.mustGetBot(socialTool)
	serverInfo, err := bot.GetSeverInfo(context.Background(), serverId)
	if err != nil {
		return err
	}

	msg := fmt.Sprintf("You are setting Rainbow-Bot, your Rainbow auth code is %v, please fill back to Rainbow to complete authentication.", v)
	if err := bot.SendDirectMessage(context.Background(), serverId, serverInfo.OwnerId, msg); err != nil {
		return err
	}

	return nil
}

func (d *BotServerService) InsertBotServer(userId uint, req InsertSocialServerReq) error {
	code, ok := d.authcodes.Load(d.GetServerAuthCodeKey(req.SocialTool, req.ServerId))
	if !ok || code.(string) != req.AuthCode {
		return errors.New("auth code not match")
	}

	val, err := models.FindBotServerByRawID(req.ServerId, &req.SocialTool)
	if val != nil {
		return errors.New("already exists")
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	// get user social id
	serverInfo, err := d.mustGetBot(req.SocialTool).GetSeverInfo(context.Background(), req.ServerId)
	if err != nil {
		return err
	}

	var p models.BotServer
	p.RainbowUserId = userId
	p.SocialTool = req.SocialTool
	p.RawServerId = req.ServerId
	p.OwnerSocialId = serverInfo.OwnerId

	if err := models.GetDB().Save(&p).Error; err != nil {
		return err
	}
	return nil
}

func (d *BotServerService) GetBotServers(userId uint, socialType *models.SocialToolType) ([]*models.BotServer, error) {
	return models.FindBotServers(userId, socialType)
}

func (d *BotServerService) GetBotServer(userId uint, serverId uint) (*models.BotServer, error) {
	return VerifyServerBelongsToUser(userId, serverId)
}

func (d *BotServerService) AddActivity(userId uint, serverId uint, pushInfo PushInfoReq) (*models.PushInfo, error) {
	// check server belongs to user
	botServer, err := VerifyServerBelongsToUser(userId, serverId)
	if err != nil {
		return nil, err
	}

	if botServer.PushInfo != nil {
		return nil, errors.New("already exist")
	}

	var activity models.Activity
	if err := models.GetDB().Model(&models.Activity{}).Where("id=?", pushInfo.ActivityID).First(&activity).Error; err != nil {
		return nil, err
	}

	botServer.PushInfo, err = pushInfo.ToModel()
	if err != nil {
		return nil, err
	}

	if err := models.GetDB().Save(botServer).Error; err != nil {
		return nil, err
	}
	return botServer.PushInfo, nil
}

// send message to channel
func (d *BotServerService) Push(userId uint, pushInfoId uint) error {
	var pushInfo models.PushInfo
	if err := models.GetDB().Preload("Activity").Where("id=?", pushInfoId).Find(&pushInfo).Error; err != nil {
		return err
	}

	// check server belongs to user
	botServer, err := VerifyServerBelongsToUser(userId, pushInfo.BotServerID)
	if err != nil {
		return err
	}

	activity := pushInfo.Activity
	return d.mustGetBot(botServer.SocialTool).Push(
		pushInfo.ChannelId,
		strings.Split(pushInfo.Roles, ","),
		activity.AppName,
		activity.ActivityID,
		pushInfo.Content,
		pushInfo.ColorTheme)
}

func (d *BotServerService) UpdateActivity(userId uint, serverId uint, pushInfo PushInfoReq) (*models.PushInfo, error) {
	// check server belongs to user
	if _, err := VerifyServerBelongsToUser(userId, serverId); err != nil {
		return nil, err
	}

	if err := models.GetDB().Save(&pushInfo).Error; err != nil {
		return nil, err
	}
	return nil, errors.New("not implemented")
}

func (d *BotServerService) GetChannels(socialTool models.SocialToolType, rawServerId string) ([]*Channel, error) {
	return d.mustGetBot(socialTool).GetChannels(rawServerId)
}

func (d *BotServerService) GetRoles(socialTool models.SocialToolType, rawServerId string) ([]*Role, error) {
	return d.mustGetBot(socialTool).GetRoles(rawServerId)
}

func (d *BotServerService) GetServerAuthCodeKey(socialTool models.SocialToolType, serverId string) string {
	return fmt.Sprintf("%s%s", serverId, d.mustGetBot(socialTool).GetSocialToolType())
}

func VerifyServerBelongsToUser(userId uint, serverId uint) (*models.BotServer, error) {
	s := models.BotServer{RainbowUserId: userId}
	s.ID = serverId
	if err := models.GetDB().Where(&s).First(&s).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("server not belongs to user")
		}
		return nil, err
	}
	return &s, nil
}

// func BindDiscordProjectConfig(config *models.SocialToolServer, id uint) error {
// 	// info, err := GetDiscordGuildInfo(config.GuildId)
// 	// if err != nil {
// 	// 	return err
// 	// }

// 	// config.GuildName = info.Name
// 	// config.RainbowUserId = int32(id)

// 	// res := models.GetDB().Create(&config)
// 	// if res.Error != nil {
// 	// 	return res.Error
// 	// }
// 	return nil
// }

// func BindDoDoProjectConfig(config *models.SocialToolServer, id uint) error {
// 	// info, err := GetDoDoIslandInfo(config.IslandId)
// 	// if err != nil {
// 	// 	return err
// 	// }

// 	// config.IslandName = info.IslandName
// 	// config.RainbowUserId = int32(id)

// 	// res := models.GetDB().Create(&config)
// 	// if res.Error != nil {
// 	// 	return res.Error
// 	// }
// 	return nil
// }

// func DiscordCustomActivityConfig(config *models.CustomActivityConfig, userId uint) error {
// 	token, err := middlewares.GenDiscordOpenJWTByRainbowUserId(userId, uint(config.AppId))
// 	if err != nil {
// 		return err
// 	}
// 	info, err := GetContractInfo(config.ContractID, token)
// 	if err != nil {
// 		return err
// 	}
// 	config.ContractType = *info.Type
// 	config.Chain = *info.ChainType
// 	config.AppId = uint(*info.AppId)
// 	config.ContractAddress = *info.Address

// 	res := models.GetDB().Create(&config)
// 	if res.Error != nil {
// 		return res.Error
// 	}

// 	return nil
// }

// func DoDoCustomActivityConfig(config *models.CustomActivityConfig, userId uint) error {
// 	token, err := middlewares.GenDoDoOpenJWTByRainbowUserId(userId, uint(config.AppId))
// 	if err != nil {
// 		return err
// 	}
// info, err := GetContractInfo(config.ContractID, token)
// 	if err != nil {
// 		return err
// 	}
// 	config.ContractType = *info.Type
// 	config.Chain = *info.ChainType
// 	config.AppId = uint(*info.AppId)
// 	config.ContractAddress = *info.Address

// 	res := models.GetDB().Create(&config)
// 	if res.Error != nil {
// 		return res.Error
// 	}

// 	return nil
// }

// func HandleCustomMint(userId, channelId, platform string) (*openapiclient.ModelsMintTask, string, int32, error) {
// 	req := models.CustomMintReq{
// 		UserID:    userId,
// 		ChannelID: channelId,
// 	}
// 	if platform == "dodo" {
// 		resp, token, contractId, err := dodoCustomMint(&req)
// 		return resp, token, contractId, err
// 	} else if platform == "discord" {
// 		resp, token, contractId, err := discordCustomMint(&req)
// 		return resp, token, contractId, err
// 	} else {
// 		return nil, "", 0, nil
// 	}
// }

// func dodoCustomMint(req *models.CustomMintReq) (*openapiclient.ModelsMintTask, string, int32, error) {
// 	config, err := models.FindDoDoCustomActivityConfigByChannelId(req.ChannelID)
// 	if err != nil {
// 		return nil, "", 0, err
// 	}

// 	ok, err := models.CheckDoDoCustomCount(req.UserID, req.ChannelID, config.MaxMintCount)
// 	if err != nil {
// 		return nil, "", 0, err
// 	}
// 	if !ok {
// 		return nil, "", 0, errors.New("This number of the NFTs the account minted has reached the maximum")
// 	}

// 	cfxAddress, err := GetDoDoBindCFXAddress(req.UserID)
// 	if err != nil {
// 		return nil, "", 0, err
// 	}

// 	token, _ := middlewares.GenerateDoDoOpenJWT(req.ChannelID)
// 	chain, err := utils.ChainById(uint(config.Chain))
// 	if err != nil {
// 		return nil, "", 0, err
// 	}

// 	resp, err := sendCustomMintRequest(middlewares.PrefixToken(token), openapiclient.ServicesCustomMintDto{
// 		Chain:           chain,
// 		ContractAddress: config.ContractAddress,
// 		MintToAddress:   cfxAddress,
// 		MetadataUri:     &config.MetadataURI,
// 	})
// 	if err != nil {
// 		return nil, "", 0, err
// 	}

// 	return resp, token, config.ContractID, err
// }

// func discordCustomMint(req *models.CustomMintReq) (*openapiclient.ModelsMintTask, string, int32, error) {
// 	config, err := models.FindDiscordCustomActivityConfigByChannelId(req.ChannelID)
// 	if err != nil {
// 		return nil, "", 0, err
// 	}

// 	ok, err := models.CheckDiscordCustomCount(req.UserID, req.ChannelID, config.MaxMintCount)
// 	if err != nil {
// 		return nil, "", 0, err
// 	}
// 	if !ok {
// 		return nil, "", 0, errors.New("This number of the NFTs the account minted has reached the maximum")
// 	}

// 	cfxAddress, err := GetDiscordBindCFXAddress(req.UserID)
// 	if err != nil {
// 		return nil, "", 0, err
// 	}

// 	token, _ := middlewares.GenerateDiscordOpenJWT(req.ChannelID)

// 	chain, err := utils.ChainById(uint(config.Chain))
// 	if err != nil {
// 		return nil, "", 0, err
// 	}

// 	resp, err := sendCustomMintRequest(middlewares.PrefixToken(token), openapiclient.ServicesCustomMintDto{
// 		Chain:           chain,
// 		ContractAddress: config.ContractAddress,
// 		MintToAddress:   cfxAddress,
// 		MetadataUri:     &config.MetadataURI,
// 	})
// 	if err != nil {
// 		return nil, "", 0, err
// 	}

// 	return resp, token, config.ContractID, err
// }
