package services

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/nft-rainbow/rainbow-app-service/models/enums"
	"github.com/spf13/viper"
)

type (
	Bot interface {
		SendChannelMessage(ctx context.Context, channedId string, msg string, referMsgId ...string) error
		SendDirectMessage(ctx context.Context, serverId string, userId string, msg string) error
		GetSeverInfo(ctx context.Context, serverId string) (*SeverInfo, error)
		GetSocialToolType() enums.SocialToolType
		GetChannels(serverId string) ([]*Channel, error)
		GetRoles(serverId string) ([]*Role, error)
		GetInviteUrl() string
		Push(channelId string, pushData PushData) error
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

	ChannelMsgSource struct {
		serverId         string
		channelId        string
		userDodoSourceId string
		messageId        string
	}

	PushData struct {
		Roles         []string
		Content       string
		PushInfoID    uint
		ActivityName  string
		StartTime     time.Time
		EndTime       time.Time
		ActivityImage string
		ClaimLink     string
	}
)

var (
	dodoBot           *DodoBot
	dodoBotCreateOnce sync.Once
)

func getSocialToolBot(socialToolType enums.SocialToolType) (Bot, error) {
	switch socialToolType {
	case enums.SOCIAL_TOOL_DODO:
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
