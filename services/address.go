package services

import (
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/utils"
)

func bindCFXAddress(req *models.BindCFXAddress) error{
	res := models.GetDB().Create(&req)
	if res.Error != nil {
		return  res.Error
	}
	return nil
}

func GetBindCFXAddress(userID string) (*models.BindCFXAddress, error) {
	resp, err := models.FindBindingCFXAddressById(userID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func HandleBindCfxAddress(userId, userAddress string) error{
	_, err := utils.CheckCfxAddress(utils.CONFLUX_TEST, userAddress)
	if err != nil {
		return err
	}
	dto := models.BindCFXAddress{
		DiscordId: userId,
		CFXAddress: userAddress,
	}

	err = bindCFXAddress(&dto)
	if err != nil {
		return err
	}

	return nil
}


