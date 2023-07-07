package services

import (
	"bytes"
	"encoding/json"
	"fmt"

	"io/ioutil"
	"net/http"

	"github.com/nft-rainbow/rainbow-app-service/config"
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/models/enums"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Cellar struct {
	Host       string
	httpClient *http.Client
}

func NewCellarClient(host ...string) *Cellar {
	if len(host) == 0 {
		host = append(host, "wallet.metacellar.art")
	}
	return &Cellar{host[0], &http.Client{}}
}

func (a *Cellar) InsertUser(userReq AddWalletUserReq) error {
	if userReq.Wallet != enums.WALLET_CELLAR {
		return errors.New("not cellar wallet")
	}

	cu, err := a.getUserInfoByToken(userReq.Code)
	if err != nil {
		return err
	}

	if userReq.Address != "" && userReq.Address != cu.Data.Wallet {
		return errors.New("address mismatch")
	}

	if userReq.Phone != "" && userReq.Phone != cu.Data.Phone {
		return errors.New("phone mismatch")
	}

	_, err = models.FindWalletUser(enums.WALLET_CELLAR, cu.Data.Wallet)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			user := &models.WalletUser{
				Wallet:  enums.WALLET_CELLAR,
				Address: cu.Data.Wallet,
				Phone:   cu.Data.Phone,
			}
			return models.GetDB().Create(user).Error
		}
		return err
	}

	return errors.New("user exists already")
}

type cellarResp[T any] struct {
	Message string `json:"msg"`
	Code    int    `json:"code"`
	Data    *T
}
type getCellarUserResp cellarResp[struct {
	Phone  string `json:"phone"`
	Wallet string `json:"wallet"`
}]

type getOrCreateCellarUserResp cellarResp[struct {
	Phone  string `json:"userPhone"`
	Wallet string `json:"wallet"`
	Code   string `json:"userCode"`
}]

func (a *Cellar) getUserInfoByToken(token string) (*getCellarUserResp, error) {
	url := fmt.Sprintf("https://%s/web3/userAuth/getUserInfoByToken", a.Host) //PROD
	// url := "https://wallet-pre.maytek.cn/web3/userAuth/getUserInfoByToken" //DEV

	payload := map[string]string{"token": token}
	payloadJ, _ := json.Marshal(payload)

	fmt.Println(string(payloadJ))
	resp, err := http.Post(url, "application/json", bytes.NewReader(payloadJ))
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var cu getCellarUserResp
	if err := json.Unmarshal(body, &cu); err != nil {
		return nil, err
	}

	if cu.Code != http.StatusOK {
		return nil, errors.Errorf("failed, code: %v, message: %v", cu.Code, cu.Message)
	}

	return &cu, nil
}

func (a *Cellar) getOrCreateAccount(phone string) (*getOrCreateCellarUserResp, error) {
	url := fmt.Sprintf("https://%s/web3/open/api/loginUnify", a.Host)

	// payload := map[string]string{"userPhone": phone}
	// payloadJ, _ := json.Marshal(payload)

	// req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader(payloadJ))
	// req.Header.Set("appId", config.GetConfig().Wallet.Cellar.Appid)

	// resp, err := http.DefaultClient.Do(req)
	// if err != nil {
	// 	return nil, err
	// }

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return nil, err
	// }

	resp, err := sendHttp[any, getOrCreateCellarUserResp](http.MethodPost, url,
		map[string]string{"userPhone": phone},
		map[string]string{"appId": config.GetConfig().Wallet.Cellar.Appid})
	if err != nil {
		return nil, err
	}
	if resp.Code != http.StatusOK {
		return nil, errors.Errorf("failed, code: %v, message: %v", resp.Code, resp.Message)
	}
	return resp, nil
}

func sendHttp[TPayload any, TResp any](method string, url string, payload TPayload, headers map[string]string) (*TResp, error) {
	payloadJ, _ := json.Marshal(payload)

	req, _ := http.NewRequest(method, url, bytes.NewReader(payloadJ))
	// req.Header.Set("appId", config.GetConfig().Wallet.Cellar.Appid)

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var cu TResp
	if err := json.Unmarshal(body, &cu); err != nil {
		return nil, err
	}

	return &cu, nil
}
