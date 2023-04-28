package services

import (
	"errors"

	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/models/enums"
	"gorm.io/gorm"
)

type Cellar struct {
}

func (a *Cellar) InsertUser(userReq AddWalletUserReq) error {
	if userReq.Wallet != enums.WALLET_CELLAR {
		return errors.New("not cellar wallet")
	}

	user, err := models.FindWalletUser(userReq.Wallet, userReq.Address)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	if user != nil {
		return errors.New("user exists already")
	}

	// if user == nil {
	user = &models.WalletUser{
		Wallet:  userReq.Wallet,
		Address: userReq.Address,
		Phone:   userReq.Phone,
	}
	// }

	return models.GetDB().Save(user).Error
}
