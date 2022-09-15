package models

type AdminConfig struct {
	BaseModel
	AppId string `gorm:"type:varchar(256)" json:"app_id"`
	ChannelId string `gorm:"type:varchar(256)" json:"channel_id"`
	UserId string `gorm:"type:varchar(256)" json:"user_id"`
	Token string `gorm:"type:varchar(256)" json:"token"`
}

type CustomMintConfig struct {
	BaseModel
	FileUrl string `gorm:"type:varchar(256)" json:"file_url" binding:"required"`
	Name string `gorm:"type:varchar(256)" json:"name" binding:"required"`
	Description string `gorm:"type:varchar(256)" json:"description" binding:"required"`
	ContractType string `gorm:"type:varchar(256)" json:"contract_type" binding:"required"`
	ContractAddress string `gorm:"type:varchar(256)" json:"contract_address" binding:"required"`
	ChannelID string `gorm:"type:varchar(256)" json:"channel_id" binding:"required"`
	MaxMintCount uint `gorm:"type:varchar(256)" json:"max_mint_count"`
}

type EasyMintConfig struct {
	BaseModel
	FileUrl string `gorm:"type:varchar(256)" json:"file_url" binding:"required"`
	Name string `gorm:"type:varchar(256)" json:"name" binding:"required"`
	Description string `gorm:"type:varchar(256)" json:"description" binding:"required"`
	ChannelID string `gorm:"type:varchar(256)" json:"channel_id" binding:"required"`
	MaxMintCount uint `gorm:"type:varchar(256)" json:"max_mint_count"`
}

func FindBindingTokenById(id string) (string, error) {
	var item AdminConfig
	err := db.Where("user_id = ?", id).First(&item).Error
	return item.Token, err
}

func FindBindingCustomMintConfigById(id string) (*CustomMintConfig, error){
	var item CustomMintConfig
	err := db.Where("channel_id = ?", id).First(&item).Error
	return &item, err
}

func FindBindingEasyMintConfigById(id string) (*EasyMintConfig, error){
	var item EasyMintConfig
	err := db.Where("channel_id = ?", id).First(&item).Error
	return &item, err
}

func FindBindingConfigById(id string) (*AdminConfig, error) {
	var item AdminConfig
	err := db.Where("user_id = ?", id).First(&item).Error
	return &item, err
}

