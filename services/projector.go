package services

import (
	"context"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/nft-rainbow/rainbow-app-service/models"
	openapiclient "github.com/nft-rainbow/rainbow-sdk-go"
	"github.com/spf13/viper"
)

func BindProjectorConfig(config *models.AdminConfig, id uint) error{
	info, err := GetGuildInfo(config.GuildId)
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

func GetChannelInfo(guildId string) ([]*discordgo.Channel, error){
	st, err := GetSession().GuildChannels(guildId)

	if err != nil {
		return nil, err
	}
	return st, err
}

func GetGuildInfo(guildId string) (st *discordgo.Guild, err error){
	st, err = GetSession().Guild(guildId)
	if err != nil {
		return nil, err
	}
	return st, err
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