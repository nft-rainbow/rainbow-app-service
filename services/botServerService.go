package services

import (
	"context"
	"strings"

	"fmt"
	"sync"
	"time"

	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/models/enums"
	"github.com/nft-rainbow/rainbow-app-service/utils/rand"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	utils "github.com/nft-rainbow/rainbow-app-service/utils"
	"gorm.io/gorm"
)

type (
	VerifyBotServerReq struct {
		ServerId string `form:"server_id" json:"server_id" binding:"required"`
		SocialToolQueryReq
	}
	InsertBotServerReq struct {
		SocialTool       string `json:"social_tool" binding:"required,oneof=dodo discord"`
		ServerId         string `json:"server_id" binding:"required"`
		OutdatedServerId string `json:"outdated_server_id"`
		AuthCode         string `json:"auth_code" binding:"required"`
	}
	GetBotServersReq struct {
		SocialToolQueryReq
		models.Pagination
	}
)
type BotServerService struct {
	authcodes sync.Map
	bots      map[enums.SocialToolType]Bot
	botsLock  sync.Mutex
}

func NewBotServerService() (*BotServerService, error) {
	dodo, err := getSocialToolBot(enums.SOCIAL_TOOL_DODO)
	if err != nil {
		return nil, err
	}
	return &BotServerService{
		bots: map[enums.SocialToolType]Bot{
			enums.SOCIAL_TOOL_DODO: dodo,
		},
	}, nil
}

func (d *BotServerService) getBot(socialTool enums.SocialToolType) (Bot, error) {
	d.botsLock.Lock()
	defer d.botsLock.Unlock()
	if bot, ok := d.bots[socialTool]; !ok {
		return nil, errors.Errorf("not support social tool %v", socialTool)
	} else {
		return bot, nil
	}
}

func (d *BotServerService) mustGetBot(socialTool enums.SocialToolType) Bot {
	b, err := d.getBot(socialTool)
	if err != nil {
		panic(err)
	}
	return b
}

func (d *BotServerService) GetAuthcode(socialTool enums.SocialToolType, serverId string) error {
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

	msg := fmt.Sprintf("您的授权码是 %v, 有效期5分钟。请前往NFTRainbow管理后台-「添加机器人」填写此授权码, 完成NFTRainbow机器人配置。", v)
	if _, err := bot.SendDirectMessage(context.Background(), serverId, serverInfo.OwnerId, msg); err != nil {
		return err
	}

	return nil
}

func (d *BotServerService) InsertBotServer(userId uint, req InsertBotServerReq) (*models.BotServer, error) {

	socialTool, err := enums.ParseSocialToolType(req.SocialTool)
	if err != nil {
		return nil, err
	}

	code, ok := d.authcodes.Load(d.GetServerAuthCodeKey(*socialTool, req.ServerId))
	if !ok || code.(string) != req.AuthCode {
		return nil, errors.New("auth code not match")
	}

	val, err := models.FindBotServerByRawID(req.ServerId, socialTool)
	if val != nil {
		return nil, errors.New("already exists")
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	// get user social id
	bot := d.mustGetBot(*socialTool)
	serverInfo, err := d.mustGetBot(*socialTool).GetSeverInfo(context.Background(), req.ServerId)
	if err != nil {
		return nil, err
	}

	// create channel
	// 自动创建一个文字频道，该频道作为指令交互频道。频道名称：NFT领取通道，频道图片使用NFTRainbow logo；
	channelId, err := bot.CreateChannel(context.Background(), req.ServerId, "NFT领取通道", 1)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create rainbow channel")
	}

	var msgId string
	if err = utils.Retry(3, time.Millisecond*100, func() error {
		msgId, err = bot.SendChannelMessage(context.Background(), channelId, CrAllCommandsZh)
		return err
	}); err != nil {
		logrus.WithError(err).Info("failed to send help message to channel")
	}

	if msgId != "" {
		if err = utils.Retry(3, time.Millisecond*100, func() error {
			return bot.SetChannelMessageTop(context.Background(), msgId, true)
		}); err != nil {
			logrus.WithField("message id", msgId).WithError(err).Info("failed to set message to top")
		}
	}

	var p models.BotServer
	p.RainbowUserId = userId
	p.OutdatedServerId = req.OutdatedServerId
	p.SocialTool = *socialTool
	p.RawServerId = req.ServerId
	p.ServerName = serverInfo.Name
	p.OwnerSocialId = serverInfo.OwnerId
	p.DefaultActivityChannelId = channelId

	if err := models.GetDB().Save(&p).Error; err != nil {
		return nil, err
	}
	return &p, nil
}

func (d *BotServerService) GetBotServers(userId uint, queryParams *GetBotServersReq) (*models.FindBotServersResult, error) {
	socialTool, err := enums.ParseSocialToolType(queryParams.SocialTool)
	if err != nil {
		return nil, err
	}
	return models.FindBotServers(userId, socialTool, queryParams.Pagination)
}

func (d *BotServerService) GetActivitiesOfBotServers(userId uint, cond *models.FindBotServerActivitiesCond) (*models.FindBotServerActivitiesResult, error) {
	return models.FindActivitiesOfUserBotServers(userId, cond)
}

func (d *BotServerService) GetBotServer(userId uint, serverId uint) (*models.BotServer, error) {
	return VerifyServerBelongsToUser(userId, serverId)
}

func (d *BotServerService) AddPushInfo(userId uint, serverId uint, pushInfoReq PushInfoReq) (*models.PushInfo, error) {
	botServer, err := VerifyServerBelongsToUser(userId, serverId)
	if err != nil {
		return nil, err
	}

	if pushInfoReq.ChannelId == "" {
		pushInfoReq.ChannelId = botServer.DefaultActivityChannelId
	}

	exists, err := models.IsPushInfoExists(pushInfoReq.ActivityID, pushInfoReq.ChannelId)
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, errors.New("the channel has configured the activity")
	}

	var activity models.Activity
	if err := models.GetDB().Model(&models.Activity{}).Where("id=?", pushInfoReq.ActivityID).First(&activity).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	pushInfo, err := pushInfoReq.ToModel(false)
	if err != nil {
		return nil, err
	}
	pushInfo.BotServerID = botServer.ID
	pushInfo.Activity = &activity

	if err := models.GetDB().Save(pushInfo).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return pushInfo, nil
}

// send message to channel
func (d *BotServerService) Push(userId uint, channelId string, pushInfoId uint) error {
	pushInfo, err := models.FindPushInfoById(pushInfoId)
	if err != nil {
		return err
	}

	// check server belongs to user
	botServer, err := VerifyServerBelongsToUser(userId, pushInfo.BotServerID)
	if err != nil {
		return err
	}

	if channelId == "" {
		channelId = pushInfo.ChannelId
	}

	// startTime, err := time.ParseDuration(fmt.Sprintf("%ds", pushInfo.Activity.StartedTime))
	if err := d.mustGetBot(botServer.SocialTool).Push(
		channelId,
		PushData{
			Roles:         strings.Split(pushInfo.Roles, ","),
			Content:       pushInfo.Content,
			PushInfoID:    pushInfoId,
			ActivityName:  pushInfo.Activity.Name,
			StartTime:     time.Unix(pushInfo.Activity.StartedTime, 0),
			EndTime:       time.Unix(pushInfo.Activity.EndedTime, 0),
			ActivityImage: pushInfo.Activity.ActivityPictureURL,
			ClaimLink:     fmt.Sprintf("https://imdodo.com/i?gNo=%s&c=%s", botServer.OutdatedServerId, pushInfo.ChannelId),
		},
	); err != nil {
		return err
	}

	pushInfo.LastPushTime = time.Now().Unix()
	return models.GetDB().Save(pushInfo).Error
}

func (d *BotServerService) UpdatePushInfo(userId uint, serverId uint, pushInfoReq PushInfoReq) (*models.PushInfo, error) {
	if _, err := VerifyServerBelongsToUser(userId, serverId); err != nil {
		return nil, err
	}

	if _, err := VerifyPushInfoBelongsToServer(serverId, pushInfoReq.ID); err != nil {
		return nil, err
	}

	pushInfo, err := pushInfoReq.ToModel(true)
	if err != nil {
		return nil, err
	}

	if err := models.GetDB().Save(pushInfo).Error; err != nil {
		return nil, err
	}

	return pushInfo, nil
}

func (d *BotServerService) GetChannels(socialTool enums.SocialToolType, rawServerId string) ([]*Channel, error) {
	return d.mustGetBot(socialTool).GetChannels(rawServerId)
}

func (d *BotServerService) GetRoles(socialTool enums.SocialToolType, rawServerId string) ([]*Role, error) {
	return d.mustGetBot(socialTool).GetRoles(rawServerId)
}

func (d *BotServerService) GetInviteUrl(socialTool enums.SocialToolType) string {
	return d.mustGetBot(socialTool).GetInviteUrl()
}

func (d *BotServerService) GetServerAuthCodeKey(socialTool enums.SocialToolType, serverId string) string {
	return fmt.Sprintf("%s%s", serverId, d.mustGetBot(socialTool).GetSocialToolType())
}

func VerifyServerBelongsToUser(userId uint, serverId uint) (*models.BotServer, error) {
	botServer, err := models.FindBotServerById(serverId)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if botServer.RainbowUserId != userId {
		return nil, errors.New("server not belongs to user")
	}

	return botServer, nil
}

func VerifyPushInfoBelongsToServer(serverId uint, pushInfoId uint) (*models.PushInfo, error) {
	pushInfo, err := models.FindPushInfoById(pushInfoId)
	if err != nil {
		return nil, err
	}

	if pushInfo.BotServerID != serverId {
		return nil, errors.New("push info not belongs to server")
	}

	return pushInfo, nil
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
