package services

import "github.com/nft-rainbow/rainbow-app-service/models"

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
	ChannelId  string `gorm:"type:string" json:"channel_id"`
	Roles      string `gorm:"type:string" json:"roles"`
	Content    string `gorm:"type:string" json:"content"`
	ColorTheme string `gorm:"type:string" json:"color_theme"`
	ActivityID uint   `json:"activity_id"`
}

func (p *PushInfoReq) ToModel() (*models.PushInfo, error) {
	var activity models.Activity
	if err := models.GetDB().Model(&models.Activity{}).Where("id=?", p.ActivityID).First(&activity).Error; err != nil {
		return nil, err
	}
	result := models.PushInfo{
		ChannelId:  p.ChannelId,
		Roles:      p.Roles,
		Content:    p.Content,
		ColorTheme: p.ColorTheme,
		Activity:   activity,
	}
	return &result, nil

}
