package services

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	dodoClient "github.com/dodo-open/dodo-open-go/client"
	"github.com/dodo-open/dodo-open-go/model"
	"github.com/dodo-open/dodo-open-go/websocket"
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/sirupsen/logrus"
)

type DodoBot struct {
	ws       websocket.Client
	instance dodoClient.Client
}

func NewDodoBot() *DodoBot {
	_ws, _instance := InitDodoInstance()
	return &DodoBot{
		ws:       _ws,
		instance: _instance,
	}
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

	_, err = instance.SendChannelMessage(context.Background(), &model.SendChannelMessageReq{
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

func (d *DodoBot) ListenWebsocket() {
	logrus.Info("Start to connect")

	err := d.ws.Connect()
	if err != nil {
		panic(err)
	}
	logrus.Info("Start to listen")

	err = d.ws.Listen()
	if err != nil {
		panic(err)
	}
}
