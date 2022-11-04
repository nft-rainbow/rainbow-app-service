package services

import (
	"context"
	"fmt"
	"github.com/bwmarrin/discordgo"
	dodoModel "github.com/dodo-open/dodo-open-go/model"
	"github.com/nft-rainbow/rainbow-app-service/models"
	openapiclient "github.com/nft-rainbow/rainbow-sdk-go"
	"github.com/spf13/viper"
)

func BindDiscordProjectorConfig(config *models.DiscordAdminConfig, id uint) error{
	info, err := GetDiscordGuildInfo(config.GuildId)
	if err != nil {
		return err
	}

	config.GuildName = info.Name
	config.AppId = int32(id)
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
	config.AppId = int32(id)
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

func addContractAdmin(contract, token string) (*openapiclient.ServicesSendTxResp, error){
	fmt.Println("Start to add contract admin")
	//configuration := openapiclient.NewConfiguration()
	//apiClient := openapiclient.NewAPIClient(configuration)
	resp, _, err := newClient().ContractApi.UpdateContractAdmin(context.Background(), contract).Authorization(token).AdminInfo(*openapiclient.NewServicesContractAdminUpdateDto(viper.GetString("botAddress"))).Execute()
	if err != nil {
		return nil, err
	}
	return resp, err
}