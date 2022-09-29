package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/nft-rainbow/discordbot-service/models"
	openapiclient "github.com/nft-rainbow/rainbow-sdk-go"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

func ActivityConfig(config *models.ActivityConfig) error {
	if config.MaxMintCount == 0 {
		config.MaxMintCount = 1
	}
	config.Event = "customMint"
	res := models.GetDB().Create(&config)
	if res.Error != nil {
		return  res.Error
	}
	token, err := Login()
	if err != nil {
		return err
	}
	info, err := GetContractInfo(config.ContractID, token)
	if err != nil {
		return err
	}
	_, err = addContractAdmin(*info.Address, "Bearer " + token)
	if err != nil {
		return err
	}
	return nil
}

func CustomMint(req *models.MintReq) (*models.MintResp, error){
	config, err := models.FindBindingActivityConfigById(req.ChannelID)
	if err != nil {
		return nil, err
	}

	ok, err := models.CheckCustomCount(req.UserID, req.ChannelID, config.MaxMintCount)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("This number of the NFTs the account minted has reached the maximum")
	}

	tmp, err := GetBindCFXAddress(req.UserID)
	if err != nil {
		return nil, err
	}

	token, err := Login()
	if err != nil {
		return nil, err
	}

	info, err := GetContractInfo(int32(config.ContractID), token)
	if err != nil {
		return nil, err
	}
	var contractType string
	if *info.Type == 1 {
		contractType = "erc721"
	}else {
		contractType = "erc1155"
	}
	var chainType string
	if *info.ChainType == 1 {
		chainType = "conflux_test"
	}else {
		chainType = "conflux_main"
	}
	uri := "https://dev.nftrainbow.xyz/assets/file/2/nft/86db42aac9db6dbead473d7d49e1eaa4d6e9fcb3be86684ee56c210bc284b551.png"
	resp , err := sendCustomMintRequest("Bearer " + token, openapiclient.ServicesCustomMintDto{
		Chain: chainType,
		ContractType: contractType,
		ContractAddress: *info.Address,
		MintToAddress: tmp.UserAddress,
		MetadataUri: &uri,
	})
	if err != nil {
		return nil, err
	}

	_, err = models.UpdateCustomCount(req.UserID, req.ChannelID)
	if err != nil {
		return nil, err
	}

	err = models.StoreMintResult(models.MintResult{
		UserID: req.UserID,
		ContractID: *info.Id,
		TokenID: resp.TokenID,
	})

	return resp, err
}

func sendCustomMintRequest(token string, dto openapiclient.ServicesCustomMintDto) (*models.MintResp, error){
	//configuration := openapiclient.NewConfiguration()
	//apiClient := openapiclient.NewAPIClient(configuration)
	fmt.Println("Start to mint")
	resp, _, err := newClient().MintsApi.CustomMint(context.Background()).Authorization(token).CustomMintDto(dto).Execute()
	if err != nil {
		return nil, err
	}
	fmt.Println(resp)
	id, err := getTokenId(*resp.Id, token)
	if err != nil {
		return nil, err
	}

	res := &models.MintResp{
		UserAddress: dto.MintToAddress,
		NFTAddress: viper.GetString("customMint.mintRespPrefix") +  dto.ContractAddress + "/" + id,
		Contract: dto.ContractAddress,
		TokenID: id,
		Time: *resp.CreatedAt,
	}

	return res, nil
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
	fmt.Println(token)
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