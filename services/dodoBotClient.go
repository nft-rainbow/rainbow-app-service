package services

import (
	"context"

	dodoClient "github.com/dodo-open/dodo-open-go/client"
	"github.com/dodo-open/dodo-open-go/model"
)

type DodoBotClient struct {
	raw dodoClient.Client
}

// 发送channel消息
func (d *DodoBotClient) SendChannelMessage(ctx context.Context, channedId string, msg string) error {
	_, err := d.raw.SendChannelMessage(ctx, &model.SendChannelMessageReq{
		ChannelId:   channedId,
		MessageBody: &model.TextMessage{Content: msg},
	})
	return err
}

// 发送私信
func (d *DodoBotClient) SendDirectMessage(ctx context.Context, userId string, msg string) error {
	_, err := d.raw.SendDirectMessage(context.Background(), &model.SendDirectMessageReq{
		DodoId:      userId,
		MessageBody: &model.TextMessage{Content: msg},
	})
	return err
}
