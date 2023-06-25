package services

import (
	"bytes"
	"encoding/json"
	"fmt"

	"io/ioutil"
	"net/http"

	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/models/enums"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Cellar struct {
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

type cellarUser struct {
	Message string `json:"msg"`
	Code    int    `json:"code"`
	Data    struct {
		Phone  string `json:"phone"`
		Wallet string `json:"wallet"`
	} `json:"data"`
}

func (a *Cellar) getUserInfoByToken(token string) (*cellarUser, error) {
	url := "https://wallet.metacellar.art/web3/userAuth/getUserInfoByToken" //PROD
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

	var cu cellarUser
	if err := json.Unmarshal(body, &cu); err != nil {
		return nil, err
	}

	if cu.Code != http.StatusOK {
		return nil, errors.Errorf("failed, code: %v, message: %v", cu.Code, cu.Message)
	}

	return &cu, nil
}
