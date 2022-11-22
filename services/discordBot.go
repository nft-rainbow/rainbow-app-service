package services

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/spf13/viper"
	"log"
)

var s *discordgo.Session

var (
	Commands = []*discordgo.ApplicationCommand{
		{
			Name:        "claim",
			Description: "Command for claiming NFTs",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "custom-mint",
					Description: "Mint a nft through the contract deployed by the admin",
					Type: discordgo.ApplicationCommandOptionSubCommand,
				},
				{
					Name:        "easy-mint",
					Description: "Mint a nft through the NFTfactory contract owned by NFTRainbow",
					Type: discordgo.ApplicationCommandOptionSubCommand,
				},
			},
		},
		{
			Name:        "bind",
			Description: "Command for binding addresses",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "conflux",
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
					Type: discordgo.ApplicationCommandOptionSubCommand,
				},
			},
		},
	}

	CommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"claim": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			options := i.ApplicationCommandData().Options
			startFlag := ""
			switch options[0].Name {
			case "custom-mint":
				startFlag = "Start to mint using custom-mint model. Please wait patiently."
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: startFlag,
					},
				})

				resp, token, contactId, err := HandleCustomMint(i.Interaction.Member.User.ID, i.ChannelID, "discord")
				if err != nil {
					s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
						Embeds: failMessageEmbed(err.Error()),
					})
					return
				}
				s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
					Content: fmt.Sprintf("Create mint task successfully. The correspding transaction hash is %s", *resp.Hash),
				})

				res, err := GenDiscordMintRes(token, resp.GetCreatedAt(), resp.GetContract(), resp.GetMintTo(),  i.Interaction.Member.User.ID, i.ChannelID, resp.GetId(), contactId)
				if err != nil {
					s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
						Embeds: failMessageEmbed(err.Error()),
					})
					return
				}
				s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{

					Embeds: successfulMessageEmbed(res),
				})
			}
		},
		"bind": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			options := i.ApplicationCommandData().Options
			userAddress := options[0].Options[0].Value.(string)
			startFlag := ""
			var err error
			switch options[0].Name {
			case "conflux":
				startFlag = "Start to bind cfx address."
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: startFlag,
					},
				})
				err = HandleBindCfxAddress(i.Interaction.Member.User.ID, userAddress, "discord")
			}
			if err != nil {
				s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
					Embeds: failMessageEmbed(err.Error()),
				})
				return
			}

			s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
				Content: "success",
			})
		},
		"address": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			options := i.ApplicationCommandData().Options
			startFlag := ""
			switch options[0].Name {
			case "conflux":
				startFlag = "Start to get binding cfx address."
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: startFlag,
					},
				})
				resp, err := GetDiscordBindCFXAddress(i.Interaction.Member.User.ID)
				if err != nil {
					s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
						Embeds: failMessageEmbed(err.Error()),
					})
					return
				}
				s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
					Content: resp,
				})
			}
		},
	}
)

func successfulMessageEmbed(resp *models.CustomMintResp) []*discordgo.MessageEmbed{
	embeds := []*discordgo.MessageEmbed{
		&discordgo.MessageEmbed{
			Type: discordgo.EmbedTypeRich,
			Title: ":rainbow: Mint NFT successfully  :rainbow:",
			Description: "Congratulate on minting NFT successfully! The NFT information is showed in the following.",
			Image: &discordgo.MessageEmbedImage{
				URL: "https://img0.baidu.com/it/u=2475308105,1312864556&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=889",
			},
			Provider: &discordgo.MessageEmbedProvider{
				Name: "come",
				URL: "https://img0.baidu.com/it/u=2475308105,1312864556&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=889",
			},
			Fields: []*discordgo.MessageEmbedField{
				&discordgo.MessageEmbedField{
					Name: "Mints Time",
					Value: resp.Time,
					Inline: true,
				},
				&discordgo.MessageEmbedField{
					Name: "Contract",
					Value: resp.Contract,
					Inline: true,
				},
				&discordgo.MessageEmbedField{
					Name: "Token ID",
					Value: resp.TokenID,
					Inline: true,
				},
				&discordgo.MessageEmbedField{
					Name: "NFT URL",
					Value: fmt.Sprintf("[VIEW IN CONFLUX SCAN](%s)", resp.NFTAddress),
					Inline: false,
				},
				&discordgo.MessageEmbedField{
					Name: "Advertise",
					Value: viper.GetString("advertise"),
					Inline: false,
				},
			},
			Author: &discordgo.MessageEmbedAuthor{
				Name: "NFTRainbow",
				URL: "https://docs.nftrainbow.xyz/",
				IconURL: "https://img0.baidu.com/it/u=2475308105,1312864556&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=889",
			},
		},
	}
	return embeds
}

func failMessageEmbed(message string) []*discordgo.MessageEmbed{
	embeds := []*discordgo.MessageEmbed{
		&discordgo.MessageEmbed{
			Type: discordgo.EmbedTypeRich,
			Title: ":scream: Failed to Mint NFT  :scream:",
			Description: "There is problem during minting NFT. ",
			Image: &discordgo.MessageEmbedImage{
				URL: "https://gimg2.baidu.com/image_search/src=http%3A%2F%2Ftva1.sinaimg.cn%2Fbmiddle%2F006APoFYly1g55m70z1uvj30hs0hidhd.jpg&refer=http%3A%2F%2Ftva1.sinaimg.cn&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=auto?sec=1664935347&t=223d106a8cbc9c825b5a34ff36b3678c",
			},
			Fields: []*discordgo.MessageEmbedField{
				&discordgo.MessageEmbedField{
					Name: "Error message",
					Value: message,
					Inline: false,
				},
				&discordgo.MessageEmbedField{
					Name: "Advertise",
					Value: viper.GetString("advertise"),
					Inline: false,
				},
			},
			Author: &discordgo.MessageEmbedAuthor{
				Name: "NFTRainbow",
				URL: "https://docs.nftrainbow.xyz/",
				IconURL: "https://img0.baidu.com/it/u=2475308105,1312864556&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=889",
			},
		},
	}

	return embeds
}

func InitSession()*discordgo.Session {
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
	return s

}

func GetSession() *discordgo.Session{
	return s
}
