package clients

import (
	"fmt"

	"github.com/imroc/req/v3"
	"github.com/spf13/viper"
)

var client = req.C() //.DevMode() // http request client

type AnywebAccessTokenReq struct {
	Appid  string `json:"appid"`
	Secret string `json:"secret"`
	Code   string `json:"code"`
}

type AnywebAccessTokenMeta struct {
	UnionId       string `json:"unionid"`
	AccessToken   string `json:"accessToken"`
	Expire        int64  `json:"expire"`
	RefreshToken  string `json:"refreshToken"`
	RefreshExpire int64  `json:"refreshExpire"`
	Scop          string `json:"scope"`
}

type AnywebAccessTokenResponse struct {
	Code    interface{}           `json:"code"`
	Message string                `json:"message"`
	Data    AnywebAccessTokenMeta `json:"data"`
}

type AnywebRefreshTokenReq struct {
	RefreshToken string `json:"refreshToken"`
}

type AnywebRefreshTokenMeta struct {
	Token         string `json:"token"`
	Expire        uint64 `json:"expire"`
	RefreshToken  string `json:"refreshToken"`
	RefreshExpire uint64 `json:"refreshExpire"`
}

type AnywebRefreshTokenResponse struct {
	Code    interface{}            `json:"code"`
	Message string                 `json:"message"`
	Data    AnywebRefreshTokenMeta `json:"data"`
}

type AnywebUserInfoReq struct {
	Appid       string   `json:"appid"`
	Secret      string   `json:"secret"`
	AccessToken string   `json:"accessToken"`
	UnionId     string   `json:"unionid"`
	Scopes      []string `json:"scopes"` // baseInfo, identity
}

type AnywebUserInfo struct {
	UnionId     string   `json:"unionid"`
	AddressList []string `json:"addressList"`
	Network     uint     `json:"network"`
	Scopes      []string `json:"scopes"`
	Level       uint     `json:"level"`
	Phone       string   `json:"phone"`
	IdNumber    *string  `json:"idNumber"`
	IdName      *string  `json:"name"`
}

type AnywebUserInfoResponse struct {
	Code    interface{}    `json:"code"`
	Message string         `json:"message"`
	Data    AnywebUserInfo `json:"data"`
}

func GetAnywebAccessToken(code string) (*AnywebAccessTokenMeta, error) {
	var result AnywebAccessTokenResponse

	resp, err := client.R().
		SetBody(&AnywebAccessTokenReq{Appid: viper.GetString("anyweb.appid"), Secret: viper.GetString("anyweb.secret"), Code: code}).
		SetSuccessResult(&result).
		Post("https://api.anyweb.cc/oauth/accessToken")

	if err != nil {
		return nil, err
	}
	if !resp.IsSuccessState() {
		return nil, fmt.Errorf("bad response status: %s", resp.Status)
	}
	if val, ok := result.Code.(string); ok {
		return nil, fmt.Errorf("API call failed: %s %s", val, result.Message)
	}
	return &result.Data, nil
}

func GetAnywebRefreshToken(refreshToken string) (*AnywebRefreshTokenMeta, error) {
	var result AnywebRefreshTokenResponse

	resp, err := client.R().
		SetBody(&AnywebRefreshTokenReq{RefreshToken: refreshToken}).
		SetSuccessResult(&result).
		Post("https://api.anyweb.cc/open/refreshToken")

	if err != nil {
		return nil, err
	}
	if !resp.IsSuccessState() {
		return nil, fmt.Errorf("bad response status: %s", resp.Status)
	}
	if val, ok := result.Code.(string); ok {
		return nil, fmt.Errorf("API call failed: %s %s", val, result.Message)
	}
	return &result.Data, nil
}

func GetAnywebUserInfo(accessToken string, unionid string, scopes []string) (*AnywebUserInfo, error) {
	var result AnywebUserInfoResponse

	resp, err := client.R().
		SetBody(&AnywebUserInfoReq{Appid: viper.GetString("anyweb.appid"), Secret: viper.GetString("anyweb.secret"), AccessToken: accessToken, UnionId: unionid, Scopes: scopes}).
		SetSuccessResult(&result).
		Post("https://api.anyweb.cc/oauth/userInfo")

	if err != nil {
		return nil, err
	}
	if !resp.IsSuccessState() {
		return nil, fmt.Errorf("bad response status: %s", resp.Status)
	}
	if val, ok := result.Code.(string); ok {
		return nil, fmt.Errorf("API call failed: %s %s", val, result.Message)
	}
	return &result.Data, nil
}
