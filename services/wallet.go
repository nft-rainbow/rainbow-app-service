package services

import (
	"fmt"

	"github.com/nft-rainbow/rainbow-app-service/models/enums"
)

type AddWalletUserReq struct {
	Chain   enums.Chain      `json:"chain" binding:"required"`
	Wallet  enums.WalletType `json:"wallet" binding:"required"`
	Code    string           `json:"code" binding:"required"`
	Phone   string           `json:"phone"`
	Address string           `json:"address"`
}

type Wallet interface {
	InsertUser(user AddWalletUserReq) error
}

type WalletService struct {
	wallets map[enums.WalletType]Wallet
}

func NewWalletService() *WalletService {
	wallets := map[enums.WalletType]Wallet{
		enums.WALLET_ANYWEB: &Anyweb{},
		enums.WALLET_CELLAR: &Cellar{},
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
