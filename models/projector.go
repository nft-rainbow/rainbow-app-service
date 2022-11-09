package models

type DiscordAdminConfig struct {
	BaseModel
	AppId int32 `gorm:"index" json:"app_id" binding:"required"`
	GuildId string `gorm:"type:varchar(256)" json:"guild_id" binding:"required"`
	GuildName string `gorm:"type:varchar(256)" json:"guild_name"`
	RainbowUserId int32 `gorm:"type:integer" json:"rainbow_user_id"`
}

type DoDoAdminConfig struct {
	BaseModel
	AppId int32 `gorm:"index" json:"app_id"`
	IslandId string `gorm:"type:varchar(256)" json:"island_id" binding:"required"`
	IslandName string `gorm:"type:varchar(256)" json:"island_name"`
	RainbowUserId int32 `gorm:"type:integer" json:"rainbow_user_id"`
}

type DiscordActivityConfig struct {
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

type DoDoActivityConfig struct {
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

type DiscordActivityQueryResult struct {
	Count int64       `json:"count"`
	Items []*DiscordActivityConfig `json:"items"`
}

type DoDoActivityQueryResult struct {
	Count int64       `json:"count"`
	Items []*DoDoActivityConfig `json:"items"`
}

type DiscordAdminConfigQueryResult struct {
	Count int64       `json:"count"`
	Items []*DiscordAdminConfig `json:"items"`
}

type DoDoAdminConfigQueryResult struct {
	Count int64       `json:"count"`
	Items []*DoDoAdminConfig `json:"items"`
}

func FindBindingDiscordActivityConfigByChannelId(id string) (*DiscordActivityConfig, error){
	var item DiscordActivityConfig
	err := db.Where("channel_id = ?", id).First(&item).Error
	return &item, err
}

func FindBindingDoDoActivityConfigByChannelId(id string) (*DoDoActivityConfig, error){
	var item DoDoActivityConfig
	err := db.Where("channel_id = ?", id).First(&item).Error
	return &item, err
}

func FindBindingDiscordConfigById(id int) (*DiscordAdminConfig, error) {
	var item DiscordAdminConfig
	err := db.Where("id = ?", id).First(&item).Error
	return &item, err
}

func FindBindingDoDoConfigById(id int) (*DoDoAdminConfig, error) {
	var item DoDoAdminConfig
	err := db.Where("id = ?", id).First(&item).Error
	return &item, err
}

func FindBindingDiscordActivityConfigById(id int) (*DiscordActivityConfig, error){
	var item DiscordActivityConfig
	err := db.Where("id = ?", id).First(&item).Error
	return &item, err
}

func FindBindingDoDoActivityConfigById(id int) (*DoDoActivityConfig, error){
	var item DoDoActivityConfig
	err := db.Where("id = ?", id).First(&item).Error
	return &item, err
}

func FindAndCountDiscordActivity(id uint, offset int, limit int) (*DiscordActivityQueryResult, error) {
	var items []*DiscordActivityConfig
	cond := &DiscordActivityConfig{}
	cond.AppId = int32(id)

	var count int64
	if err := db.Find(&items).Where(cond).Count(&count).Error; err != nil {
		return nil, err
	}

	if err := db.Find(&items).Where(cond).Offset(offset).Limit(limit).Error; err != nil {
		return nil, err
	}

	return &DiscordActivityQueryResult{count, items}, nil
}

func FindAndCountDoDoActivity(id uint, offset int, limit int) (*DoDoActivityQueryResult, error) {
	var items []*DoDoActivityConfig
	cond := &DoDoActivityConfig{}
	cond.AppId = int32(id)

	var count int64
	if err := db.Find(&items).Where(cond).Count(&count).Error; err != nil {
		return nil, err
	}

	if err := db.Find(&items).Where(cond).Offset(offset).Limit(limit).Error; err != nil {
		return nil, err
	}

	return &DoDoActivityQueryResult{count, items}, nil
}

func FindAndCountDiscordAdminConfig(id uint, offset int, limit int) (*DiscordAdminConfigQueryResult, error) {
	var items []*DiscordAdminConfig
	cond := &DiscordAdminConfig{}
	cond.RainbowUserId = int32(id)

	var count int64
	if err := db.Find(&items).Where(cond).Count(&count).Error; err != nil {
		return nil, err
	}

	if err := db.Find(&items).Where(cond).Offset(offset).Limit(limit).Error; err != nil {
		return nil, err
	}

	return &DiscordAdminConfigQueryResult{count, items}, nil
}

func FindAndCountDoDoAdminConfig(id uint, offset int, limit int) (*DoDoAdminConfigQueryResult, error) {
	var items []*DoDoAdminConfig
	cond := &DoDoAdminConfig{}
	cond.RainbowUserId = int32(id)

	var count int64
	if err := db.Find(&items).Where(cond).Count(&count).Error; err != nil {
		return nil, err
	}

	if err := db.Find(&items).Where(cond).Offset(offset).Limit(limit).Error; err != nil {
		return nil, err
	}

	return &DoDoAdminConfigQueryResult{count, items}, nil
}

