package services

import (
	"context"
	"fmt"
	"github.com/bwmarrin/discordgo"
	dodoModel "github.com/dodo-open/dodo-open-go/model"
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/utils"
	openapiclient "github.com/nft-rainbow/rainbow-sdk-go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

func bindCFXAddressWithDiscord(req *models.BindCFXWithDiscord) error {
	res := models.GetDB().Create(&req)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func bindCFXAddressWithDoDo(req *models.BindCFXWithDoDo) error {
	res := models.GetDB().Create(&req)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func GetDoDoBindCFXAddress(userID string) (string, error) {
	resp, err := models.FindDoDoBindingCFXAddressById(userID)
	if err != nil {
		return "", err
	}
	return resp.CFXAddress, nil
}

func GetDiscordBindCFXAddress(userID string) (string, error) {
	resp, err := models.FindDiscordBindingCFXAddressById(userID)
	if err != nil {
		return "", err
	}
	return resp.CFXAddress, nil
}

func HandleBindCfxAddress(userId, userAddress, platform string) error {
	var err error
	_, err = utils.CheckCfxAddress(utils.CONFLUX_TEST, userAddress)
	if err != nil {
		return err
	}

	if platform == "discord" {
		err = bindCFXAddressWithDiscord(&models.BindCFXWithDiscord{
			DiscordId:  userId,
			CFXAddress: userAddress,
		})
	} else if platform == "dodo" {
		err = bindCFXAddressWithDoDo(&models.BindCFXWithDoDo{
			DoDoId:     userId,
			CFXAddress: userAddress,
		})
	}
	if err != nil {
		return err
	}

	return nil
}

func GetDiscordChannelInfo(guildId string) ([]*discordgo.Channel, error) {
	st, err := GetSession().GuildChannels(guildId)

	if err != nil {
		return nil, err
	}
	return st, err
}

func GetDoDoChannelInfo(islandId string) ([]*dodoModel.ChannelElement, error) {
	st, err := (*GetInstance()).GetChannelList(context.Background(), &dodoModel.GetChannelListReq{
		IslandId: islandId,
	})

	if err != nil {
		return nil, err
	}
	return st, err
}

func GetDiscordGuildInfo(guildId string) (st *discordgo.Guild, err error) {
	st, err = GetSession().Guild(guildId)
	if err != nil {
		return nil, err
	}
	return st, err
}

func GetDoDoIslandInfo(islandId string) (st *dodoModel.GetIslandInfoRsp, err error) {
	info, err := (*GetInstance()).GetIslandInfo(context.Background(), &dodoModel.GetIslandInfoReq{
		IslandId: islandId,
	})
	if err != nil {
		return nil, err
	}
	return info, err
}

func GenDiscordMintRes(token, createTime, contractAddress, userAddress, userID, channelID string, id, contractId int32) (*models.CustomMintResp, error) {
	tokenId, hash, err := getTokenInfo(id, "Bearer "+token)
	if err != nil {
		return nil, err
	}

	res := &models.CustomMintResp{
		UserAddress: userAddress,
		NFTAddress:  viper.GetString("customMint.mintRespPrefix") + contractAddress + "/" + tokenId,
		Contract:    contractAddress,
		TokenID:     tokenId,
		Time:        createTime,
	}
	_, err = models.UpdateDiscordCustomCount(userID, channelID)
	if err != nil {
		return nil, err
	}

	err = models.StoreCustomMintResult(models.CustomMintResult{
		UserID:     userID,
		ContractID: contractId,
		TokenID:    tokenId,
		Hash: hash,
	})
	return res, nil
}

func GenDoDoMintRes(token, createTime, contractAddress, userAddress, userID, channelID string, id, contractId int32) (*models.CustomMintResp, error) {
	tokenId, hash, err := getTokenInfo(id, "Bearer "+token)
	if err != nil {
		return nil, err
	}

	res := &models.CustomMintResp{
		UserAddress: userAddress,
		NFTAddress:  viper.GetString("customMint.mintRespPrefix") + contractAddress + "/" + tokenId,
		Contract:    contractAddress,
		TokenID:     tokenId,
		Time:        createTime,
	}
	_, err = models.UpdateDoDoCustomCount(userID, channelID)
	if err != nil {
		return nil, err
	}

	err = models.StoreCustomMintResult(models.CustomMintResult{
		UserID:     userID,
		ContractID: contractId,
		TokenID:    tokenId,
		Hash: hash,
	})
	return res, nil
}

func sendBurnNFTRequest(token string, dto openapiclient.ServicesBurnDto) (*openapiclient.ModelsBurnTask, error) {
	fmt.Println("Start to burn")
	resp, _, err := newClient().BurnsApi.BurnNft(context.Background()).Authorization(token).BurnDto(dto).Execute()
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func sendCustomMintRequest(token string, dto openapiclient.ServicesCustomMintDto) (*openapiclient.ModelsMintTask, error) {
	//configuration := openapiclient.NewConfiguration()
	//apiClient := openapiclient.NewAPIClient(configuration)
	fmt.Println("Start to mint")
	resp, _, err := newClient().MintsApi.CustomMint(context.Background()).Authorization(token).CustomMintDto(dto).Execute()
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func getTokenInfo(id int32, token string) (string, string, error) {
	//configuration := openapiclient.NewConfiguration()
	//apiClient := openapiclient.NewAPIClient(configuration)
	//fmt.Println("Start to get token Id")
	resp, _, err := newClient().MintsApi.GetMintDetail(context.Background(), id).Authorization(token).Execute()
	if err != nil {
		return "", "", err
	}

	for *resp.Status != 1 && *resp.Hash == ""{
		resp, _, err = newClient().MintsApi.GetMintDetail(context.Background(), id).Authorization(token).Execute()
		if err != nil {
			return "", "", err
		}
		time.Sleep(3 * time.Second)
	}
	return *resp.TokenId, *resp.Hash, nil
}

func GetContractInfo(id int32, token string) (*openapiclient.ModelsContract, error) {
	//configuration := openapiclient.NewConfiguration()
	//apiClient := openapiclient.NewAPIClient(configuration)
	fmt.Println("Start to get contract information")
	resp, _, err := newClient().ContractApi.GetContractInfo(context.Background(), id).Authorization(token).Execute()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func newClient() *openapiclient.APIClient {
	configuration := openapiclient.NewConfiguration()
	configuration.HTTPClient = http.DefaultClient
	configuration.Servers = openapiclient.ServerConfigurations{
		{
			URL: viper.GetString("rainbowOpenApi") + "/v1",
		},
	}
	apiClient := openapiclient.NewAPIClient(configuration)
	return apiClient
}

func SyncNFTMintTaskStatus(token string, res *models.POAPResult) {
	logrus.Info("start task for syncing nft mint status")
	tokenId, hash, _ := getTokenInfo(res.TxID, "Bearer "+token)

	res.TokenID = tokenId
	res.Hash = hash

	models.GetDB().Save(&res)
}
