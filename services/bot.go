package services

import (
	"context"
	"errors"
	"sync"

	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/spf13/viper"
)

type (
	Bot interface {
		SendChannelMessage(ctx context.Context, channedId string, msg string) error
		SendDirectMessage(ctx context.Context, serverId string, userId string, msg string) error
		GetSeverInfo(ctx context.Context, serverId string) (*SeverInfo, error)
		GetSocialToolType() models.SocialToolType
		GetChannels(serverId string) ([]*Channel, error)
		GetRoles(serverId string) ([]*Role, error)
		Push(channelId string, roles []string, name, activityId, content, color string) error
	}

	BotCommander interface {
		Mint() error
		Bind() error
	}
)

type (
	SeverInfo struct {
		Name    string `json:"name"`
		OwnerId string `json:"owner"`
	}

	Channel struct {
		ChannelId   string `json:"channelId"`   // 频道号
		ChannelName string `json:"channelName"` // 频道名称
	}

	Role struct {
		RoleId   string `json:"roleId"`   // 身份组ID
		RoleName string `json:"roleName"` // 身份组名称
	}
)

var (
	dodoBot           *DodoBot
	dodoBotCreateOnce sync.Once
)

func getSocialToolBot(socialToolType models.SocialToolType) (Bot, error) {
	switch socialToolType {
	case models.SOCIAL_TOOL_DODO:
		var err error
		dodoBotCreateOnce.Do(func() {
			clientId := viper.GetString("dodoBot.clientId")
			tokenId := viper.GetString("dodoBot.tokenId")
			dodoBot, err = NewDodoBot(clientId, tokenId)
		})
		return dodoBot, err
	}

	return nil, errors.New("unsupported social tool")
}
