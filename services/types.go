package services

import (
	"strings"

	"github.com/nft-rainbow/rainbow-app-service/models"
)

type SocialToolQueryReq struct {
	SocialTool string `uri:"social_tool" form:"social_tool" binding:"required,oneof=dodo discord"`
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
