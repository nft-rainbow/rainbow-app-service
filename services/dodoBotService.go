package services

import (
	"context"

	dodoClient "github.com/dodo-open/dodo-open-go/client"
	"github.com/dodo-open/dodo-open-go/model"
	"github.com/dodo-open/dodo-open-go/websocket"
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

func (d *DodoBot) SendChannelMessage(ctx context.Context, channedId string, msg string) error {
	_, err := d.instance.SendChannelMessage(context.Background(), &model.SendChannelMessageReq{
		ChannelId:   channedId,
		MessageBody: &model.TextMessage{Content: msg},
	})
	return err
}

func (d *DodoBot) SendDirectMessage(ctx context.Context, userId string, msg string) error {
	_, err := d.instance.SendDirectMessage(ctx, &model.SendDirectMessageReq{
		DodoId:      userId,
		MessageBody: &model.TextMessage{Content: msg},
	})
	return err
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
