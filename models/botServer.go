package models

// type DiscordCustomProjectConfig struct {
// 	BaseModel
// 	AppId         int32  `gorm:"index" json:"app_id" binding:"required"`
// 	GuildId       string `gorm:"type:varchar(256)" json:"guild_id" binding:"required"`
// 	GuildName     string `gorm:"type:varchar(256)" json:"guild_name"`
// 	RainbowUserId int32  `gorm:"type:integer" json:"rainbow_user_id"`
// 	ProjectName   string `gorm:"type:string" json:"Project_name" binding:"required"`
// 	Description   string `gorm:"type:string" json:"description" binding:"required"`
// 	ChainType     string `gorm:"type:string" json:"chain_type" binding:"required"`
// }

// dodo/discord 群组
type BotServer struct {
	BaseModel
	RainbowUserId uint           `gorm:"type:integer" json:"rainbow_user_id" binding:"required"`
	SocialTool    SocialToolType `json:"social_tool"`
	RawServerId   string         `json:"server_id" binding:"required"`
	OwnerSocialId string         `json:"user_social_id" binding:"required"`
	PushInfo      *PushInfo

	// Platform PlatformType `json:"platform" binding:"required"`
	// AppId         int32  `gorm:"index" json:"app_id" binding:"required"`
	// IslandId      string `gorm:"type:varchar(256)" json:"island_id" binding:"required"`
	// IslandName    string `gorm:"type:varchar(256)" json:"island_name"`
	// RainbowUserId int32 `gorm:"type:integer" json:"rainbow_user_id" binding:"required"`
	// ProjectName   string `gorm:"type:string" json:"Project_name" binding:"required"`
	// Description   string `gorm:"type:string" json:"description" binding:"required"`
	// ChainType     string `gorm:"type:string" json:"chain_type" binding:"required"`
	// PlatformUserId string `gorm:"type:varchar(255)" json:"platform_user_id" binding:"required"`
}

func FindBotServers(rainbowUserId uint, socialTool *SocialToolType) ([]*BotServer, error) {
	cond := BotServer{RainbowUserId: rainbowUserId}
	if socialTool != nil {
		cond.SocialTool = *socialTool
	}
	var result []*BotServer
	err := GetDB().Debug().Where(&cond).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func FindBotServerByChannel(channelId string) (*BotServer, error) {
	var result *BotServer
	var pi PushInfo
	if err := db.Model(&PushInfo{}).Where("channel_id=?", channelId).First(&pi).Error; err != nil {
		return nil, err
	}
	err := db.Model(&BotServer{}).Preload("PushInfo").Where("id=?", pi.BotServerID).First(&result).Error
	return result, err
}

func FindBotServerByRawID(rawServerId string, socialTool *SocialToolType) (*BotServer, error) {
	cond := BotServer{
		RawServerId: rawServerId,
	}
	if socialTool != nil {
		cond.SocialTool = *socialTool
	}
	if err := db.Preload("PushInfo").Where(&cond).First(&cond).Error; err != nil {
		return nil, err
	}
	return &cond, nil
}

func FindBotServerById(id uint) (*BotServer, error) {
	var result *BotServer
	err := db.Preload("PushInfo").Where("id=?", id).Find(&result).Error
	return result, err
}

// type DiscordCustomActivityConfig struct {
// 	BaseModel
// 	ContractID      int32  `gorm:"type:integer" json:"contract_id" binding:"required"`
// 	ChannelID       string `gorm:"type:varchar(256)" json:"channel_id" binding:"required"`
// 	Amount          int32  `gorm:"type:integer" json:"amount" binding:"required"`
// 	MaxMintCount    int32  `gorm:"type:varchar(256)" json:"max_mint_count" binding:"required"`
// 	Event           string `gorm:"type:string" json:"event" binding:"required"`
// 	Name            string `gorm:"type:string" json:"name" binding:"required"`
// 	Description     string `gorm:"type:string" json:"description" binding:"required"`
// 	AppId           int32  `gorm:"index" json:"app_id"`
// 	ContractType    int32  `gorm:"type:int" json:"contract_type"`
// 	ContractAddress string `gorm:"type:string" json:"contract_address"`
// 	Chain           int32  `gorm:"type:int" json:"chain_type"`
// 	MetadataURI     string `gorm:"type:string" json:"metadata_uri"`
// }

// type CustomActivityConfig struct {
// 	BaseModel
// 	Platform        SocialToolType `json:"platform" binding:"required"`
// 	GroupID         string         `gorm:"type:varchar(256)" json:"group_id" binding:"required"`
// 	ChannelID       string         `gorm:"type:varchar(256)" json:"channel_id" binding:"required"`
// 	ContractID      int32          `gorm:"type:integer" json:"contract_id" binding:"required"`
// 	Amount          int32          `gorm:"type:integer" json:"amount" binding:"required"`
// 	MaxMintCount    int32          `gorm:"type:varchar(256)" json:"max_mint_count" binding:"required"`
// 	Event           string         `gorm:"type:string" json:"event" binding:"required"`
// 	Name            string         `gorm:"type:string" json:"name" binding:"required"`
// 	Description     string         `gorm:"type:string" json:"description" binding:"required"`
// 	AppId           uint           `gorm:"index" json:"app_id"`
// 	ContractType    int32          `gorm:"type:int" json:"contract_type"`
// 	ContractAddress string         `gorm:"type:string" json:"contract_address"`
// 	Chain           int32          `gorm:"type:int" json:"chain_type"`
// 	MetadataURI     string         `gorm:"type:string" json:"metadata_uri"`
// }

type PushInfo struct {
	BaseModel
	BotServerID uint
	Activity    POAPActivityConfig
	ChannelId   string `gorm:"type:string" json:"channel_id"`
	Roles       string `gorm:"type:string" json:"roles"`
	Content     string `gorm:"type:string" json:"content"`
	ColorTheme  string `gorm:"type:string" json:"color_theme"`
}

// type PushInfo struct {
// 	BaseModel
// ServerId      string  `gorm:"type:varchar(256);index" json:"server_id"`
// ServerName    string  `gorm:"type:varchar(256)" json:"server_name"`
// Activity  Activity `gorm:"type:string;index" json:"activity_id"`
// ChannelId string   `gorm:"type:string" json:"channel_id"`

// ActivityName  string  `gorm:"type:string" json:"activity_name"`
// AccountLimit  int     `gorm:"type:integer" json:"account_limit"`
// ContractID    *int32  `gorm:"type:integer" json:"contract_id"`
// EndedTime     int64   `gorm:"type:integer" json:"end_time"`
// StartedTime   int64   `gorm:"type:integer" json:"start_time"`
// ActivityType  uint    `gorm:"type:uint" json:"activity_type"`
// Bot           uint    `gorm:"type:integer" json:"bot"`
// Contract      *string `gorm:"type:string" json:"contract"`
// RainbowUserId int32   `gorm:"type:integer" json:"rainbow_user_id"`
// }

// type SocialToolServer struct {
// 	BaseModel
// 	ServerId      string `gorm:"type:varchar(256);index" json:"server_id" binding:"required"`
// 	ServerName    string `gorm:"type:varchar(256)" json:"server_name"`
// 	RainbowUserId int32  `gorm:"type:integer" json:"rainbow_user_id"`
// 	UserId        int32  `gorm:"type:integer" json:"user_id"`
// 	Bot           uint   `gorm:"type:integer" json:"bot" binding:"required"`
// }

// type PushInfoQueryResult struct {
// 	Count int64       `json:"count"`
// 	Items []*PushInfo `json:"items"`
// }

type UserServerQueryResult struct {
	Count int64        `json:"count"`
	Items []*BotServer `json:"items"`
}

// type DiscordActivityQueryResult struct {
// 	Count int64                   `json:"count"`
// 	Items []*CustomActivityConfig `json:"items"`
// }

// type DoDoActivityQueryResult struct {
// 	Count int64                   `json:"count"`
// 	Items []*CustomActivityConfig `json:"items"`
// }

type DiscordCustomProjectConfigQueryResult struct {
	Count int64        `json:"count"`
	Items []*BotServer `json:"items"`
}

type DoDoCustomProjectConfigQueryResult struct {
	Count int64        `json:"count"`
	Items []*BotServer `json:"items"`
}

// func FindPushInfo(serverId, activityId string) (*PushInfo, error) {
// 	var res PushInfo
// 	var cond PushInfo
// 	cond.ActivityId = activityId
// 	cond.ServerId = serverId

// 	err := db.Where(&cond).Last(&res).Error
// 	return &res, err
// }

// func FindAndCountPushInfo(offset, limit, userId int, bot uint) (*PushInfoQueryResult, error) {
// 	var items []*PushInfo
// 	var cond PushInfo
// 	cond.RainbowUserId = int32(userId)
// 	cond.Bot = bot

// 	var count int64
// 	if err := db.Find(&items).Where(cond).Count(&count).Error; err != nil {
// 		return nil, err
// 	}

// 	if err := db.Find(&items).Where(cond).Offset(offset).Limit(limit).Error; err != nil {
// 		return nil, err
// 	}
// 	return &PushInfoQueryResult{count, items}, nil
// }

// func FindDiscordCustomActivityConfigByChannelId(id string) (*CustomActivityConfig, error) {
// 	var item CustomActivityConfig
// 	err := db.Where("channel_id = ?", id).First(&item).Error
// 	return &item, err
// }

// func FindDoDoCustomActivityConfigByChannelId(id string) (*CustomActivityConfig, error) {
// 	var item CustomActivityConfig
// 	err := db.Where("channel_id = ?", id).First(&item).Error
// 	return &item, err
// }

// func FindDiscordConfigById(id int) (*SocialToolServer, error) {
// 	var item SocialToolServer
// 	err := db.Where("id = ?", id).First(&item).Error
// 	return &item, err
// }

// func FindDiscordConfigByUserId(id int) (*SocialToolServer, error) {
// 	var item SocialToolServer
// 	err := db.Where("rainbow_user_id = ?", id).First(&item).Error
// 	return &item, err
// }

// func FindDoDoConfigById(id int) (*SocialToolServer, error) {
// 	var item SocialToolServer
// 	err := db.Where("id = ?", id).First(&item).Error
// 	return &item, err
// }

func FirstBotServerByUserId(rainbowUserId int) (*BotServer, error) {
	var item BotServer
	err := db.Preload("PushInfo").Where("rainbow_user_id = ?", rainbowUserId).First(&item).Error
	return &item, err
}

// func FindDiscordCustomActivityConfigById(id int) (*CustomActivityConfig, error) {
// 	var item CustomActivityConfig
// 	err := db.Where("id = ?", id).First(&item).Error
// 	return &item, err
// }

// func FindDoDoCustomActivityConfigById(id int) (*CustomActivityConfig, error) {
// 	var item CustomActivityConfig
// 	err := db.Where("id = ?", id).First(&item).Error
// 	return &item, err
// }

// func FindAndCountDiscordActivity(id uint, offset int, limit int) (*DiscordActivityQueryResult, error) {
// 	var items []*CustomActivityConfig
// 	cond := &CustomActivityConfig{}
// 	cond.AppId = id

// 	var count int64
// 	if err := db.Find(&items).Where(cond).Count(&count).Error; err != nil {
// 		return nil, err
// 	}

// 	if err := db.Find(&items).Where(cond).Offset(offset).Limit(limit).Error; err != nil {
// 		return nil, err
// 	}

// 	return &DiscordActivityQueryResult{count, items}, nil
// }

// func FindAndCountDoDoActivity(id uint, offset int, limit int) (*DoDoActivityQueryResult, error) {
// 	var items []*CustomActivityConfig
// 	cond := &CustomActivityConfig{}
// 	cond.AppId = id

// 	var count int64
// 	if err := db.Find(&items).Where(cond).Count(&count).Error; err != nil {
// 		return nil, err
// 	}

// 	if err := db.Find(&items).Where(cond).Offset(offset).Limit(limit).Error; err != nil {
// 		return nil, err
// 	}

// 	return &DoDoActivityQueryResult{count, items}, nil
// }

func FindAndCountDiscordCustomProjectConfig(id uint, offset int, limit int) (*DiscordCustomProjectConfigQueryResult, error) {
	var items []*BotServer
	cond := &BotServer{}
	cond.RainbowUserId = id

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
	var items []*BotServer
	cond := &BotServer{}
	cond.RainbowUserId = id

	var count int64
	if err := db.Find(&items).Where(cond).Count(&count).Error; err != nil {
		return nil, err
	}

	if err := db.Find(&items).Where(cond).Offset(offset).Limit(limit).Error; err != nil {
		return nil, err
	}

	return &DoDoCustomProjectConfigQueryResult{count, items}, nil
}
