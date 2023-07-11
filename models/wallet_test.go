package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
	"testing"

	"github.com/nft-rainbow/rainbow-app-service/models/enums"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestUnmarshalWalletType(t *testing.T) {
	var w *enums.WalletType
	err := json.Unmarshal([]byte("\"cellar\""), &w)
	assert.NoError(t, err)
	assert.Equal(t, enums.WALLET_CELLAR, *w)
	fmt.Println(*w)

	type AddWalletUserReq struct {
		Wallet  enums.WalletType `json:"wallet"`
		Code    string           `json:"code"`
		Phone   string           `json:"phone"`
		Address string           `json:"address"`
	}
	var req *AddWalletUserReq
	err = json.Unmarshal([]byte(`{
		"wallet":"cellar",
		"phone": "17011112222",
		"address": "cfx:aamgvyzht7h1zxdghb9ee9w26wrz8rd3gj837392dp"
	}`), &req)
	assert.NoError(t, err)
	assert.Equal(t, enums.WALLET_CELLAR, req.Wallet)
	fmt.Println(*req)
}

func TestCellarAutoGenAccount(t *testing.T) {
	type (
		LoginReq struct {
			UserPhone string `json:"userPhone"`
		}

		LoginResp struct {
			Msg  string `json:"msg"`
			Code int    `json:"code"`
			Data struct {
				Wallet    string `json:"wallet"`
				UserPhone string `json:"userPhone"`
				UserCode  string `json:"userCode"`
			} `json:"data"`
		}
	)

	initConfig()
	connectDB()
	getOrCreateCellarAccount := func(phone string) (*LoginResp, error) {
		j, _ := json.Marshal(LoginReq{
			UserPhone: phone,
		})
		url := "https://wallet.metacellar.art/web3/open/api/loginUnify"
		req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(j))
		if err != nil {
			return nil, errors.WithMessage(err, "failed to new request")
		}
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("appId", "3e438a5834cb43159fa45aad62671333")
		// req.Header.Add("User-Agent", "Apifox/1.0.0 (https://www.apifox.cn)")

		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return nil, errors.WithMessage(err, "failed to send request")
		}
		defer resp.Body.Close()

		respStr, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, errors.WithMessage(err, "failed to read resp body")
		}
		// fmt.Printf("%s", respStr)

		var loginResp LoginResp
		err = json.Unmarshal(respStr, &loginResp)
		if err != nil {
			return nil, errors.New(string(respStr))
		}

		return &loginResp, nil
	}

	phones := []string{"15116936766", "18035158224"}
	onebatch := 2
	for i := 0; i < len(phones); i = i + onebatch {
		max := i + onebatch
		if len(phones) < max {
			max = len(phones)
		}

		fragment := phones[i:max]
		var w sync.WaitGroup
		for _, phone := range fragment {
			w.Add(1)
			go func(_phone string) {
				defer w.Done()
				loginResp, err := getOrCreateCellarAccount(_phone)
				if err != nil {
					fmt.Printf("%s fail %v", _phone, err)
					return
				}
				fmt.Printf("%s %s\n", loginResp.Data.UserPhone, loginResp.Data.Wallet)

				var count int64
				if err := GetDB().Model(&WalletUser{}).Where("phone=? and address=?", _phone, loginResp.Data.Wallet).Count(&count).Error; err != nil {
					fmt.Printf("%s find db error %v", _phone, err)
					return
				}
				if count > 0 {
					return
				}

				if err := GetDB().Save(&WalletUser{
					Phone:   loginResp.Data.UserPhone,
					Address: loginResp.Data.Wallet,
				}).Error; err != nil {
					fmt.Printf("%s save db error %v", _phone, err)
					return
				}
			}(phone)
		}
		w.Wait()
		fmt.Printf("%d one batch done\n", i)
	}
}

func TestFindWalletAddress(t *testing.T) {
	initConfig()
	connectDB()

	phones := []string{"13801002119", "15000213993"}
	notSetted := []string{}
	for _, phone := range phones {
		var address string
		if err := GetDB().Model(&WalletUser{}).Where("phone=?", phone).Select("address").Scan(&address).Error; err != nil {
			fmt.Printf("%s \n", phone)
			continue
		}
		fmt.Printf("%s %s\n", phone, address)
		if address == "" {
			notSetted = append(notSetted, phone)
		}
	}
	fmt.Printf("not setted %v %v\n", len(notSetted), strings.Join(notSetted, "\",\""))
}

func initConfig() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	viper.AddConfigPath("..")     // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		log.Fatalln(fmt.Errorf("fatal error config file: %w", err))
	}
}
