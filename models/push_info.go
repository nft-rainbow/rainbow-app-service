package models

import (
	"github.com/pkg/errors"
)

type PushInfo struct {
	BaseModel
	BotServerID  uint      `json:"bot_server_id"`
	ChannelId    string    `gorm:"unique_index:idx_member;type:varchar(255)" json:"channel_id"`
	Roles        string    `gorm:"type:string" json:"roles"`
	Content      string    `gorm:"type:string" json:"content"`
	ColorTheme   string    `gorm:"type:string" json:"color_theme"`
	LastPushTime int64     `json:"last_push_time"`
	ActivityId   uint      `gorm:"unique_index:idx_member" json:"activity_id"`
	Activity     *Activity `gorm:"-" json:"activity"`
}

func (p *PushInfo) LoadActivity() error {
	var activity Activity
	activity.ID = p.ActivityId
	if err := GetDB().Where(&activity).First(&activity).Error; err != nil {
		return err
	}
	if err := activity.LoadBindedContract(); err != nil {
		return err
	}
	p.Activity = &activity
	return nil
}

func IsPushInfoExists(activityId uint, channelId string) (bool, error) {
	pushInfos, err := FindPushInfos(PushInfo{ActivityId: activityId, ChannelId: channelId})
	return len(pushInfos) > 0, errors.WithStack(err)
}

func LoadPushInfosActivity(pushInfos ...*PushInfo) error {
	for _, p := range pushInfos {
		if err := p.LoadActivity(); err != nil {
			return err
		}
	}
	return nil
}

func FindPushInfoById(id uint) (*PushInfo, error) {
	var pushInfo PushInfo
	pushInfo.ID = id
	if err := GetDB().Where(&pushInfo).First(&pushInfo).Error; err != nil {
		return nil, err
	}
	if err := pushInfo.LoadActivity(); err != nil {
		return nil, err
	}
	return &pushInfo, nil
}

func FindPushInfo(cond PushInfo) (*PushInfo, error) {
	if err := GetDB().Where(&cond).First(&cond).Error; err != nil {
		return nil, err
	}
	if err := cond.LoadActivity(); err != nil {
		return nil, err
	}
	return &cond, nil
}

func FindPushInfos(cond PushInfo) ([]*PushInfo, error) {
	var result []*PushInfo
	if err := GetDB().Debug().Where(&cond).Find(&result).Error; err != nil {
		return nil, err
	}

	if err := LoadPushInfosActivity(result...); err != nil {
		return nil, err
	}

	return result, nil
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

// type UserServerQueryResult struct {
// 	Count int64        `json:"count"`
// 	Items []*BotServer `json:"items"`
// }

// type DiscordActivityQueryResult struct {
// 	Count int64                   `json:"count"`
// 	Items []*CustomActivityConfig `json:"items"`
// }

// type DoDoActivityQueryResult struct {
// 	Count int64                   `json:"count"`
// 	Items []*CustomActivityConfig `json:"items"`
// }

// type DiscordCustomProjectConfigQueryResult struct {
// 	Count int64        `json:"count"`
// 	Items []*BotServer `json:"items"`
// }

// type DoDoCustomProjectConfigQueryResult struct {
// 	Count int64        `json:"count"`
// 	Items []*BotServer `json:"items"`
// }

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

// func FindAndCountDiscordCustomProjectConfig(id uint, offset int, limit int) (*DiscordCustomProjectConfigQueryResult, error) {
// 	var items []*BotServer
// 	cond := &BotServer{}
// 	cond.RainbowUserId = id

// 	var count int64
// 	if err := db.Find(&items).Where(cond).Count(&count).Error; err != nil {
// 		return nil, err
// 	}

// 	if err := db.Find(&items).Where(cond).Offset(offset).Limit(limit).Error; err != nil {
// 		return nil, err
// 	}

// 	return &DiscordCustomProjectConfigQueryResult{count, items}, nil
// }

// func FindAndCountDoDoCustomProjectConfig(id uint, offset int, limit int) (*DoDoCustomProjectConfigQueryResult, error) {
// 	var items []*BotServer
// 	cond := &BotServer{}
// 	cond.RainbowUserId = id

// 	var count int64
// 	if err := db.Find(&items).Where(cond).Count(&count).Error; err != nil {
// 		return nil, err
// 	}

// 	if err := db.Find(&items).Where(cond).Offset(offset).Limit(limit).Error; err != nil {
// 		return nil, err
// 	}

// 	return &DoDoCustomProjectConfigQueryResult{count, items}, nil
// }
