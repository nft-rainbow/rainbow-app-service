package models

type AdminConfig struct {
	BaseModel
	AppId string `gorm:"type:varchar(256)" json:"app_id"`
	GuildId string `gorm:"type:varchar(256)" json:"guild_id"`
	GuildName string `gorm:"type:varchar(256)" json:"guild_name"`
}

type CustomMintConfig struct {
	BaseModel
	ContractID int32 `gorm:"type:integer" json:"contract_id" binding:"required"`
	ChannelID string `gorm:"type:varchar(256)" json:"channel_id" binding:"required"`
	Amount int32 `gorm:"type:integer" json:"amount" binding:"required"`
	MaxMintCount uint `gorm:"type:varchar(256)" json:"max_mint_count"`
	Event string `gorm:"type:string" json:"event"`
}


func FindBindingCustomMintConfigById(id string) (*CustomMintConfig, error){
	var item CustomMintConfig
	err := db.Where("channel_id = ?", id).First(&item).Error
	return &item, err
}

func FindBindingConfigById(id string) (*AdminConfig, error) {
	var item AdminConfig
	err := db.Where("user_id = ?", id).First(&item).Error
	return &item, err
}

