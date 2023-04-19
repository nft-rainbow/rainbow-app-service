package models

type WalletUser struct {
	BaseModel
	Wallet        WalletType `gorm:"type:text" json:"wallet"`
	UnionId       string     `gorm:"type:varchar(256);index" json:"unionid"`
	AccessToken   string     `gorm:"type:text" json:"access_token"`
	Expire        int64      `gorm:"type:integer" json:"expire"` // access token expire time in timestamp
	RefreshToken  string     `gorm:"type:text" json:"refresh_token"`
	RefreshExpire int64      `gorm:"type:integer" json:"refresh_expire"` // refresh expire time in timestamp
	Phone         string     `gorm:"type:varchar(256);index" json:"phone"`
	Address       string     `gorm:"type:varchar(256);index" json:"address"`
}

func FindWalletUserByAddress(address string) ([]*WalletUser, error) {
	var users []*WalletUser
	err := GetDB().Where("address = ?", address).First(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func FindWalletUser(wallet WalletType, address string) (*WalletUser, error) {
	var user WalletUser
	err := GetDB().Where("wallet=? and address=?", wallet, address).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
