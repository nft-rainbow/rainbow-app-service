package models

type AdminConfig struct {
	BaseModel
	AppId int32 `gorm:"index" json:"app_id"`
	GuildId string `gorm:"type:varchar(256)" json:"guild_id" binding:"required"`
	GuildName string `gorm:"type:varchar(256)" json:"guild_name"`
	RainbowUserId int32 `gorm:"type:integer" json:"rainbow_user_id"`
}

type ActivityConfig struct {
	BaseModel
	ContractID int32 `gorm:"type:integer" json:"contract_id" binding:"required"`
	ChannelID string `gorm:"type:varchar(256)" json:"channel_id" binding:"required"`
	Amount int32 `gorm:"type:integer" json:"amount" binding:"required"`
	MaxMintCount uint `gorm:"type:varchar(256)" json:"max_mint_count" binding:"required"`
	Event string `gorm:"type:string" json:"event" binding:"required"`
	Name string `gorm:"type:string" json:"name" binding:"required"`
	Description string `gorm:"type:string" json:"description" binding:"required"`
	AppId int32 `gorm:"index" json:"app_id"`
	ContractType int32 `gorm:"type:int" json:"contract_type"`
	ContractAddress string `gorm:"type:string" json:"contract_address"`
	Chain    int32   `gorm:"type:int" json:"chain_type"`
	MetadataURI string `gorm:"type:string" json:"metadata_uri"`
}

type ActivityQueryResult struct {
	Count int64       `json:"count"`
	Items []*ActivityConfig `json:"items"`
}

type AdminConfigQueryResult struct {
	Count int64       `json:"count"`
	Items []*AdminConfig `json:"items"`
}

func FindBindingActivityConfigByChannelId(id string) (*ActivityConfig, error){
	var item ActivityConfig
	err := db.Where("channel_id = ?", id).First(&item).Error
	return &item, err
}

func FindBindingConfigById(id int) (*AdminConfig, error) {
	var item AdminConfig
	err := db.Where("id = ?", id).First(&item).Error
	return &item, err
}

func FindBindingActivityConfigById(id int) (*ActivityConfig, error){
	var item ActivityConfig
	err := db.Where("id = ?", id).First(&item).Error
	return &item, err
}

func FindAndCountActivity(id uint, offset int, limit int) (*ActivityQueryResult, error) {
	var items []*ActivityConfig
	cond := &ActivityConfig{}
	cond.AppId = int32(id)

	var count int64
	if err := db.Find(&items).Where(cond).Count(&count).Error; err != nil {
		return nil, err
	}

	if err := db.Find(&items).Where(cond).Offset(offset).Limit(limit).Error; err != nil {
		return nil, err
	}

	return &ActivityQueryResult{count, items}, nil
}

func FindAndCountAdminConfig(id uint, offset int, limit int) (*AdminConfigQueryResult, error) {
	var items []*AdminConfig
	cond := &AdminConfig{}
	cond.RainbowUserId = int32(id)

	var count int64
	if err := db.Find(&items).Where(cond).Count(&count).Error; err != nil {
		return nil, err
	}

	if err := db.Find(&items).Where(cond).Offset(offset).Limit(limit).Error; err != nil {
		return nil, err
	}

	return &AdminConfigQueryResult{count, items}, nil
}

