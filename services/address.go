package services

import (
	"github.com/nft-rainbow/discordbot-service/models"
)

func BindCFXAddress(req *models.BindCFXAddress) error{
	res := models.GetDB().Create(&req)
	if res.Error != nil {
		return  res.Error
	}
	return nil
}

func GetBindCFXAddress(userID string) (*models.BindCFXAddress, error) {
	resp, err := models.FindBindingAddressById(userID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

