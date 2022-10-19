package services

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/nft-rainbow/discordbot-service/models"
	openapiclient "github.com/nft-rainbow/rainbow-sdk-go"
	"github.com/spf13/viper"
	"net/http"
	"net/url"
)

func BindAdminConfig(config *models.AdminConfig, id uint) error{
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
	token := viper.GetString("botToken")
	session, err := discordgo.New("Bot " + token)

	if err != nil {
		return nil, err
	}
	proxy, _ := url.Parse(viper.GetString("proxy"))

	tr := &http.Transport{
		Proxy:           http.ProxyURL(proxy),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	session.Client.Transport = tr

	st, err := session.GuildChannels(guildId)

	if err != nil {
		return nil, err
	}
	return st, err
}

func GetGuildInfo(guildId string) (st *discordgo.Guild, err error){
	token := viper.GetString("botToken")
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}
	proxy, _ := url.Parse(viper.GetString("proxy"))
	tr := &http.Transport{
		Proxy:           http.ProxyURL(proxy),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	session.Client.Transport = tr
	st, err = session.Guild(guildId)
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