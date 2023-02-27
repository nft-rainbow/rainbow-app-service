package models

type DiscordCustomProjectConfig struct {
	BaseModel
	AppId         int32  `gorm:"index" json:"app_id" binding:"required"`
	GuildId       string `gorm:"type:varchar(256)" json:"guild_id" binding:"required"`
	GuildName     string `gorm:"type:varchar(256)" json:"guild_name"`
	RainbowUserId int32  `gorm:"type:integer" json:"rainbow_user_id"`
	ProjectName   string `gorm:"type:string" json:"Project_name" binding:"required"`
	Description   string `gorm:"type:string" json:"description" binding:"required"`
	ChainType     string `gorm:"type:string" json:"chain_type" binding:"required"`
}

type DoDoCustomProjectConfig struct {
	BaseModel
	AppId         int32  `gorm:"index" json:"app_id" binding:"required"`
	IslandId      string `gorm:"type:varchar(256)" json:"island_id" binding:"required"`
	IslandName    string `gorm:"type:varchar(256)" json:"island_name"`
	RainbowUserId int32  `gorm:"type:integer" json:"rainbow_user_id"`
	ProjectName   string `gorm:"type:string" json:"Project_name" binding:"required"`
	Description   string `gorm:"type:string" json:"description" binding:"required"`
	ChainType     string `gorm:"type:string" json:"chain_type" binding:"required"`
}

type DiscordCustomActivityConfig struct {
	BaseModel
	ContractID      int32  `gorm:"type:integer" json:"contract_id" binding:"required"`
	ChannelID       string `gorm:"type:varchar(256)" json:"channel_id" binding:"required"`
	Amount          int32  `gorm:"type:integer" json:"amount" binding:"required"`
	MaxMintCount    int32  `gorm:"type:varchar(256)" json:"max_mint_count" binding:"required"`
	Event           string `gorm:"type:string" json:"event" binding:"required"`
	Name            string `gorm:"type:string" json:"name" binding:"required"`
	Description     string `gorm:"type:string" json:"description" binding:"required"`
	AppId           int32  `gorm:"index" json:"app_id"`
	ContractType    int32  `gorm:"type:int" json:"contract_type"`
	ContractAddress string `gorm:"type:string" json:"contract_address"`
	Chain           int32  `gorm:"type:int" json:"chain_type"`
	MetadataURI     string `gorm:"type:string" json:"metadata_uri"`
}

type DoDoCustomActivityConfig struct {
	BaseModel
	ContractID      int32  `gorm:"type:integer" json:"contract_id" binding:"required"`
	ChannelID       string `gorm:"type:varchar(256)" json:"channel_id" binding:"required"`
	Amount          int32  `gorm:"type:integer" json:"amount" binding:"required"`
	MaxMintCount    int32  `gorm:"type:varchar(256)" json:"max_mint_count" binding:"required"`
	Event           string `gorm:"type:string" json:"event" binding:"required"`
	Name            string `gorm:"type:string" json:"name" binding:"required"`
	Description     string `gorm:"type:string" json:"description" binding:"required"`
	AppId           int32  `gorm:"index" json:"app_id"`
	ContractType    int32  `gorm:"type:int" json:"contract_type"`
	ContractAddress string `gorm:"type:string" json:"contract_address"`
	Chain           int32  `gorm:"type:int" json:"chain_type"`
	MetadataURI     string `gorm:"type:string" json:"metadata_uri"`
}

type PushReq struct {
	ServerId     string `gorm:"type:varchar(256)" json:"server_id" binding:"required"`
	ServerName   string `gorm:"type:varchar(256)" json:"server_name"`
	ChannelId    string `gorm:"type:varchar(256)" json:"channel_id"`
	Roles        string `gorm:"type:varchar(256)" json:"roles"`
	AccountLimit string `gorm:"type:varchar(256)" json:"account_limit"`
	Color        string `gorm:"type:varchar(256)" json:"color"`
	Content      string `gorm:"type:varchar(256)" json:"content"`
	Bot          uint   `gorm:"type:integer" json:"bot"`
	UserId       uint   `gorm:"type:integer" json:"user_id"`
	ActivityId   string `gorm:"type:string" json:"activity_id"`
}

type PushInfo struct {
	BaseModel
	ServerId     string `gorm:"type:varchar(256)" json:"server_id"`
	ServerName   string `gorm:"type:varchar(256)" json:"server_name"`
	ActivityId   string `gorm:"type:string" json:"activity_id"`
	ActivityName string `gorm:"type:string" json:"activity_name"`
	ContractID   int32  `gorm:"type:integer" json:"contract_id"`
	EndedTime    int64  `gorm:"type:integer" json:"end_time"`
	StartedTime  int64  `gorm:"type:integer" json:"start_time"`
	ActivityType uint   `gorm:"type:uint" json:"activity_type"`
	Contract     string `gorm:"type:string" json:"contract"`
}

type DiscordActivityQueryResult struct {
	Count int64                          `json:"count"`
	Items []*DiscordCustomActivityConfig `json:"items"`
}

type DoDoActivityQueryResult struct {
	Count int64                       `json:"count"`
	Items []*DoDoCustomActivityConfig `json:"items"`
}

type DiscordCustomProjectConfigQueryResult struct {
	Count int64                         `json:"count"`
	Items []*DiscordCustomProjectConfig `json:"items"`
}

type DoDoCustomProjectConfigQueryResult struct {
	Count int64                      `json:"count"`
	Items []*DoDoCustomProjectConfig `json:"items"`
}

func FindDiscordCustomActivityConfigByChannelId(id string) (*DiscordCustomActivityConfig, error) {
	var item DiscordCustomActivityConfig
	err := db.Where("channel_id = ?", id).First(&item).Error
	return &item, err
}

func FindDoDoCustomActivityConfigByChannelId(id string) (*DoDoCustomActivityConfig, error) {
	var item DoDoCustomActivityConfig
	err := db.Where("channel_id = ?", id).First(&item).Error
	return &item, err
}

func FindDiscordConfigById(id int) (*DiscordCustomProjectConfig, error) {
	var item DiscordCustomProjectConfig
	err := db.Where("id = ?", id).First(&item).Error
	return &item, err
}

func FindDiscordConfigByUserId(id int) (*DiscordCustomProjectConfig, error) {
	var item DiscordCustomProjectConfig
	err := db.Where("rainbow_user_id = ?", id).First(&item).Error
	return &item, err
}

func FindDoDoConfigById(id int) (*DoDoCustomProjectConfig, error) {
	var item DoDoCustomProjectConfig
	err := db.Where("id = ?", id).First(&item).Error
	return &item, err
}

func FindDoDoConfigByUserId(id int) (*DoDoCustomProjectConfig, error) {
	var item DoDoCustomProjectConfig
	err := db.Where("rainbow_user_id = ?", id).First(&item).Error
	return &item, err
}

func FindDiscordCustomActivityConfigById(id int) (*DiscordCustomActivityConfig, error) {
	var item DiscordCustomActivityConfig
	err := db.Where("id = ?", id).First(&item).Error
	return &item, err
}

func FindDoDoCustomActivityConfigById(id int) (*DoDoCustomActivityConfig, error) {
	var item DoDoCustomActivityConfig
	err := db.Where("id = ?", id).First(&item).Error
	return &item, err
}

func FindAndCountDiscordActivity(id uint, offset int, limit int) (*DiscordActivityQueryResult, error) {
	var items []*DiscordCustomActivityConfig
	cond := &DiscordCustomActivityConfig{}
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
	var items []*DoDoCustomActivityConfig
	cond := &DoDoCustomActivityConfig{}
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

func FindAndCountDiscordCustomProjectConfig(id uint, offset int, limit int) (*DiscordCustomProjectConfigQueryResult, error) {
	var items []*DiscordCustomProjectConfig
	cond := &DiscordCustomProjectConfig{}
	cond.RainbowUserId = int32(id)

	var count int64
	if err := db.Find(&items).Where(cond).Count(&count).Error; err != nil {
		return nil, err
	}

	if err := db.Find(&items).Where(cond).Offset(offset).Limit(limit).Error; err != nil {
		return nil, err
	}

	return &DiscordCustomProjectConfigQueryResult{count, items}, nil
}

func FindAndCountDoDoCustomProjectConfig(id uint, offset int, limit int) (*DoDoCustomProjectConfigQueryResult, error) {
	var items []*DoDoCustomProjectConfig
	cond := &DoDoCustomProjectConfig{}
	cond.RainbowUserId = int32(id)

	var count int64
	if err := db.Find(&items).Where(cond).Count(&count).Error; err != nil {
		return nil, err
	}

	if err := db.Find(&items).Where(cond).Offset(offset).Limit(limit).Error; err != nil {
		return nil, err
	}

	return &DoDoCustomProjectConfigQueryResult{count, items}, nil
}
