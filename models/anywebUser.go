package models

type AnywebUser struct {
	BaseModel
	UnionId       string `gorm:"type:varchar(256);index" json:"unionid"`
	AccessToken   string `gorm:"type:text" json:"access_token"`
	Expire        int64  `gorm:"type:integer" json:"expire"` // access token expire time in timestamp
	RefreshToken  string `gorm:"type:text" json:"refresh_token"`
	RefreshExpire int64  `gorm:"type:integer" json:"refresh_expire"` // refresh expire time in timestamp
	Phone         string `gorm:"type:varchar(256);index" json:"phone"`
	Address       string `gorm:"type:varchar(256);index" json:"address"`
}

func FindAnywebUserByAddress(address string) (*AnywebUser, error) {
	var user AnywebUser
	err := GetDB().Where("address = ?", address).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
