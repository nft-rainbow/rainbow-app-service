package models

import "errors"

type BindCFXWithDiscord struct {
	BaseModel
	DiscordId string `gorm:"type:varchar(256)" json:"discord_id" binding:"required"`
	CFXAddress string `gorm:"type:varchar(256)" json:"cfx_address" binding:"required"`
}

type BindCFXWithDoDo struct {
	BaseModel
	DoDoId string `gorm:"type:varchar(256)" json:"do_do_id" binding:"required"`
	CFXAddress string `gorm:"type:varchar(256)" json:"cfx_address" binding:"required"`
}

type CustomMintCount struct {
	BaseModel
	ChannelId string `gorm:"type:varchar(256)" json:"channel_id" binding:"required"`
	UserId string `gorm:"type:varchar(256)" json:"user_id" binding:"required"`
	Count uint `gorm:"type:integer" json:"count" binding:"required"`
}

func FindDiscordBindingCFXAddressById(id string) (*BindCFXWithDiscord, error) {
	var item BindCFXWithDiscord
	err := db.Where("discord_id = ?", id).First(&item).Error
	return &item, err
}

func FindDoDoBindingCFXAddressById(id string) (*BindCFXWithDoDo, error) {
	var item BindCFXWithDoDo
	err := db.Where("do_do_id = ?", id).First(&item).Error
	return &item, err
}

func CheckDiscordCustomCount(id, channelId string, maxCount uint) (bool, error){
	config, err := FindBindingDiscordActivityConfigByChannelId(channelId)
	if err != nil {
		return false, err
	}
	if config.Amount == 0 {
		return false, errors.New("The number of the NFTs has reached the maximum in this channel")
	}
	var item CustomMintCount
	err = db.Where("user_id = ?", id).First(&item).Where("channel_id = ?", channelId).First(&item).Error
	if err != nil {
		err := InsertCustomCount(id,channelId)
		if err != nil {
			return false, err
		}
	}
	if item.Count == maxCount {
		return false, nil
	}
	return true, nil
}

func CheckDoDoCustomCount(id, channelId string, maxCount uint) (bool, error){
	config, err := FindBindingDoDoActivityConfigByChannelId(channelId)
	if err != nil {
		return false, err
	}
	if config.Amount == 0 {
		return false, errors.New("The number of the NFTs has reached the maximum in this channel")
	}
	var item CustomMintCount
	err = db.Where("user_id = ?", id).First(&item).Where("channel_id = ?", channelId).First(&item).Error
	if err != nil {
		err := InsertCustomCount(id,channelId)
		if err != nil {
			return false, err
		}
	}
	if item.Count == maxCount {
		return false, nil
	}
	return true, nil
}

func UpdateDiscordCustomCount(id, channelId string) (*CustomMintCount, error){
	var item CustomMintCount
	err := db.Where("user_id = ?", id).First(&item).Where("channel_id = ?", channelId).First(&item).Error
	if err != nil {
		return nil, err
	}
	db.Model(&item).Update("count", item.Count+1)

	var t DiscordActivityConfig
	err = db.Where("channel_id = ?", channelId).First(&t).Error
	if err != nil {
		return nil, err
	}
	db.Model(&t).Update("amount", t.Amount - 1)

	return &item, nil
}

func UpdateDoDoCustomCount(id, channelId string) (*CustomMintCount, error){
	var item CustomMintCount
	err := db.Where("user_id = ?", id).First(&item).Where("channel_id = ?", channelId).First(&item).Error
	if err != nil {
		return nil, err
	}
	db.Model(&item).Update("count", item.Count+1)

	var t DoDoActivityConfig
	err = db.Where("channel_id = ?", channelId).First(&t).Error
	if err != nil {
		return nil, err
	}
	db.Model(&t).Update("amount", t.Amount - 1)

	return &item, nil
}

func InsertCustomCount(id, channelId string) error{
	res := db.Create(&CustomMintCount{
		ChannelId: channelId,
		UserId: id,
		Count: 0,
	})
	if res.Error != nil {
		return  res.Error
	}
	return nil
}



