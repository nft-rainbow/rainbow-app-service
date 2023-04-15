package services

import (
	"strings"

	"github.com/nft-rainbow/rainbow-app-service/models"
)

var (
	Success = map[string]string{"msg": "success"}
)

type VerifySocialServerReq struct {
	ServerId   string                `json:"server_id" binding:"required"`
	SocialTool models.SocialToolType `json:"social_tool" binding:"required"`
}
type InsertSocialServerReq struct {
	SocialTool models.SocialToolType `json:"social_tool" binding:"required"`
	ServerId   string                `json:"server_id" binding:"required"`
	AuthCode   string                `json:"auth_code" binding:"required"`
}

type PushInfoReq struct {
	ID         uint     `json:"id"`
	ChannelId  string   `json:"channel_id"`
	Roles      []string `json:"roles"`
	Content    string   `json:"content"`
	ColorTheme string   `json:"color_theme"`
	ActivityID uint     `json:"activity_id"`
}

func (p *PushInfoReq) ToModel(fillByRaw bool) (*models.PushInfo, error) {
	var result models.PushInfo
	if fillByRaw {
		raw, err := models.FindPushInfoById(p.ID)
		if err != nil {
			return nil, err
		}
		result = *raw
	}
	result.ChannelId = p.ChannelId
	result.Roles = strings.Join(p.Roles, ",")
	result.Content = p.Content
	result.ColorTheme = p.ColorTheme
	result.ActivityId = p.ActivityID
	result.ID = p.ID

	return &result, nil

}
