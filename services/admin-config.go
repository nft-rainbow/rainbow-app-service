package services

import "github.com/nft-rainbow/discordbot-service/models"

func BindAdminConfig(config *models.AdminConfig) error{
	res := models.GetDB().Create(&config)
	if res.Error != nil {
		return  res.Error
	}
	return nil
}
