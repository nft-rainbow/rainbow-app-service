package models

import "fmt"

type BindCFXAddress struct {
	BaseModel
	UserId string `gorm:"type:varchar(256)" json:"user_id" binding:"required"`
	UserAddress string `gorm:"type:varchar(256)" json:"user_address" binding:"required"`
}

type GetBindCFXAddressResp struct{
	CFXAddress string `json:"cfx_address"`
	UserId string `json:"user_id"`
}

type CustomMintCount struct {
	BaseModel
	ChannelId string `gorm:"type:varchar(256)" json:"channel_id" binding:"required"`
	UserId string `gorm:"type:varchar(256)" json:"user_id" binding:"required"`
	Count uint `gorm:"type:integer" json:"count" binding:"required"`
}

type EasyMintCount struct {
	BaseModel
	ChannelId string `gorm:"type:varchar(256)" json:"channel_id" binding:"required"`
	UserId string `gorm:"type:varchar(256)" json:"user_id" binding:"required"`
	Count uint `gorm:"type:integer" json:"count" binding:"required"`
}

func FindBindingAddressById(id string) (*BindCFXAddress, error) {
	var item BindCFXAddress
	err := db.Where("user_id = ?", id).First(&item).Error
	return &item, err
}

func CheckCustomCount(id, channelId string, maxCount uint) (bool, error){
	var item CustomMintCount
	err := db.Where("user_id = ?", id).First(&item).Where("channel_id = ?", channelId).First(&item).Error
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

func CheckEasyCount(id, channelId string, maxCount uint) (bool, error){
	var item EasyMintCount
	err := db.Where("user_id = ?", id).First(&item).Where("channel_id = ?", channelId).First(&item).Error
	if err != nil {
		return false, err
	}
	if item.Count == maxCount {
		return false, nil
	}
	return true, nil
}

func UpdateCustomCount(id, channelId string) (*CustomMintCount, error){
	var item CustomMintCount
	err := db.Where("user_id = ?", id).First(&item).Where("channel_id = ?", channelId).First(&item).Error
	if err != nil {
		return nil, err
	}
	fmt.Println(item)
	//item.Count += 1
	//db.Save(&item)
	db.Model(&item).Update("count", item.Count+1)

	return &item, nil
}

func UpdateEasyCount(id, channelId string)(*EasyMintCount, error) {
	var item EasyMintCount
	err := db.Where("user_id = ?", id).First(&item).Where("channel_id = ?", channelId).First(&item).Error
	if err != nil {
		return nil, err
	}
	item.Count += 1
	db.Save(&item)

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



