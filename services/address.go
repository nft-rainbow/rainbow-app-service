package services

import (
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/utils"
)

func bindCFXAddressWithDiscord(req *models.BindCFXWithDiscord) error{
	res := models.GetDB().Create(&req)
	if res.Error != nil {
		return  res.Error
	}
	return nil
}

func bindCFXAddressWithDoDo(req *models.BindCFXWithDoDo) error{
	res := models.GetDB().Create(&req)
	if res.Error != nil {
		return  res.Error
	}
	return nil
}

func GetDoDoBindCFXAddress(userID string) (string, error) {
	resp, err := models.FindDoDoBindingCFXAddressById(userID)
	if err != nil {
		return "", err
	}
	return resp.CFXAddress, nil
}

func GetDiscordBindCFXAddress(userID string) (string, error) {
	resp, err := models.FindDiscordBindingCFXAddressById(userID)
	if err != nil {
		return "", err
	}
	return resp.CFXAddress, nil
}

func HandleBindCfxAddress(userId, userAddress, platform string) error{
	var err error
	_, err = utils.CheckCfxAddress(utils.CONFLUX_TEST, userAddress)
	if err != nil {
		return err
	}

	if platform == "discord" {
		err = bindCFXAddressWithDiscord(&models.BindCFXWithDiscord{
			DiscordId: userId,
			CFXAddress: userAddress,
		})
	}else if platform == "dodo"{
		err = bindCFXAddressWithDoDo(&models.BindCFXWithDoDo{
			DoDoId: userId,
			CFXAddress: userAddress,
		})
	}
	if err != nil {
		return err
	}

	return nil
}


