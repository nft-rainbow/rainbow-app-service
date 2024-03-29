package services

import (
	"fmt"
	"log"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/models/enums"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type PushReq struct {
	ServerId      string   `json:"server_id" binding:"required"`
	ServerName    string   `json:"server_name"  binding:"required"`
	ChannelId     string   `json:"channel_id"  binding:"required"`
	Roles         []string `json:"roles" binding:"required"`
	AccountLimit  int      `json:"account_limit"`
	Color         string   `json:"color"`
	Content       string   `json:"content"`
	Bot           uint     `json:"bot" binding:"required"`
	UserId        uint     `json:"user_id" binding:"required"`
	ActivityId    string   `json:"activity_id" binding:"required"`
	RainbowUserId int32    `gorm:"type:integer" json:"rainbow_user_id"`
	AppId         int32    `gorm:"index" json:"app_id"`
}

var s *discordgo.Session

var (
	guide       = "教程还没出"
	anywebH5    = "https://open.imdodo.com/dev/api/channel-text.html#%E5%8F%91%E9%80%81%E6%B6%88%E6%81%AF"
	channelName = "nft-rainbow-ai"
)

var (
	Commands = []*discordgo.ApplicationCommand{
		{
			Name:        "mint",
			Description: "Command for minting NFTs",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "activity",
					Description: "The id of the activity",
					Type:        discordgo.ApplicationCommandOptionString,
					Required:    true,
				},
				{
					Name:        "command",
					Description: "The mint command of the activity",
					Type:        discordgo.ApplicationCommandOptionString,
					Required:    false,
				},
			},
		},
		{
			Name:        "bind",
			Description: "Command for binding addresses",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "cfx",
					Description: "Bind the discord account with the conflux account",
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionString,
							Name:        "conflux_address",
							Description: "User's conflux address",
							Required:    true,
						},
					},
					Type: discordgo.ApplicationCommandOptionSubCommand,
				},
			},
		},
		{
			Name:        "address",
			Description: "Command for binding addresses",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "conflux",
					Description: "Get the binding conflux account",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
				},
			},
		},
		{
			Name:        "command",
			Description: "query the command for the specific activity",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "activity_id",
					Description: "Get the binding conflux account",
					Type:        discordgo.ApplicationCommandOptionString,
					Required:    true,
				},
			},
		},
		{
			Name:        "guide",
			Description: "query the guide",
		},
		{
			Name:        "create",
			Description: "create the cfx address",
		},
	}

	CommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"mint": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			if !checkDiscordChannel(i.ChannelID, i.GuildID) {
				return
			}
			options := i.ApplicationCommandData().Options
			activityId := options[0].Value
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: fmt.Sprintf("%v Start to mint using custom-mint model. Please wait patiently.", i.Interaction.Member.User.Mention()),
				},
			})
			bind, err := models.FindSocialUserConfig(i.Interaction.Member.User.ID, enums.SOCIAL_TOOL_DISCORD)
			if err != nil {
				s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
					Embeds: failMessageEmbed(err.Error()),
				})
				return
			}

			config, err := models.FindActivityByCode(activityId.(string))
			if err != nil {
				s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
					Embeds: failMessageEmbed(err.Error()),
				})
				return
			}
			// if err = config.CheckActivityValid(); err != nil {
			// 	s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
			// 		Embeds: failMessageEmbed(err.Error()),
			// 	})
			// 	return
			// }

			// err = checkSocialLimit(i.Interaction.GuildID, i.Interaction.Member.User.ID, config.ActivityID, utils.DoDo)
			// if err != nil {
			// 	s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
			// 		Embeds: failMessageEmbed(err.Error()),
			// 	})
			// 	return
			// }
			var command string
			if len(options) > 1 {
				command = options[1].Value.(string)
			}
			res, err := GetActivityService().H5Mint(&MintReq{
				ActivityID:  config.ActivityCode,
				UserAddress: bind.CFXAddress,
				Command:     command,
			})
			if err != nil {
				s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
					Embeds: failMessageEmbed(err.Error()),
				})
				return
			}

			for {
				resp, _ := models.FindPOAPResultById(config.ActivityCode, int(res.ID))
				if resp.Hash == "" {
					continue
				}
				s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
					Content: fmt.Sprintf("%v Mint NFT successfully. The correspding transaction hash is %s", i.Interaction.Member.User.Mention(), resp.Hash),
				})
				break
			}

		},
		"bind": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			if !checkDiscordChannel(i.ChannelID, i.GuildID) {
				return
			}
			options := i.ApplicationCommandData().Options
			userAddress := options[0].Options[0].Value.(string)
			startFlag := ""
			var err error
			switch options[0].Name {
			case "cfx":
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: startFlag,
					},
				})
				_, _, err = BindCfxAddress(i.Interaction.Member.User.ID, userAddress, enums.SOCIAL_TOOL_DISCORD)
			}
			if err != nil {
				s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
					Embeds: failMessageEmbed(err.Error()),
				})
				return
			}

			s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
				Content: fmt.Sprintf("%v success", i.Interaction.Member.User.Mention()),
			})
		},
		"address": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			if !checkDiscordChannel(i.ChannelID, i.GuildID) {
				return
			}
			options := i.ApplicationCommandData().Options
			switch options[0].Name {
			case "conflux":
				resp, err := models.FindSocialUserConfig(i.Interaction.Member.User.ID, enums.SOCIAL_TOOL_DISCORD)
				if err != nil {
					s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
						Embeds: failMessageEmbed(err.Error()),
					})
					return
				}
				s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
					Content: resp.CFXAddress,
				})
			}
		},
		"command": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			if !checkDiscordChannel(i.ChannelID, i.GuildID) {
				return
			}
			options := i.ApplicationCommandData().Options
			activity := options[0].Value
			config, err := models.FindActivityByCode(activity.(string))
			if err != nil {
				s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
					Embeds: failMessageEmbed(err.Error()),
				})
				return
			}
			channel, err := s.UserChannelCreate(i.Interaction.Member.User.ID)
			if err != nil {
				s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
					Embeds: failMessageEmbed(err.Error()),
				})
				return
			}
			if config.Command == "" {
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: fmt.Sprintf("%v The command is not needed in this activity", i.Interaction.Member.User.Mention()),
					},
				})
				return
			}
			s.ChannelMessageSend(channel.ID, config.Command)
		},
		"guide": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			if !checkDiscordChannel(i.ChannelID, i.GuildID) {
				return
			}
			s.ChannelMessageSend(i.ChannelID, fmt.Sprintf(i.Interaction.Member.User.Mention()+guide))
		},
		"create": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			if !checkDiscordChannel(i.ChannelID, i.GuildID) {
				return
			}
			s.ChannelMessageSend(i.ChannelID, fmt.Sprintf(i.Interaction.Member.User.Mention()+anywebH5))
		},
	}
)

func successfulMessageEmbed(resp *models.POAPResult, contract string) []*discordgo.MessageEmbed {
	embeds := []*discordgo.MessageEmbed{
		&discordgo.MessageEmbed{
			Type:        discordgo.EmbedTypeRich,
			Title:       ":rainbow: Mint NFT successfully  :rainbow:",
			Description: "Congratulate on minting NFT successfully! The NFT information is showed in the following.",
			Image: &discordgo.MessageEmbedImage{
				URL: "https://img0.baidu.com/it/u=2475308105,1312864556&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=889",
			},
			Provider: &discordgo.MessageEmbedProvider{
				Name: "come",
				URL:  "https://img0.baidu.com/it/u=2475308105,1312864556&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=889",
			},
			Fields: []*discordgo.MessageEmbedField{
				//&discordgo.MessageEmbedField{
				//	Name:   "Mints Time",
				//	Value:  resp.CreatedAt.String(),
				//	Inline: true,
				//},
				&discordgo.MessageEmbedField{
					Name:   "Hash",
					Value:  resp.Hash,
					Inline: true,
				},
				&discordgo.MessageEmbedField{
					Name:   "Contract",
					Value:  contract,
					Inline: true,
				},
				&discordgo.MessageEmbedField{
					Name:   "Token ID",
					Value:  resp.TokenID,
					Inline: true,
				},
				&discordgo.MessageEmbedField{
					Name:   "NFT URL",
					Value:  fmt.Sprintf("[VIEW IN CONFLUX SCAN](%s)", viper.GetString("customMint.mintRespPrefix")+contract+"/"+resp.TokenID),
					Inline: false,
				},
				&discordgo.MessageEmbedField{
					Name:   "Advertise",
					Value:  viper.GetString("advertise"),
					Inline: false,
				},
			},
			//Author: &discordgo.MessageEmbedAuthor{
			//	Name:    "NFTRainbow",
			//	URL:     "https://docs.nftrainbow.xyz/",
			//	IconURL: "https://img0.baidu.com/it/u=2475308105,1312864556&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=889",
			//},
		},
	}
	return embeds
}

func failMessageEmbed(message string) []*discordgo.MessageEmbed {
	embeds := []*discordgo.MessageEmbed{
		&discordgo.MessageEmbed{
			Type:        discordgo.EmbedTypeRich,
			Title:       ":scream: Failed to Mint NFT  :scream:",
			Description: "There is problem during minting NFT. ",
			Image: &discordgo.MessageEmbedImage{
				URL: "https://gimg2.baidu.com/image_search/src=http%3A%2F%2Ftva1.sinaimg.cn%2Fbmiddle%2F006APoFYly1g55m70z1uvj30hs0hidhd.jpg&refer=http%3A%2F%2Ftva1.sinaimg.cn&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=auto?sec=1664935347&t=223d106a8cbc9c825b5a34ff36b3678c",
			},
			Fields: []*discordgo.MessageEmbedField{
				&discordgo.MessageEmbedField{
					Name:   "Error message",
					Value:  message,
					Inline: false,
				},
				&discordgo.MessageEmbedField{
					Name:   "Advertise",
					Value:  viper.GetString("advertise"),
					Inline: false,
				},
			},
			Author: &discordgo.MessageEmbedAuthor{
				Name:    "NFTRainbow",
				URL:     "https://docs.nftrainbow.xyz/",
				IconURL: "https://img0.baidu.com/it/u=2475308105,1312864556&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=889",
			},
		},
	}

	return embeds
}

func DiscordPushActivity(req *PushReq) (*discordgo.Message, error) {
	config, err := models.FindActivityByCode(req.ActivityId)
	if err != nil {
		return nil, err
	}

	color, err := strconv.ParseInt(req.Color, 0, 64)
	if err != nil {
		return nil, err
	}
	roles := ""
	for _, v := range req.Roles {
		roles += fmt.Sprintf("<@&%s>", v)
	}
	msg, err := s.ChannelMessageSendEmbeds(req.ChannelId, createPushEmbed(config, roles, req.Content, int(color)))
	if err != nil {
		return nil, err
	}

	models.GetDB().Create(&models.PushInfo{
		// ServerId:      req.ServerId,
		// ServerName:    req.ServerName,
		// ContractID:    config.ContractID,
		// ActivityId:    req.ActivityId,
		// ActivityName:  config.Name,
		// StartedTime:   config.StartedTime,
		// EndedTime:     config.EndedTime,
		// AccountLimit:  req.AccountLimit,
		// Contract:      config.ContractAddress,
		// ChannelId:     req.ChannelId,
		// Bot:           utils.Discord,
		// RainbowUserId: req.RainbowUserId,
	})

	return msg, nil
}

func createPushEmbed(config *models.Activity, roles, content string, color int) []*discordgo.MessageEmbed {
	embeds := []*discordgo.MessageEmbed{
		&discordgo.MessageEmbed{
			Type:        discordgo.EmbedTypeRich,
			Title:       "新活动发布",
			Description: fmt.Sprintf("<@&%v> %v #%v 来了！\n在频道中调用【\\guide】，机器人将私信你领取教程", roles, config.Name, config.ActivityCode),
			Color:       color,
			Fields: []*discordgo.MessageEmbedField{
				&discordgo.MessageEmbedField{
					Name:   "内容",
					Value:  content,
					Inline: false,
				},
			},
			Author: &discordgo.MessageEmbedAuthor{
				Name:    "NFTRainbow",
				URL:     "https://docs.nftrainbow.xyz/",
				IconURL: "https://img0.baidu.com/it/u=2475308105,1312864556&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=889",
			},
			Image: &discordgo.MessageEmbedImage{
				URL: "https://gimg2.baidu.com/image_search/src=http%3A%2F%2Ftva1.sinaimg.cn%2Fbmiddle%2F006APoFYly1g55m70z1uvj30hs0hidhd.jpg&refer=http%3A%2F%2Ftva1.sinaimg.cn&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=auto?sec=1664935347&t=223d106a8cbc9c825b5a34ff36b3678c",
			},
		},
	}
	return embeds
}

func checkDiscordChannel(channelId, guildId string) bool {
	botServer, err := models.FindBotServerByChannel(channelId)
	if err != nil {
		logrus.Errorf("Failed to find channel: %v", err.Error())
		return false
	}
	if botServer.RawServerId != guildId {
		return false
	}
	//channels, err := s.GuildChannels(guildId)
	//if err != nil {
	//	logrus.Errorf("Failed to find channel: %v", err.Error())
	//	return false
	//}
	return true
}

func guildCreate(s *discordgo.Session, event *discordgo.GuildCreate) {
	channels, err := s.GuildChannels(event.Guild.ID)
	if err != nil {
		logrus.Errorf("Failed to find channel: %v", err.Error())
		return
	}
	for _, v := range channels {
		if v.Name == channelName {
			return
		}
	}
	_, err = s.GuildChannelCreateComplex(event.Guild.ID, discordgo.GuildChannelCreateData{
		Name:                 channelName,
		Type:                 discordgo.ChannelTypeGuildText,
		PermissionOverwrites: []*discordgo.PermissionOverwrite{},
		Topic:                "",
		NSFW:                 false,
	})
	if err != nil {
		logrus.Errorf("Failed to create channel: %v", err.Error())
		return
	}
}

func GetDiscordChannels(s *discordgo.Session, guild string) ([]*discordgo.Channel, error) {
	channels, err := s.GuildChannels(guild)
	if err != nil {
		return nil, err
	}

	return channels, nil
}

func GetDiscordRoles(s *discordgo.Session, guild string) ([]*discordgo.Role, error) {
	roles, err := s.GuildRoles(guild)
	if err != nil {
		return nil, err
	}

	return roles, nil
}

func CheckGuildIsActive(s *discordgo.Session, guild string) bool {
	_, err := s.Guild(guild)
	if err != nil {
		return false
	}
	return true
}

func InitSession() *discordgo.Session {
	var err error

	s, err = discordgo.New("Bot " + viper.GetString("discordBotToken"))
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}

	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := CommandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})

	s.AddHandler(guildCreate)
	return s
}

func GetSession() *discordgo.Session {
	return s
}
