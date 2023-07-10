package services

import (
	"fmt"

	"net/http"

	"github.com/nft-rainbow/rainbow-app-service/config"
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/models/enums"
	"github.com/nft-rainbow/rainbow-app-service/utils"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Cellar struct {
	Host       string
	httpClient *http.Client
}

func NewCellarClient(chain enums.Chain) *Cellar {
	switch chain {
	case enums.CHAIN_CONFLUX:
		return &Cellar{"wallet.metacellar.art", &http.Client{}}
	default:
		return &Cellar{"wallet-pre.maytek.cn", &http.Client{}}
	}
}

func (a *Cellar) InsertUser(userReq AddWalletUserReq) error {
	if userReq.Wallet != enums.WALLET_CELLAR {
		return errors.New("not cellar wallet")
	}

	cu, err := a.getAccount(userReq.Code)
	if err != nil {
		return err
	}

	if userReq.Address != "" && userReq.Address != cu.Wallet {
		return errors.New("address mismatch")
	}

	if userReq.Phone != "" && userReq.Phone != cu.Phone {
		return errors.New("phone mismatch")
	}

	_, err = models.FindWalletUser(enums.WALLET_CELLAR, cu.Wallet)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			user := &models.WalletUser{
				Wallet:  enums.WALLET_CELLAR,
				Address: cu.Wallet,
				Phone:   cu.Phone,
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
type getCellarUserResp struct {
	Phone  string `json:"phone"`
	Wallet string `json:"wallet"`
}

type getOrCreateCellarUserResp struct {
	Phone  string `json:"userPhone"`
	Wallet string `json:"wallet"`
	Code   string `json:"userCode"`
}

func (a *Cellar) getAccount(token string) (*getCellarUserResp, error) {
	url := fmt.Sprintf("https://%s/web3/userAuth/getUserInfoByToken", a.Host)
	resp, err := requestCellar[any, getCellarUserResp](http.MethodPost, url,
		map[string]string{"token": token},
		map[string]string{"Content-Type": "application/json"})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *Cellar) getOrCreateAccount(phone string) (*getOrCreateCellarUserResp, error) {
	url := fmt.Sprintf("https://%s/web3/open/api/loginUnify", a.Host)
	resp, err := requestCellar[any, getOrCreateCellarUserResp](http.MethodPost, url,
		map[string]string{"userPhone": phone},
		map[string]string{
			"appId":        config.GetConfig().Wallet.Cellar.Appid,
			"Content-Type": "application/json",
		})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func requestCellar[TPayload any, TResp any](method string, url string, payload TPayload, headers map[string]string) (*TResp, error) {

	cr, err := utils.SendHttp[any, cellarResp[TResp]](method, url, payload, headers)
	if err != nil {
		return nil, err
	}

	if cr.Code != http.StatusOK {
		return nil, errors.Errorf("failed, code: %v, message: %v", cr.Code, cr.Message)
	}

	return cr.Data, nil
}
