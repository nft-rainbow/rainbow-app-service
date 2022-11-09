package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/nft-rainbow/rainbow-app-service/middlewares"
	"github.com/nft-rainbow/rainbow-app-service/models"
	openapiclient "github.com/nft-rainbow/rainbow-sdk-go"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

func DiscordActivityConfig(config *models.DiscordActivityConfig, token string) error {
	info, err := GetContractInfo(config.ContractID, token)
	if err != nil {
		return err
	}
	config.ContractType = *info.Type
	config.Chain = *info.ChainType
	config.AppId = *info.AppId
	config.ContractAddress = *info.Address

	res := models.GetDB().Create(&config)
	if res.Error != nil {
		return  res.Error
	}

	return nil
}

func DoDoActivityConfig(config *models.DoDoActivityConfig, token string) error {
	info, err := GetContractInfo(config.ContractID, token)
	if err != nil {
		return err
	}
	config.ContractType = *info.Type
	config.Chain = *info.ChainType
	config.AppId = *info.AppId
	config.ContractAddress = *info.Address

	res := models.GetDB().Create(&config)
	if res.Error != nil {
		return  res.Error
	}

	return nil
}

func HandleCustomMint(userId, channelId, platform string) (*openapiclient.ModelsMintTask, string, int32, error){
	req := models.MintReq{
		UserID: userId,
		ChannelID: channelId,
	}
	if platform == "dodo" {
		resp, token, contractId, err := dodoCustomMint(&req)
		return resp, token, contractId, err
	}else if platform == "discord" {
		resp, token, contractId, err := discordCustomMint(&req)
		return resp, token, contractId, err
	}else {
		return nil, "", 0, nil
	}
}

func dodoCustomMint(req *models.MintReq) (*openapiclient.ModelsMintTask, string, int32, error){
	config, err := models.FindBindingDoDoActivityConfigByChannelId(req.ChannelID)
	if err != nil {
		return nil, "", 0, err
	}

	ok, err := models.CheckDoDoCustomCount(req.UserID, req.ChannelID, config.MaxMintCount)
	if err != nil {
		return nil, "", 0, err
	}
	if !ok {
		return nil, "", 0, errors.New("This number of the NFTs the account minted has reached the maximum")
	}

	cfxAddress, err := GetDoDoBindCFXAddress(req.UserID)
	if err != nil {
		return nil, "", 0, err
	}

	token, _ := middlewares.GenerateDoDoOpenJWT(req.ChannelID)
	var contractType string
	if config.ContractType == 1 {
		contractType = "erc721"
	}else {
		contractType = "erc1155"
	}
	var chainType string
	if config.Chain == 1 {
		chainType = "conflux_test"
	}else {
		chainType = "conflux"
	}
	resp , err := sendCustomMintRequest("Bearer " + token, openapiclient.ServicesCustomMintDto{
		Chain: chainType,
		ContractType: contractType,
		ContractAddress: config.ContractAddress,
		MintToAddress: cfxAddress,
		MetadataUri: &config.MetadataURI,
	})
	if err != nil {
		return nil, "", 0, err
	}

	return resp, token, config.ContractID, err
}


func discordCustomMint(req *models.MintReq) (*openapiclient.ModelsMintTask, string, int32, error){
	config, err := models.FindBindingDiscordActivityConfigByChannelId(req.ChannelID)
	if err != nil {
		return nil, "", 0, err
	}

	ok, err := models.CheckDiscordCustomCount(req.UserID, req.ChannelID, config.MaxMintCount)
	if err != nil {
		return nil, "", 0, err
	}
	if !ok {
		return nil, "", 0, errors.New("This number of the NFTs the account minted has reached the maximum")
	}

	cfxAddress, err := GetDiscordBindCFXAddress(req.UserID)
	if err != nil {
		return nil, "", 0, err
	}

	token, _ := middlewares.GenerateDiscordOpenJWT(req.ChannelID)

	var contractType string
	if config.ContractType == 1 {
		contractType = "erc721"
	}else {
		contractType = "erc1155"
	}
	var chainType string
	if config.Chain == 1 {
		chainType = "conflux_test"
	}else {
		chainType = "conflux"
	}
	resp , err := sendCustomMintRequest("Bearer " + token, openapiclient.ServicesCustomMintDto{
		Chain: chainType,
		ContractType: contractType,
		ContractAddress: config.ContractAddress,
		MintToAddress: cfxAddress,
		MetadataUri: &config.MetadataURI,
	})
	if err != nil {
		return nil, "", 0, err
	}

	return resp, token, config.ContractID, err
}

func GenDiscordMintRes(token, createTime, contractAddress, userAddress, userID, channelID string, id, contractId int32)(*models.MintResp, error){
	tokenId, err := getTokenId(id, "Bearer " + token)
	if err != nil {
		return nil, err
	}

	res := &models.MintResp{
		UserAddress: userAddress,
		NFTAddress: viper.GetString("customMint.mintRespPrefix") +  contractAddress + "/" + tokenId,
		Contract: contractAddress,
		TokenID: tokenId,
		Time: createTime,
	}
	_, err = models.UpdateDiscordCustomCount(userID, channelID)
	if err != nil {
		return nil, err
	}

	err = models.StoreMintResult(models.MintResult{
		UserID: userID,
		ContractID: contractId,
		TokenID: tokenId,
	})
	return res, nil
}

func GenDoDoMintRes(token, createTime, contractAddress, userAddress, userID, channelID string, id, contractId int32)(*models.MintResp, error){
	tokenId, err := getTokenId(id, "Bearer " + token)
	if err != nil {
		return nil, err
	}

	res := &models.MintResp{
		UserAddress: userAddress,
		NFTAddress: viper.GetString("customMint.mintRespPrefix") +  contractAddress + "/" + tokenId,
		Contract: contractAddress,
		TokenID: tokenId,
		Time: createTime,
	}
	_, err = models.UpdateDoDoCustomCount(userID, channelID)
	if err != nil {
		return nil, err
	}

	err = models.StoreMintResult(models.MintResult{
		UserID: userID,
		ContractID: contractId,
		TokenID: tokenId,
	})
	return res, nil
}

func sendCustomMintRequest(token string, dto openapiclient.ServicesCustomMintDto) (*openapiclient.ModelsMintTask, error){
	//configuration := openapiclient.NewConfiguration()
	//apiClient := openapiclient.NewAPIClient(configuration)
	fmt.Println("Start to mint")
	resp, _, err := newClient().MintsApi.CustomMint(context.Background()).Authorization(token).CustomMintDto(dto).Execute()
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func getTokenId(id int32, token string) (string, error) {
	//configuration := openapiclient.NewConfiguration()
	//apiClient := openapiclient.NewAPIClient(configuration)
	fmt.Println("Start to get token Id")
	resp, _, err := newClient().MintsApi.GetMintDetail(context.Background(), id).Authorization(token).Execute()
	if err != nil {
		return "", err
	}
	for resp.TokenId == nil && *resp.Status != 1 {
		resp, _, err = newClient().MintsApi.GetMintDetail(context.Background(), id).Authorization(token).Execute()
		if err != nil {
			return "", err
		}
		time.Sleep(10 * time.Second)
	}
	return *resp.TokenId, nil
}

func GetContractInfo(id int32, token string) (*openapiclient.ModelsContract, error){
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
			URL: "https://dev.nftrainbow.xyz/v1",
		},
	}
	apiClient := openapiclient.NewAPIClient(configuration)
	return apiClient
}