package services

import (
	"time"

	"github.com/nft-rainbow/rainbow-app-service/clients"
	"github.com/nft-rainbow/rainbow-app-service/models"
)

func GetAnywebUserInfo(address, code string) error {
	// check exist in db first if have directly return
	user, err := models.FindAnywebUserByAddress(address)
	if err == nil && user != nil {
		// user exist
		return nil
	}

	// retrieve accessToken through code
	tokenInfo, err := clients.GetAnywebAccessToken(code)
	if err != nil {
		return err
	}

	// get userInfo through accessToken
	userInfo, err := clients.GetAnywebUserInfo(tokenInfo.AccessToken, tokenInfo.UnionId, []string{"baseInfo"})
	if err != nil {
		return err
	}

	// save db
	now := time.Now().Unix()
	anywebUser := &models.AnywebUser{
		UnionId:       tokenInfo.UnionId,
		AccessToken:   tokenInfo.AccessToken,
		Expire:        now + tokenInfo.Expire,
		RefreshToken:  tokenInfo.RefreshToken,
		RefreshExpire: now + tokenInfo.RefreshExpire,
		Phone:         userInfo.Phone,
		Address:       address,
	}

	err = models.GetDB().Create(anywebUser).Error
	return err
}
