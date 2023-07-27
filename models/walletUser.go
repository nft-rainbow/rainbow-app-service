package models

import (
	"fmt"

	"github.com/nft-rainbow/rainbow-app-service/models/enums"
	"github.com/nft-rainbow/rainbow-app-service/utils/gormutils"
	"gorm.io/gorm"
)

type WalletUser struct {
	BaseModel
	Wallet        enums.WalletType `gorm:"type:varchar(256);index:idx_wallet_phone,priority:2" json:"wallet"`
	Chain         enums.Chain      `gorm:"type:uint" json:"chain"`
	UnionId       string           `gorm:"type:varchar(256);index" json:"unionid"`
	AccessToken   string           `gorm:"type:text" json:"access_token"`
	Expire        int64            `gorm:"type:integer" json:"expire"` // access token expire time in timestamp
	RefreshToken  string           `gorm:"type:text" json:"refresh_token"`
	RefreshExpire int64            `gorm:"type:integer" json:"refresh_expire"` // refresh expire time in timestamp
	Phone         string           `gorm:"type:varchar(256);index:idx_wallet_phone,priority:1" json:"phone"`
	Address       string           `gorm:"type:varchar(256);index" json:"address"`
}

// If record not found, return nil user and nil error
func FindWalletUserByAddress(address string) (*WalletUser, error) {
	var user WalletUser
	err := GetDB().Where("address = ?", address).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

type WalletUserFilter struct {
	Wallet  enums.WalletType
	Chain   enums.Chain
	Phone   string
	Address string
}

func (w *WalletUserFilter) where() *gorm.DB {
	tx := GetDB()
	if int(w.Wallet) > 0 {
		tx = tx.Where("wallet=?", w.Wallet)
	}
	if int(w.Chain) > 0 {
		tx = tx.Where("chain=?", w.Chain)
	}
	if w.Phone != "" {
		tx = tx.Where("phone=?", w.Phone)
	}
	if w.Address != "" {
		tx = tx.Where("address=?", w.Address)
	}
	return tx
}

func FindWalletUser(filter WalletUserFilter) (walletUser *WalletUser, err error) {
	err = GetDB().Where(filter.where()).First(&walletUser).Error
	if err != nil {
		return nil, err
	}
	return
}

func FindTopWalletUsersByPhones(wallet enums.WalletType, chain enums.Chain, phones []string) (map[string]string, error) {
	type phoneXaddr struct {
		Phone   string `json:"phone"`
		Address string `json:"address"`
	}

	var phoneXaddrs []phoneXaddr
	// select phone,address from wallet_users where id in (select min(id) from wallet_users group by phone)
	getIdsSql := GetDB().Table("wallet_users").Select("min(id)").Group("phone").Where("wallet=? and chain=? and phone in ?", enums.WALLET_CELLAR, chain, phones)
	if err := GetDB().Debug().Table("wallet_users").Select("phone,address").Where("id in (?)", getIdsSql).Find(&phoneXaddrs).Error; err != nil {
		return nil, err
	}
	fmt.Println("find phone with address, len: ", len(phoneXaddrs))

	result := make(map[string]string)
	for _, pxa := range phoneXaddrs {

		result[pxa.Phone] = pxa.Address
	}
	return result, nil
}

func FindAllUserAddrsOfPhone(phone string) (addrs []string, err error) {
	err = GetDB().Debug().Model(&WalletUser{}).Select("address").Distinct().Where("phone=?", phone).Distinct("address").Find(&addrs).Error
	return
}

func FindRelatedAddressWithSamePhone(address string) (addrs []string, err error) {
	user, err := FindWalletUserByAddress(address)
	if err != nil {
		if gormutils.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return FindAllUserAddrsOfPhone(user.Phone)
}
