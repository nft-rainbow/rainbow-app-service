package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/nft-rainbow/rainbow-app-service/models"
	openapiclient "github.com/nft-rainbow/rainbow-sdk-go"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

func ActivityConfig(config *models.ActivityConfig) error {
	token, err := Login()
	if err != nil {
		return err
	}
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

	_, err = addContractAdmin(*info.Address, "Bearer " + token)
	if err != nil {
		return err
	}
	return nil
}

func HandleCustomMint(userId, channelId string) (*openapiclient.ModelsMintTask, string, int32, error){
	req := models.MintReq{
		UserID: userId,
		ChannelID: channelId,
	}

	resp, token, contractId, err := customMint(&req)

	return resp, token, contractId, err
}

func customMint(req *models.MintReq) (*openapiclient.ModelsMintTask, string, int32, error){
	config, err := models.FindBindingActivityConfigByChannelId(req.ChannelID)
	if err != nil {
		return nil, "", 0, err
	}

	ok, err := models.CheckCustomCount(req.UserID, req.ChannelID, config.MaxMintCount)
	if err != nil {
		return nil, "", 0, err
	}
	if !ok {
		return nil, "", 0, errors.New("This number of the NFTs the account minted has reached the maximum")
	}

	tmp, err := GetBindCFXAddress(req.UserID)
	if err != nil {
		return nil, "", 0, err
	}

	token, err := Login()
	if err != nil {
		return nil, "", 0, err
	}

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
		MintToAddress: tmp.CFXAddress,
		MetadataUri: &config.MetadataURI,
	})
	if err != nil {
		return nil, "", 0, err
	}

	return resp, token, config.ContractID, err
}

func GenMintRes(token, createTime, contractAddress, userAddress, userID, channelID string, id, contractId int32)(*models.MintResp, error){
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
	_, err = models.UpdateCustomCount(userID, channelID)
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

func Login() (string, error) {
	fmt.Println("Start to login")
	resp, _, err := newClient().LoginApi.LoginApp(context.Background()).AppLoginInfo(openapiclient.MiddlewaresAppLogin{
		AppId: viper.GetString("app.appId"),
		AppSecret: viper.GetString("app.appSecret"),
	}).Execute()
	if err != nil {
		return "", err
	}
	t := make(map[string]interface{})
	err = json.Unmarshal([]byte(resp), &t)
	if err != nil {
		return "", err
	}
	if t["code"] != nil {
		return "", errors.New(t["message"].(string))
	}

	return t["token"].(string), nil
}

func GetContractInfo(id int32, token string) (*openapiclient.ModelsContract, error){
	//configuration := openapiclient.NewConfiguration()
	//apiClient := openapiclient.NewAPIClient(configuration)
	fmt.Println("Start to get contract information")
	resp, _, err := newClient().ContractApi.GetContractInfo(context.Background(), id).Authorization("Bearer " + token).Execute()
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