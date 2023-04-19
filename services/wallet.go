package services

import (
	"fmt"

	"github.com/nft-rainbow/rainbow-app-service/models"
)

type AddWalletUserReq struct {
	Wallet  models.WalletType `json:"wallet"`
	Code    string            `json:"code"`
	Phone   string            `json:"phone"`
	Address string            `json:"address"`
}

type Wallet interface {
	InsertUser(user AddWalletUserReq) error
}

type WalletService struct {
	wallets map[models.WalletType]Wallet
}

func NewWalletService() *WalletService {
	wallets := map[models.WalletType]Wallet{
		models.WALLET_ANYWEB: &Anyweb{},
		models.WALLET_CELLAR: &Cellar{},
	}

	return &WalletService{wallets: wallets}
}

func (w *WalletService) InsertUser(user AddWalletUserReq) error {
	wallet, ok := w.wallets[user.Wallet]
	if !ok {
		return fmt.Errorf("not support %v", user.Wallet)
	}
	return wallet.InsertUser(user)
}
