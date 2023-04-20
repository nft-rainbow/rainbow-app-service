package services

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	dodo "github.com/dodo-open/dodo-open-go"
	dodoClient "github.com/dodo-open/dodo-open-go/client"
	"github.com/dodo-open/dodo-open-go/model"
	"github.com/dodo-open/dodo-open-go/tools"
	"github.com/dodo-open/dodo-open-go/websocket"
	jsoniter "github.com/json-iterator/go"
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type DodoBot struct {
	instance        dodoClient.Client
	instanceBotInfo *model.GetBotInfoRsp
	commander       *DodoBotCommander
}

func NewDodoBot(clientId, tokenId string) (*DodoBot, error) {
	_instance, err := dodo.NewInstance(clientId, tokenId, dodoClient.WithTimeout(time.Second*3))
	if err != nil {
		return nil, err
	}

	botInfo, err := _instance.GetBotInfo(context.Background())
	if err != nil {
		return nil, err
	}

	b := &DodoBot{
		instance:        _instance,
		instanceBotInfo: botInfo,
	}
	b.commander = NewDodoBotCommander(b)

	go b.ListenWebsocket()
	return b, nil
}

func (d *DodoBot) GetSocialToolType() models.SocialToolType {
	return models.SOCIAL_TOOL_DODO
}

func (d *DodoBot) SendChannelMessage(ctx context.Context, channedId string, msg string) error {
	_, err := d.instance.SendChannelMessage(context.Background(), &model.SendChannelMessageReq{
		ChannelId:   channedId,
		MessageBody: &model.TextMessage{Content: msg},
	})
	return err
}

func (d *DodoBot) SendDirectMessage(ctx context.Context, serverId string, userId string, msg string) error {
	_, err := d.instance.SendDirectMessage(ctx, &model.SendDirectMessageReq{
		IslandSourceId: serverId,
		DodoSourceId:   userId,
		MessageBody:    &model.TextMessage{Content: msg},
	})
	return err
}

func (d *DodoBot) GetSeverInfo(ctx context.Context, serverId string) (*SeverInfo, error) {
	info, err := d.instance.GetIslandInfo(ctx, &model.GetIslandInfoReq{
		IslandSourceId: serverId,
	})
	if err != nil {
		return nil, err
	}
	return &SeverInfo{
		OwnerId: info.OwnerDodoSourceId,
	}, nil
}

func (d *DodoBot) GetChannels(serverId string) ([]*Channel, error) {
	_channels, err := d.instance.GetChannelList(context.Background(), &model.GetChannelListReq{
		IslandSourceId: serverId,
	})
	if err != nil {
		return nil, err
	}

	channels := []*Channel{}
	for _, v := range _channels {
		channels = append(channels, &Channel{
			ChannelId:   v.ChannelId,
			ChannelName: v.ChannelName,
		})
	}

	return channels, nil
}

func (d *DodoBot) Push(channelId string, roles []string, name, activityId, content, color string) error {
	var card = `{
		"content": "",
		"card": {
		  "type": "card",
		  "components": [
			{
			  "type": "section",
			  "text": {
				"type": "dodo-md",
				"content": "{roles} {name}#{activity} 来了！\n在频道中发送【教程】，机器人将私信你领取教程"
			  }
			},
			{
			  "type": "section",
			  "text": {
				"type": "dodo-md",
				"content": "{content}"
			  }
			}
		  ],
		  "theme": "{color}",
		  "title": "新活动发布啦！"
		}
	  }`
	var message model.CardMessage

	_roles := ""
	if len(roles) == 1 && roles[0] == "all" {
		_roles = "<@all>"
	} else {
		for _, v := range roles {
			_roles += fmt.Sprintf("<@&%s>", v)
		}
	}

	card = strings.Replace(card, "{roles}", _roles, -1)
	card = strings.Replace(card, "{name}", name, -1)
	card = strings.Replace(card, "{activity}", activityId, -1)
	card = strings.Replace(card, "{content}", content, -1)
	card = strings.Replace(card, "{color}", color, -1)
	err := json.Unmarshal([]byte(card), &message)
	if err != nil {
		return err
	}

	_, err = d.instance.SendChannelMessage(context.Background(), &model.SendChannelMessageReq{
		ChannelId:   channelId,
		MessageBody: &message,
	})
	return err
}

func (d *DodoBot) GetRoles(serverId string) ([]*Role, error) {
	_roles, err := d.instance.GetRoleList(context.Background(), &model.GetRoleListReq{
		IslandSourceId: serverId,
	})
	if err != nil {
		return nil, err
	}

	roles := []*Role{}
	for _, v := range _roles {
		roles = append(roles, &Role{
			RoleId:   v.RoleId,
			RoleName: v.RoleName,
		})
	}

	return roles, nil
}

func (d *DodoBot) GetInviteUrl() string {
	return viper.GetString("dodoBot.inviteUrl")
}

func (d *DodoBot) RunCommand(channelId string, userDodoSourceId string, command string) error {
	return d.commander.ExcuteCommand(channelId, userDodoSourceId, command)
}

func (d *DodoBot) ListenWebsocket() {
	logrus.Info("Start to connect dodo websocket")
	handlers := &websocket.MessageHandlers{ChannelMessage: d.dodoChannelMsgHandler}

	ws, err := websocket.New(d.instance,
		websocket.WithMessageQueueSize(128),
		websocket.WithMessageHandlers(handlers),
	)
	if err != nil {
		panic(err)
	}

	if err = ws.Connect(); err != nil {
		panic(err)
	}
	logrus.Info("Start to listen dodo websocket")

	err = ws.Listen()
	if err != nil {
		panic(err)
	}
}

func (d *DodoBot) dodoChannelMsgHandler(event *websocket.WSEventMessage, data *websocket.ChannelMessageEventBody) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	j, _ := json.Marshal(data)
	logrus.WithField("msg", string(j)).Info("got message")

	if data.MessageType != model.TextMsg {
		return nil
	}

	messageBody := &model.TextMessage{}
	if err := tools.JSON.Unmarshal(data.MessageBody, &messageBody); err != nil {
		return err
	}

	isCommand := len(messageBody.Content) > 0 && messageBody.Content[0] == byte('/')
	if !isCommand {
		return nil
	}
	logrus.WithField("command", messageBody.Content).Info("got command")

	return d.RunCommand(data.ChannelId, data.DodoSourceId, messageBody.Content)
}
