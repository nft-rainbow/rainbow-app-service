package services

import (
	"github.com/nft-rainbow/rainbow-app-service/clients/cellar"
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/models/enums"
	"github.com/nft-rainbow/rainbow-app-service/utils/gormutils"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Cellar struct {
}

func (c *Cellar) GetOrCreateAccount(chain enums.Chain, phone string) (*models.WalletUser, error) {

	u, err := models.FindWalletUser(models.WalletUserFilter{Chain: chain, Phone: phone})
	if err == nil {
		return u, nil
	}

	if gormutils.IsRecordNotFoundError(err) {
		client := cellar.NewCellarClient(chain)
		resp, err := client.GetOrCreateAccount(phone)
		if err != nil {
			return nil, err
		}

		u = &models.WalletUser{
			Wallet:  enums.WALLET_CELLAR,
			Chain:   chain,
			Address: resp.Wallet,
			Phone:   resp.Phone,
		}
		if err = models.GetDB().Create(u).Error; err != nil {
			return nil, err
		}
		return u, nil
	}

	return nil, err
}

// create user by token
func (c *Cellar) InsertUser(userReq AddWalletUserReq) error {
	if userReq.Wallet != enums.WALLET_CELLAR {
		return errors.New("not cellar wallet")
	}

	client := cellar.NewCellarClient(userReq.Chain)
	cu, err := client.GetAccount(userReq.Code)
	if err != nil {
		return err
	}

	if userReq.Address != "" && userReq.Address != cu.Wallet {
		return errors.New("address mismatch")
	}

	if userReq.Phone != "" && userReq.Phone != cu.Phone {
		return errors.New("phone mismatch")
	}

	_, err = models.FindWalletUser(models.WalletUserFilter{Wallet: enums.WALLET_CELLAR, Address: cu.Wallet})
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
