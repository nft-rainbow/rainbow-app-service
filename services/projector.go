package services

import (
	"context"
	"github.com/bwmarrin/discordgo"
	dodoModel "github.com/dodo-open/dodo-open-go/model"
	"github.com/nft-rainbow/rainbow-app-service/models"
)

func BindDiscordProjectorConfig(config *models.DiscordAdminConfig, id uint) error{
	info, err := GetDiscordGuildInfo(config.GuildId)
	if err != nil {
		return err
	}

	config.GuildName = info.Name
	config.RainbowUserId = int32(id)

	res := models.GetDB().Create(&config)
	if res.Error != nil {
		return  res.Error
	}
	return nil
}

func BindDoDoProjectorConfig(config *models.DoDoAdminConfig, id uint) error{
	info, err := GetDoDoIslandInfo(config.IslandId)
	if err != nil {
		return err
	}

	config.IslandName = info.IslandName
	config.RainbowUserId = int32(id)

	res := models.GetDB().Create(&config)
	if res.Error != nil {
		return  res.Error
	}
	return nil
}

func GetDiscordChannelInfo(guildId string) ([]*discordgo.Channel, error){
	st, err := GetSession().GuildChannels(guildId)

	if err != nil {
		return nil, err
	}
	return st, err
}

func GetDoDoChannelInfo(islandId string) ([]*dodoModel.ChannelElement, error){
	st, err := (*GetInstance()).GetChannelList(context.Background(), &dodoModel.GetChannelListReq{
		IslandId: islandId,
	})

	if err != nil {
		return nil, err
	}
	return st, err
}

func GetDiscordGuildInfo(guildId string) (st *discordgo.Guild, err error){
	st, err = GetSession().Guild(guildId)
	if err != nil {
		return nil, err
	}
	return st, err
}

func GetDoDoIslandInfo(islandId string) (st *dodoModel.GetIslandInfoRsp, err error){
	info, err := (*GetInstance()).GetIslandInfo(context.Background(), &dodoModel.GetIslandInfoReq{
		IslandId: islandId,
	})
	if err != nil {
		return nil, err
	}
	return info, err
}