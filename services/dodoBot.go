package services

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	dodo "github.com/dodo-open/dodo-open-go"
	dodoClient "github.com/dodo-open/dodo-open-go/client"
	"github.com/dodo-open/dodo-open-go/model"
	"github.com/dodo-open/dodo-open-go/tools"
	"github.com/dodo-open/dodo-open-go/websocket"
	jsoniter "github.com/json-iterator/go"
	"github.com/nft-rainbow/rainbow-app-service/models/enums"
	"github.com/nft-rainbow/rainbow-app-service/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	pushTemplate = `{
		"content": "",
		"card": {
		  "type": "card",
		  "components": [
			{
			  "type": "section",
			  "text": {
				"type": "dodo-md",
				"content": "{roles} {name}#{activity} 来了！\n在频道中发送【/教程】，机器人将私信你领取教程"
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

func (d *DodoBot) GetSocialToolType() enums.SocialToolType {
	return enums.SOCIAL_TOOL_DODO
}

func (d *DodoBot) CreateChannel(ctx context.Context, serverId string, channelName string, channelType int) (string, error) {
	resp, err := d.instance.CreateChannel(ctx, &model.CreateChannelReq{
		IslandSourceId: serverId,
		ChannelName:    channelName,
		ChannelType:    model.ChannelType(channelType),
	})
	if err != nil {
		return "", err
	}
	return resp.ChannelId, nil
}

func (d *DodoBot) SendChannelMessage(ctx context.Context, channedId string, msg string, referMsgId ...string) (string, error) {
	if len(referMsgId) == 0 {
		referMsgId = append(referMsgId, "")
	}
	resp, err := d.instance.SendChannelMessage(context.Background(), &model.SendChannelMessageReq{
		ChannelId:           channedId,
		MessageBody:         &model.TextMessage{Content: msg},
		ReferencedMessageId: referMsgId[0],
	})
	if err != nil {
		return "", err
	}
	return resp.MessageId, nil
}

func (d *DodoBot) SendDirectMessage(ctx context.Context, serverId string, userId string, msg string) (string, error) {
	resp, err := d.instance.SendDirectMessage(ctx, &model.SendDirectMessageReq{
		IslandSourceId: serverId,
		DodoSourceId:   userId,
		MessageBody:    &model.TextMessage{Content: msg},
	})
	if err != nil {
		return "", err
	}
	return resp.MessageId, nil
}

func (d *DodoBot) SetChannelMessageTop(ctx context.Context, messageId string, setTop bool) error {
	_, err := d.instance.SetChannelMessageTop(ctx, &model.SetChannelMessageTopReq{
		MessageId:   messageId,
		OperateType: utils.Bool2Int(setTop),
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
		Name:    info.IslandName,
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

func (d *DodoBot) Push(channelId string, pushData PushData) error {

	_roles := ""
	if len(pushData.Roles) == 0 || pushData.Roles[0] == "" {
		_roles = "<@all>"
	} else {
		_roles = "<@&" + strings.Join(pushData.Roles, "><@&") + ">"
	}

	pushDataForTemplate := struct {
		PushData
		Roles               string
		StartTime           string
		EndTime             string
		StartTimeInMillisec int64
		CountdownStyle      string
	}{
		PushData:            pushData,
		Roles:               _roles,
		StartTime:           pushData.StartTime.Format("2006-01-02 15:04:05"),
		EndTime:             pushData.EndTime.Format("2006-01-02 15:04:05"),
		StartTimeInMillisec: pushData.StartTime.UnixMilli(),
		// 小于1天也按day style倒计时
		CountdownStyle: "day", //"hour",
	}
	if pushData.StartTime.Before(time.Unix(1, 0)) {
		pushDataForTemplate.StartTime = "无"
	}
	if pushData.EndTime.Before(time.Unix(1, 0)) {
		pushDataForTemplate.EndTime = "无"
	}

	countDown := time.Until(pushData.StartTime)
	if countDown > time.Hour*24 {
		pushDataForTemplate.CountdownStyle = "day"
	}

	pushMsgInJson := ExcuteTemplate(CrPushJsonTemplate, pushDataForTemplate)
	var message model.CardMessage
	err := json.Unmarshal([]byte(pushMsgInJson), &message)
	if err != nil {
		return err
	}

	if pushData.StartTime.Before(time.Now()) {
		message.Card.Components = append(message.Card.Components[:2], message.Card.Components[3:]...)
	} else {
		message.Card.Components = append(message.Card.Components[:4], message.Card.Components[5:]...)
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

func (d *DodoBot) RunCommand(channelMsgSource ChannelMsgSource, command string) error {
	return d.commander.ExcuteCommand(channelMsgSource, command)
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

	channelMsgSource := ChannelMsgSource{
		serverId:         data.IslandSourceId,
		channelId:        data.ChannelId,
		userDodoSourceId: data.DodoSourceId,
		messageId:        data.MessageId,
	}
	return d.RunCommand(channelMsgSource, messageBody.Content)
}
