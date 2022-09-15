package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/nft-rainbow/discordbot-service/models"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func CustomMintConfig(config *models.CustomMintConfig) error {
	if config.MaxMintCount == 0 {
		config.MaxMintCount = 1
	}
	res := models.GetDB().Create(&config)
	if res.Error != nil {
		return  res.Error
	}
	return nil
}

func EasyMintMintConfig(config *models.EasyMintConfig) error {
	if config.MaxMintCount == 0 {
		config.MaxMintCount = 1
	}
	res := models.GetDB().Create(&config)
	if res.Error != nil {
		return  res.Error
	}
	return nil
}

func CustomMint(req *models.MintReq) (*models.MintResp, error){
	config, err := models.FindBindingCustomMintConfigById(req.ChannelID)
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
	//token, err := models.FindBindingTokenById(req.UserID)
	if err != nil {
		return nil, err
	}

	metadataUri, err := createMetadata(token, config.FileUrl, config.Name, config.Description)
	if err != nil {
		return nil, err
	}

	resp , err := sendCustomMintRequest(token, models.CustomMintDto{
		models.ContractInfoDto{
			Chain: viper.GetString("chainType"),
			ContractType: config.ContractType,
			ContractAddress: config.ContractAddress,
		},
		models.MintItemDto{
			MintToAddress: tmp.UserAddress,
			MetadataUri: metadataUri,
		},
	})
	if err != nil {
		return nil, err
	}
	_, err = models.UpdateCustomCount(req.UserID, req.ChannelID)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func EasyMint(req *models.MintReq) (*models.MintResp, error) {
	config, err := models.FindBindingEasyMintConfigById(req.ChannelID)
	if err != nil {
		return nil, err
	}

	ok, err := models.CheckEasyCount(req.UserID, req.ChannelID, config.MaxMintCount)
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
	//token, err := models.FindBindingTokenById(req.UserID)
	if err != nil {
		return nil, err
	}

	resp, err := sendEasyMintRequest(token, models.EasyMintMetaDto{
		Chain: viper.GetString("chainType"),
		Name: config.Name,
		Description: config.Description,
		MintToAddress: tmp.UserAddress,
		FileUrl: config.FileUrl,
	})
	if err != nil {
		return nil, err
	}
	_, err = models.UpdateEasyCount(req.UserID, req.ChannelID)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func sendCustomMintRequest(token string, dto models.CustomMintDto) (*models.MintResp, error){
	b, err := json.Marshal(dto)
	if err != nil {
		return nil, err
	}


	fmt.Println("Start to custom mint")
	req, _ := http.NewRequest("POST", viper.GetString("host") + "v1/mints/", bytes.NewBuffer(b))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer " + token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return  nil, err
	}

	var tmp models.MintTask
	content, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(content, &tmp)
	if err != nil {
		return nil, err
	}
	if tmp.ErrMessage != "" {
		return nil, errors.New(tmp.ErrMessage)
	}

	id, err := getTokenId(tmp.ID, token)
	if err != nil {
		return nil, err
	}

	res := &models.MintResp{
		UserAddress: dto.MintToAddress,
		NFTAddress: viper.GetString("customMint.mintRespPrefix") +  dto.ContractAddress + "/" + id,
		Contract: dto.ContractAddress,
		TokenID: id,
		Time: tmp.BaseModel.CreatedAt.String(),
	}

	defer resp.Body.Close()
	return res, nil
}

func sendEasyMintRequest(token string, dto models.EasyMintMetaDto) (*models.MintResp, error){
	b, err := json.Marshal(dto)
	if err != nil {
		return nil, err
	}
	fmt.Println("Start to easy mint")
	req, _ := http.NewRequest("POST", viper.GetString("host") + "v1/mints/easy/urls", bytes.NewBuffer(b))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer " + token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return  nil, err
	}

	var tmp models.MintTask
	content, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(content, &tmp)
	if err != nil {
		return nil, err
	}
	if tmp.ErrMessage != "" {
		return nil, errors.New(tmp.ErrMessage)
	}
	id, err := getTokenId(tmp.ID, token)
	if err != nil {
		return nil, err
	}

	res := &models.MintResp{
		UserAddress: dto.MintToAddress,
		Contract: viper.GetString("easyMint.contract"),
		NFTAddress: viper.GetString("easyMint.mintRespPrefix") + viper.GetString("easyMint.contract") + "/" + id,
		TokenID: id,
		Time: tmp.BaseModel.CreatedAt.String(),
	}

	defer resp.Body.Close()
	return res, nil
}

func createMetadata(token, fileUrl, name, description string) (string, error) {
	metadata := models.Metadata{
		Name: name,
		Description: description,
		Image: fileUrl,
	}
	fmt.Println(metadata)

	b, err := json.Marshal(metadata)
	if err != nil {
		return "", err
	}
	fmt.Println("Start to create metadata")
	req, _ := http.NewRequest("POST", viper.GetString("host") + "v1/metadata/", bytes.NewBuffer(b))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer " + token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
		return  "", err
	}

	var tmp models.CreateMetadataResponse
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(content, &tmp)
	if err != nil {
		return "", err
	}
	if tmp.Message != "" {
		return "", errors.New(tmp.Message)
	}

	return tmp.MetadataURI, nil
}

func getTokenId(id uint, token string) (string, error) {
	t := models.MintTask{}
	fmt.Println("Start to get token id")
	for t.TokenId == "" && t.Status != 1{
		req, err := http.NewRequest("GET", viper.GetString("host") + "v1/mints/" + strconv.Itoa(int(id)),nil)
		if err != nil {
			panic(err)
		}
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "Bearer " + token)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return "", err
		}
		content, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}

		err = json.Unmarshal(content, &t)
		if err != nil {
			return "", err
		}
		if t.Error != "" {
			return "", errors.New(t.Error)
		}
		time.Sleep(10 * time.Second)
	}
	return t.TokenId, nil
}

func Login() (string, error) {
	data := make(map[string]string)
	data["app_id"] = viper.GetString("app.appId")
	data["app_secret"] = viper.GetString("app.appSecret")
	b, _ := json.Marshal(data)
	fmt.Println("Start to login")
	req, err := http.NewRequest("POST", viper.GetString("host") + "v1/login", bytes.NewBuffer(b))
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return  "", err
	}
	content, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}
	t := make(map[string]interface{})
	err = json.Unmarshal(content, &t)
	if err != nil {
		return "", err
	}
	if t["code"] != nil {
		return "", errors.New(t["message"].(string))
	}

	return t["token"].(string), nil
}