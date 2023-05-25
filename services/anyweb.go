package services

import (
	"errors"
	"time"

	"github.com/nft-rainbow/rainbow-app-service/clients"
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/models/enums"
)

type Anyweb struct {
}

func (a *Anyweb) InsertUser(userReq AddWalletUserReq) error {
	if userReq.Wallet != enums.WALLET_ANYWEB {
		return errors.New("not anyweb wallet")
	}
	// check exist in db first if have directly return
	user, err := models.FindWalletUser(userReq.Wallet, userReq.Address)
	if err == nil && user != nil {
		// user exist
		return nil
	}

	// retrieve accessToken through code
	tokenInfo, err := clients.GetAnywebAccessToken(userReq.Code)
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
	anywebUser := &models.WalletUser{
		UnionId:       tokenInfo.UnionId,
		AccessToken:   tokenInfo.AccessToken,
		Expire:        now + tokenInfo.Expire,
		RefreshToken:  tokenInfo.RefreshToken,
		RefreshExpire: now + tokenInfo.RefreshExpire,
		Phone:         userInfo.Phone,
		Address:       userReq.Address,
		Wallet:        userReq.Wallet,
	}

	err = models.GetDB().Create(anywebUser).Error
	return err
}
