package models

import "github.com/nft-rainbow/rainbow-app-service/models/enums"

type WalletUser struct {
	BaseModel
	Wallet        enums.WalletType `gorm:"type:varchar(256);index:idx_wallet_phone,priority:2" json:"wallet"`
	UnionId       string           `gorm:"type:varchar(256);index" json:"unionid"`
	AccessToken   string           `gorm:"type:text" json:"access_token"`
	Expire        int64            `gorm:"type:integer" json:"expire"` // access token expire time in timestamp
	RefreshToken  string           `gorm:"type:text" json:"refresh_token"`
	RefreshExpire int64            `gorm:"type:integer" json:"refresh_expire"` // refresh expire time in timestamp
	Phone         string           `gorm:"type:varchar(256);index:idx_wallet_phone,priority:1" json:"phone"`
	Address       string           `gorm:"type:varchar(256);index" json:"address"`
}

func FindWalletUserByAddress(address string) ([]*WalletUser, error) {
	var users []*WalletUser
	err := GetDB().Where("address = ?", address).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func FindWalletUser(wallet enums.WalletType, address string) (walletUser *WalletUser, err error) {
	err = GetDB().Where("wallet=? and address=?", wallet, address).First(&walletUser).Error
	if err != nil {
		return nil, err
	}
	return
}

func FindAllUserAddrsOfPhone(phone string) (addrs []string, err error) {
	err = GetDB().Debug().Model(&WalletUser{}).Select("address").Distinct().Where("phone=?", phone).Distinct("address").Find(&addrs).Error
	return
}
