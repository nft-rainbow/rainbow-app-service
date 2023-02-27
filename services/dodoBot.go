package services

import (
	"context"
	"fmt"
	dodo "github.com/dodo-open/dodo-open-go"
	"github.com/dodo-open/dodo-open-go/client"
	dodoClient "github.com/dodo-open/dodo-open-go/client"
	"github.com/dodo-open/dodo-open-go/model"
	"github.com/dodo-open/dodo-open-go/tools"
	"github.com/dodo-open/dodo-open-go/websocket"
	"github.com/spf13/viper"
	"strings"
	"time"
)

var instance dodoClient.Client

func InitInstance() websocket.Client {
	var err error
	instance, err = dodo.NewInstance(viper.GetString("dodoBot.clientId"), viper.GetString("dodoBot.tokenId"), client.WithTimeout(time.Second*3))
	if err != nil {
		panic(err)
	}
	handlers := &websocket.MessageHandlers{
		ChannelMessage: func(event *websocket.WSEventMessage, data *websocket.ChannelMessageEventBody) error {
			switch data.MessageType {
			case model.TextMsg:
				messageBody := &model.TextMessage{}
				if err := tools.JSON.Unmarshal(data.MessageBody, &messageBody); err != nil {
					return err
				}

				if strings.Contains(messageBody.Content, "/铸造") {
					//contents := strings.Split(messageBody.Content, " ")
					//activityId := contents[1]
					//command := contents[2]
					//
					//HandlePOAPH5Mint(&POAPRequest{})

					resp, token, contactId, err := HandleCustomMint(data.DodoId, data.ChannelId, "dodo")
					if err != nil {
						processErrorMessage(&instance, data, err.Error())
						return nil
					}

					_, _ = instance.SendChannelMessage(context.Background(), &model.SendChannelMessageReq{
						ChannelId:   data.ChannelId,
						MessageBody: &model.TextMessage{Content: fmt.Sprintf("<@!%s> Create mint task successfully. The correspding transaction hash is %s", data.DodoId, *resp.Hash)},
					})

					res, err := GenDoDoMintRes(token, resp.GetCreatedAt(), resp.GetContract(), resp.GetMintTo(), data.DodoId, data.ChannelId, resp.GetId(), contactId)
					if err != nil {
						processErrorMessage(&instance, data, err.Error())
						return nil
					}

					_, _ = instance.SendChannelMessage(context.Background(), &model.SendChannelMessageReq{
						ChannelId:   data.ChannelId,
						MessageBody: &model.TextMessage{Content: fmt.Sprintf("<@!%s> Congratulate on minting NFT for %s successfully. Check this link to view it: %s \n  %s", data.DodoId, res.UserAddress, res.NFTAddress, viper.GetString("advertise"))},
					})
					return nil
				} else if strings.Contains(messageBody.Content, "/bind CFX") {
					contents := strings.Split(messageBody.Content, " ")
					if len(contents) < 3 {
						_, _ = instance.SendChannelMessage(context.Background(), &model.SendChannelMessageReq{
							ChannelId:   data.ChannelId,
							MessageBody: &model.TextMessage{Content: fmt.Sprintf("<@!%s> %s", data.DodoId, "The input is wrong")},
						})
						return nil
					}

					userAddress := contents[2]

					err := HandleBindCfxAddress(data.DodoId, userAddress, "dodo")
					if err != nil {
						processErrorMessage(&instance, data, err.Error())
						return nil
					}

					_, _ = instance.SendChannelMessage(context.Background(), &model.SendChannelMessageReq{
						ChannelId:   data.ChannelId,
						MessageBody: &model.TextMessage{Content: fmt.Sprintf("<@!%s> success!", data.DodoId)},
					})
					return nil

				} else if strings.Contains(messageBody.Content, "/address CFX") {
					resp, err := GetDoDoBindCFXAddress(data.DodoId)
					if err != nil {
						processErrorMessage(&instance, data, err.Error())
						return nil
					}
					_, _ = instance.SendChannelMessage(context.Background(), &model.SendChannelMessageReq{
						ChannelId:   data.ChannelId,
						MessageBody: &model.TextMessage{Content: fmt.Sprintf("<@!%s> %s", data.DodoId, resp)},
					})
					return nil
				}
			}
			return nil
		},
	}

	ws, err := websocket.New(instance,
		websocket.WithMessageQueueSize(128),
		websocket.WithMessageHandlers(handlers),
	)
	if err != nil {
		panic(err)
	}

	return ws
}

func processErrorMessage(instance *client.Client, data *websocket.ChannelMessageEventBody, message string) {
	_, _ = (*instance).SendChannelMessage(context.Background(), &model.SendChannelMessageReq{
		ChannelId:   data.ChannelId,
		MessageBody: &model.TextMessage{Content: fmt.Sprintf("<@!%s> %s", data.DodoId, message)},
	})
}

func GetInstance() *dodoClient.Client {
	return &instance
}
