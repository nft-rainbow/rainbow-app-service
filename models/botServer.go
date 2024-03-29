package models

import (
	"fmt"

	"github.com/mcuadros/go-defaults"
	"github.com/nft-rainbow/rainbow-app-service/models/enums"
)

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
	RainbowUserId            uint                 `gorm:"type:integer" json:"rainbow_user_id" binding:"required"`
	SocialTool               enums.SocialToolType `json:"social_tool"`
	RawServerId              string               `gorm:"type:varchar(255);index" json:"raw_server_id" binding:"required"`
	OutdatedServerId         string               `gorm:"type:varchar(255)" json:"outdated_server_id"`
	ServerName               string               `gorm:"type:varchar(255)" json:"server_name"`
	OwnerSocialId            string               `gorm:"type:varchar(255)" json:"owner_social_id" binding:"required"`
	PushInfos                []*PushInfo          `gorm:"-" json:"push_infos"`
	DefaultActivityChannelId string               `gorm:"type:varchar(255)" json:"default_activity_channel_id"`
}

type (
	FindBotServerActivitiesCond struct {
		Pagination
		SocialTool      string  `form:"social_tool" binding:"required,oneof=dodo discord" swaggerignore:"true"`
		ActivityName    *string `form:"activity_name" swaggerignore:"true"`
		ContractAddress *string `form:"contract_address" swaggerignore:"true"`
	}

	FlattenBotServerActivity struct {
		RainbowUserId   uint                 `json:"rainbow_user_id"`
		SocialTool      enums.SocialToolType `json:"social_tool"`
		RawServerId     string               `json:"raw_server_id"`
		ServerName      string               `json:"server_name"`
		OwnerSocialId   string               `json:"owner_social_id"`
		ActivityCode    string               `json:"activity_id"`
		ActivityType    enums.ActivityType   `json:"activity_type"`
		Name            string               `json:"name" binding:"required"`
		EndedTime       int64                `json:"end_time" default:"-1"`
		StartedTime     int64                `json:"start_time" default:"-1"`
		ContractRawID   *int32               `json:"contract_id"`
		ContractAddress string               `json:"contract_address"`
		ChainId         uint                 `json:"chain_id"`
		ChainType       int32                `json:"chain_type"`
		PushInfoId      uint                 `json:"push_info_id"`
		ChannelId       string               `json:"channel_id"`
		LastPushTime    int64                `json:"last_push_time"`
	}

	FindBotServerActivitiesResult struct {
		Count int64                       `json:"count"`
		Items []*FlattenBotServerActivity `json:"items"`
	}

	FindBotServersResult struct {
		Count int64        `json:"count"`
		Items []*BotServer `json:"items"`
	}
)

func (b *BotServer) LoadPushInfos() error {
	// return GetDB().Model(&PushInfo{}).Where("bot_server_id=?", b.ID).Find(&b.PushInfos).Error
	pushInfos, err := FindPushInfos(PushInfo{BotServerID: b.ID})
	if err != nil {
		return err
	}
	b.PushInfos = pushInfos
	return nil
}

func CompleteBotServers(bs ...*BotServer) error {
	for _, b := range bs {
		if err := b.LoadPushInfos(); err != nil {
			return err
		}
	}
	return nil
}

func DoAndCompleteBotServers(f func() ([]*BotServer, error)) ([]*BotServer, error) {
	botServers, err := f()
	if err != nil {
		return nil, err
	}
	if err := CompleteBotServers(botServers...); err != nil {
		return nil, err
	}
	return botServers, nil
}

func DoAndCompleteBotServer(f func() (*BotServer, error)) (*BotServer, error) {
	botServer, err := f()
	if err != nil {
		return nil, err
	}
	if err := CompleteBotServers(botServer); err != nil {
		return nil, err
	}
	return botServer, nil
}

func FindBotServers(rainbowUserId uint, socialTool *enums.SocialToolType, pagination Pagination) (*FindBotServersResult, error) {
	defaults.SetDefaults(&pagination)

	cond := BotServer{RainbowUserId: rainbowUserId}
	if socialTool != nil {
		cond.SocialTool = *socialTool
	}

	botServers, err := DoAndCompleteBotServers(func() ([]*BotServer, error) {
		var result []*BotServer
		err := GetDB().Where(&cond).Offset(pagination.Offset()).Limit(pagination.Limit).Find(&result).Error
		if err != nil {
			return nil, err
		}
		return result, nil
	})
	if err != nil {
		return nil, err
	}

	var count int64
	if err := GetDB().Model(&cond).Where(&cond).Count(&count).Error; err != nil {
		return nil, err
	}

	return &FindBotServersResult{
		Count: count,
		Items: botServers,
	}, nil

}

func FindActivitiesOfUserBotServers(rainbowUserId uint, cond *FindBotServerActivitiesCond) (*FindBotServerActivitiesResult, error) {
	socialTool, err := enums.ParseSocialToolType(cond.SocialTool)
	if err != nil {
		return nil, err
	}

	filters := fmt.Sprintf("b.rainbow_user_id=%v and b.social_tool=%v and c.contract_address!=\"\"", rainbowUserId, uint(*socialTool))
	if cond.ActivityName != nil {
		filters += fmt.Sprintf(" and a.name like \"%%%s%%\"", *cond.ActivityName)
	}
	if cond.ContractAddress != nil {
		filters += fmt.Sprintf(" and c.contract_address=\"%s\"", *cond.ContractAddress)
	}

	fields := "b.rainbow_user_id,b.social_tool,b.raw_server_id,b.server_name,b.owner_social_id," +
		"a.activity_code,a.activity_type," +
		"a.name,a.ended_time,a.started_time," +
		"c.contract_raw_id,c.contract_address,c.chain_id,c.chain_type," +
		"p.id as push_info_id,p.channel_id,p.last_push_time"

	joins := "bot_servers as b " +
		"left join push_infos as p on b.id=p.bot_server_id " +
		"left join activities as a on p.activity_id=a.id " +
		"left join contracts as c on a.contract_raw_id=c.contract_raw_id"

	itemsSql := fmt.Sprintf("select %s from  %s  where %s order by p.id desc limit %v,%v", fields, joins, filters, cond.Offset(), cond.Limit)
	countSql := fmt.Sprintf("select %s from  %s  where %s order by p.id desc", "count(*)", joins, filters)

	var result FindBotServerActivitiesResult

	if err := db.Debug().Raw(itemsSql).Scan(&result.Items).Error; err != nil {
		return nil, err
	}

	if err := db.Raw(countSql).Scan(&result.Count).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func FindBotServerByChannel(channelId string) (*BotServer, error) {
	// select * from bot_servers left join push_infos on bot_servers.id=push_infos.bot_server_id where push_infos.channel_id=1565461 or bot_servers.default_activity_channel_id=1565461
	bs, err := DoAndCompleteBotServer(func() (*BotServer, error) {
		var result *BotServer
		err := db.Model(&BotServer{}).
			Joins("left join push_infos on bot_servers.id=push_infos.bot_server_id").
			Where("push_infos.channel_id=?", channelId).
			Or("bot_servers.default_activity_channel_id=?", channelId).
			First(&result).Error
		return result, err
	})
	return bs, err
}

func FindBotServerByRawID(rawServerId string, socialTool *enums.SocialToolType) (*BotServer, error) {
	bs, err := DoAndCompleteBotServer(func() (*BotServer, error) {
		cond := BotServer{
			RawServerId: rawServerId,
		}
		if socialTool != nil {
			cond.SocialTool = *socialTool
		}
		if err := db.Where(&cond).First(&cond).Error; err != nil {
			return nil, err
		}
		return &cond, nil
	})
	return bs, err
}

func FindBotServerById(id uint) (*BotServer, error) {
	bs, err := DoAndCompleteBotServer(func() (*BotServer, error) {
		var result *BotServer
		err := db.Where("id=?", id).First(&result).Error
		return result, err
	})
	return bs, err
}

func FirstBotServerByUserId(rainbowUserId int) (*BotServer, error) {
	bs, err := DoAndCompleteBotServer(func() (*BotServer, error) {
		var item BotServer
		err := db.Where("rainbow_user_id = ?", rainbowUserId).First(&item).Error
		return &item, err
	})
	return bs, err
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
