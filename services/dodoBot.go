package services

// import (
// 	"context"
// 	"fmt"
// 	"strings"
// 	"time"

// 	dodo "github.com/dodo-open/dodo-open-go"
// 	"github.com/dodo-open/dodo-open-go/client"
// 	dodoClient "github.com/dodo-open/dodo-open-go/client"
// 	"github.com/dodo-open/dodo-open-go/model"
// 	"github.com/dodo-open/dodo-open-go/tools"
// 	"github.com/dodo-open/dodo-open-go/websocket"
// 	"github.com/nft-rainbow/rainbow-app-service/models"
// 	"github.com/nft-rainbow/rainbow-app-service/utils"
// 	"github.com/spf13/viper"
// )

// var instance dodoClient.Client

// func InitDodoInstance() (websocket.Client, dodoClient.Client) {
// 	var err error
// 	instance, err = dodo.NewInstance(viper.GetString("dodoBot.clientId"), viper.GetString("dodoBot.tokenId"), client.WithTimeout(time.Second*3))
// 	if err != nil {
// 		panic(err)
// 	}
// 	bot, _ := instance.GetBotInfo(context.Background())
// 	handlers := &websocket.MessageHandlers{
// 		ChannelMessage: func(event *websocket.WSEventMessage, data *websocket.ChannelMessageEventBody) error {
// 			push, _ := models.FindBotServerByChannel(data.IslandSourceId)
// 			if push.PushInfos[0].ChannelId != data.ChannelId {
// 				return nil
// 			}
// 			switch data.MessageType {
// 			case model.TextMsg:
// 				messageBody := &model.TextMessage{}
// 				if err := tools.JSON.Unmarshal(data.MessageBody, &messageBody); err != nil {
// 					return err
// 				}
// 				if !strings.HasPrefix(messageBody.Content, fmt.Sprintf("<@!%v>", bot.DodoSourceId)) {
// 					return nil
// 				}
// 				commands := strings.Split(messageBody.Content, " ")

// 				if strings.HasPrefix(commands[1], "铸造") {
// 					contents := strings.Split(commands[1], "/")
// 					var activityId, command string
// 					if len(contents) == 3 {
// 						command = contents[2]
// 					}
// 					activityId = contents[1]
// 					bind, err := models.FindSocialUserConfig(data.DodoSourceId, utils.DoDo)
// 					if err != nil {
// 						processErrorMessage(&instance, data, err.Error())
// 						return nil
// 					}

// 					config, err := models.FindPOAPActivityConfigById(activityId)
// 					if err != nil {
// 						processErrorMessage(&instance, data, err.Error())
// 						return nil
// 					}

// 					// if err := config.CheckActivityValid(); err != nil {
// 					// 	processErrorMessage(&instance, data, err.Error())
// 					// 	return nil
// 					// }

// 					// err = checkSocialLimit(data.IslandSourceId, data.DodoSourceId, config.ActivityID, utils.DoDo)
// 					// if err != nil {
// 					// 	processErrorMessage(&instance, data, err.Error())
// 					// 	return nil
// 					// }
// 					_, _ = instance.SendChannelMessage(context.Background(), &model.SendChannelMessageReq{
// 						ChannelId:   data.ChannelId,
// 						MessageBody: &model.TextMessage{Content: fmt.Sprintf("<@!%s> Start to mint NFT. Please wait patiently.", data.DodoSourceId)},
// 					})

// 					res, err := HandlePOAPH5Mint(&MintReq{
// 						ActivityID:  config.ActivityCode,
// 						UserAddress: bind.CFXAddress,
// 						Command:     command,
// 					})
// 					if err != nil {
// 						processErrorMessage(&instance, data, err.Error())
// 						return nil
// 					}

// 					for {
// 						resp, _ := models.FindPOAPResultById(config.ActivityCode, int(res.ID))
// 						if resp.Hash == "" {
// 							time.Sleep(time.Second)
// 							continue
// 						}
// 						_, _ = instance.SendChannelMessage(context.Background(), &model.SendChannelMessageReq{
// 							ChannelId:   data.ChannelId,
// 							MessageBody: &model.TextMessage{Content: fmt.Sprintf("<@!%s> Mint NFT successfully. The correspding transaction hash is %v", data.DodoSourceId, resp.Hash)},
// 						})
// 						break
// 					}
// 					return nil
// 				} else if strings.HasPrefix(commands[1], "绑定") {
// 					contents := strings.Split(commands[1], "/")
// 					if len(contents) > 3 {
// 						_, _ = instance.SendChannelMessage(context.Background(), &model.SendChannelMessageReq{
// 							ChannelId:   data.ChannelId,
// 							MessageBody: &model.TextMessage{Content: fmt.Sprintf("<@!%s> %s", data.DodoSourceId, "The input is wrong")},
// 						})
// 						return nil
// 					}

// 					userAddress := contents[1]

// 					_, _, err := BindCfxAddress(data.DodoSourceId, userAddress, models.SOCIAL_TOOL_DISCORD)
// 					if err != nil {
// 						processErrorMessage(&instance, data, err.Error())
// 						return nil
// 					}

// 					_, _ = instance.SendChannelMessage(context.Background(), &model.SendChannelMessageReq{
// 						ChannelId:   data.ChannelId,
// 						MessageBody: &model.TextMessage{Content: fmt.Sprintf("<@!%s> success!", data.DodoSourceId)},
// 					})
// 					return nil

// 				} else if strings.HasPrefix(commands[1], "查地址") {
// 					resp, err := models.FindSocialUserConfig(data.DodoSourceId, models.SOCIAL_TOOL_DODO)
// 					if err != nil {
// 						processErrorMessage(&instance, data, err.Error())
// 						return nil
// 					}
// 					_, _ = instance.SendChannelMessage(context.Background(), &model.SendChannelMessageReq{
// 						ChannelId:   data.ChannelId,
// 						MessageBody: &model.TextMessage{Content: fmt.Sprintf("<@!%s> %s", data.DodoSourceId, resp.CFXAddress)},
// 					})
// 					return nil
// 				} else if strings.HasPrefix(commands[1], "教程") {
// 					_, _ = instance.SendChannelMessage(context.Background(), &model.SendChannelMessageReq{
// 						ChannelId:   data.ChannelId,
// 						MessageBody: &model.TextMessage{Content: fmt.Sprintf("<@!%s> %s", data.DodoSourceId, guide)},
// 					})
// 				} else if strings.HasPrefix(commands[1], "创建") {
// 					_, _ = instance.SendChannelMessage(context.Background(), &model.SendChannelMessageReq{
// 						ChannelId:   data.ChannelId,
// 						MessageBody: &model.TextMessage{Content: fmt.Sprintf("<@!%s> %s", data.DodoSourceId, anywebH5)},
// 					})
// 				} else if strings.HasPrefix(commands[1], "查口令") {
// 					contents := strings.Split(commands[1], "/")
// 					activity := contents[1]
// 					config, err := models.FindPOAPActivityConfigById(activity)
// 					if err != nil {
// 						processErrorMessage(&instance, data, err.Error())
// 						return nil
// 					}
// 					if config.Command == "" {
// 						processErrorMessage(&instance, data, fmt.Sprintf("<@!%s> The command is not needed in this activity", data.DodoSourceId))
// 						return nil
// 					}
// 					instance.SendDirectMessage(context.Background(), &model.SendDirectMessageReq{
// 						DodoSourceId: data.DodoSourceId,
// 						MessageBody:  &model.TextMessage{Content: config.Command},
// 					})
// 				}
// 			}
// 			return nil
// 		},
// 	}

// 	ws, err := websocket.New(instance,
// 		websocket.WithMessageQueueSize(128),
// 		websocket.WithMessageHandlers(handlers),
// 	)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return ws, instance
// }

// // func DoDoPushActivity(req *PushReq) (*model.SendChannelMessageRsp, error) {
// // 	config, err := models.FindPOAPActivityConfigById(req.ActivityId)
// // 	if err != nil {
// // 		return nil, err
// // 	}

// // 	if err := config.CheckActivityValid(); err != nil {
// // 		return nil, err
// // 	}

// // 	var message model.CardMessage

// // 	roles := ""
// // 	if len(req.Roles) == 1 && req.Roles[0] == "all" {
// // 		roles = "<@all>"
// // 	} else {
// // 		for _, v := range req.Roles {
// // 			roles += fmt.Sprintf("<@&%s>", v)
// // 		}
// // 	}

// // 	card = strings.Replace(card, "{roles}", roles, -1)
// // 	card = strings.Replace(card, "{name}", config.Name, -1)
// // 	card = strings.Replace(card, "{activity}", config.ActivityID, -1)
// // 	card = strings.Replace(card, "{content}", req.Content, -1)
// // 	card = strings.Replace(card, "{color}", req.Color, -1)
// // 	err = json.Unmarshal([]byte(card), &message)
// // 	if err != nil {
// // 		return nil, err
// // 	}

// // 	msg, err := instance.SendChannelMessage(context.Background(), &model.SendChannelMessageReq{
// // 		ChannelId:   req.ChannelId,
// // 		MessageBody: &message,
// // 	})
// // 	if err != nil {
// // 		return nil, err
// // 	}

// // 	models.GetDB().Create(&models.PushInfo{
// // 		// ServerId:     req.ServerId,
// // 		// ServerName:   req.ServerName,
// // 		// ContractID:   config.ContractID,
// // 		// ActivityId:   req.ActivityId,
// // 		// ActivityName: config.Name,
// // 		// StartedTime:  config.StartedTime,
// // 		// EndedTime:    config.EndedTime,
// // 		// Contract:     config.ContractAddress,
// // 		// AccountLimit: req.AccountLimit,
// // 		ChannelId: req.ChannelId,
// // 		// Bot:           utils.DoDo,
// // 		// RainbowUserId: req.RainbowUserId,
// // 	})

// // 	return msg, nil
// // }

// func processErrorMessage(instance *client.Client, data *websocket.ChannelMessageEventBody, message string) {
// 	_, _ = (*instance).SendChannelMessage(context.Background(), &model.SendChannelMessageReq{
// 		ChannelId:   data.ChannelId,
// 		MessageBody: &model.TextMessage{Content: fmt.Sprintf("<@!%s> %s", data.DodoSourceId, message)},
// 	})
// }

// // func checkSocialLimit(serverId, userId, activity string, socialType int) error {
// // 	push, err := models.FindServerByChannel(serverId, activity)
// // 	if err != nil {
// // 		return err
// // 	}
// // 	if push.AccountLimit == -1 {
// // 		return nil
// // 	}

// // 	count, err := models.CountPOAPResultBySocial(userId, activity, uint(socialType))
// // 	if err != nil {
// // 		return err
// // 	}
// // 	if int(count) >= push.AccountLimit {
// // 		return fmt.Errorf("The userId has exceeded the account limit")
// // 	}
// // 	return nil
// // }

// func CheckIslandIsActive(instance *client.Client, islandId string) bool {
// 	_, err := (*instance).GetChannelList(context.Background(), &model.GetChannelListReq{
// 		IslandSourceId: islandId,
// 	})
// 	if err != nil {
// 		return false
// 	}

// 	return true
// }

// // func GetDoDoChannels(instance *client.Client, islandId string) ([]*model.ChannelElement, error) {
// // 	channels, err := (*instance).GetChannelList(context.Background(), &model.GetChannelListReq{
// // 		IslandSourceId: islandId,
// // 	})
// // 	if err != nil {
// // 		return nil, err
// // 	}
// // 	return channels, nil
// // }

// // func GetDoDoRoles(instance *client.Client, islandId string) ([]*model.RoleElement, error) {
// // 	channels, err := (*instance).GetRoleList(context.Background(), &model.GetRoleListReq{
// // 		IslandSourceId: islandId,
// // 	})
// // 	if err != nil {
// // 		return nil, err
// // 	}
// // 	return channels, nil
// // }

// func checkDoDoChannels(instance *client.Client, islandId string) bool {
// 	channels, err := (*instance).GetChannelList(context.Background(), &model.GetChannelListReq{
// 		IslandSourceId: islandId,
// 	})
// 	if err != nil {
// 		return false
// 	}

// 	for _, v := range channels {
// 		if v.ChannelName == channelName {
// 			return true
// 		}
// 	}
// 	return false
// }

// func GetInstance() *dodoClient.Client {
// 	return &instance
// }
