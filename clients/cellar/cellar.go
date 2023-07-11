package cellar

import (
	"fmt"

	"net/http"

	"github.com/nft-rainbow/rainbow-app-service/config"
	"github.com/nft-rainbow/rainbow-app-service/models/enums"
	"github.com/nft-rainbow/rainbow-app-service/utils"
	"github.com/pkg/errors"
)

type Cellar struct {
	httpClient *http.Client
	chain      enums.Chain
	config.Cellar
}

func NewCellarClient(chain enums.Chain) *Cellar {
	return &Cellar{
		httpClient: &http.Client{},
		chain:      chain,
		Cellar:     config.GetCellarByChain(chain),
	}
}

type cellarResp[T any] struct {
	Message string `json:"msg"`
	Code    int    `json:"code"`
	Data    *T
}
type GetCellarUserResp struct {
	Phone  string `json:"phone"`
	Wallet string `json:"wallet"`
}

type GetOrCreateCellarUserResp struct {
	Phone  string `json:"userPhone"`
	Wallet string `json:"wallet"`
	Code   string `json:"userCode"`
}

func (a *Cellar) GetAccount(token string) (*GetCellarUserResp, error) {
	url := fmt.Sprintf("https://%s/web3/userAuth/getUserInfoByToken", a.Host)
	resp, err := requestCellar[any, GetCellarUserResp](http.MethodPost, url,
		map[string]string{"token": token},
		map[string]string{"Content-Type": "application/json"})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *Cellar) GetOrCreateAccount(phone string) (*GetOrCreateCellarUserResp, error) {
	url := fmt.Sprintf("https://%s/web3/open/api/loginUnify", a.Host)
	resp, err := requestCellar[any, GetOrCreateCellarUserResp](http.MethodPost, url,
		map[string]string{"userPhone": phone},
		map[string]string{
			"appId":        a.Appid,
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
