package models

import "github.com/nft-rainbow/rainbow-app-service/models/enums"

type SocialUserConfig struct {
	BaseModel
	UserId     string               `gorm:"type:varchar(256)" json:"user_id" binding:"required"`
	CFXAddress string               `gorm:"type:varchar(256)" json:"cfx_address" binding:"required"`
	SocialTool enums.SocialToolType `gorm:"type:integer" json:"social_tool" binding:"required"`
}

type CustomMintCount struct {
	BaseModel
	ChannelId string `gorm:"type:varchar(256)" json:"channel_id" binding:"required"`
	UserId    string `gorm:"type:varchar(256)" json:"user_id" binding:"required"`
	Count     uint   `gorm:"type:integer" json:"count" binding:"required"`
}

func FindSocialUserConfig(userSocialId string, socialTool enums.SocialToolType) (*SocialUserConfig, error) {
	result := SocialUserConfig{UserId: userSocialId, SocialTool: socialTool}
	if err := db.Where(&result).First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

// func CheckDiscordCustomCount(id, channelId string, maxCount int32) (bool, error) {
// 	config, err := FindDiscordCustomActivityConfigByChannelId(channelId)
// 	if err != nil {
// 		return false, err
// 	}
// 	if config.Amount == 0 {
// 		return false, errors.New("The number of the NFTs has reached the maximum in this channel")
// 	}
// 	var item CustomMintCount
// 	err = db.Where("user_id = ?", id).First(&item).Where("channel_id = ?", channelId).First(&item).Error
// 	if err != nil {
// 		err := InsertCustomCount(id, channelId)
// 		if err != nil {
// 			return false, err
// 		}
// 	}
// 	if item.Count == uint(maxCount) {
// 		return false, nil
// 	}
// 	return true, nil
// }

// func CheckDoDoCustomCount(id, channelId string, maxCount int32) (bool, error) {
// 	config, err := FindDoDoCustomActivityConfigByChannelId(channelId)
// 	if err != nil {
// 		return false, err
// 	}
// 	if config.Amount == 0 {
// 		return false, errors.New("The number of the NFTs has reached the maximum in this channel")
// 	}
// 	var item CustomMintCount
// 	err = db.Where("user_id = ?", id).First(&item).Where("channel_id = ?", channelId).First(&item).Error
// 	if err != nil {
// 		err := InsertCustomCount(id, channelId)
// 		if err != nil {
// 			return false, err
// 		}
// 	}
// 	if item.Count == uint(maxCount) {
// 		return false, nil
// 	}
// 	return true, nil
// }

// func UpdateDiscordCustomCount(id, channelId string) (*CustomMintCount, error) {
// 	var item CustomMintCount
// 	err := db.Where("user_id = ?", id).First(&item).Where("channel_id = ?", channelId).First(&item).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	db.Model(&item).Update("count", item.Count+1)

// 	var t CustomActivityConfig
// 	err = db.Where("channel_id = ?", channelId).First(&t).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	db.Model(&t).Update("amount", t.Amount-1)

// 	return &item, nil
// }

// func UpdateDoDoCustomCount(id, channelId string) (*CustomMintCount, error) {
// 	var item CustomMintCount
// 	err := db.Where("user_id = ?", id).First(&item).Where("channel_id = ?", channelId).First(&item).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	db.Model(&item).Update("count", item.Count+1)

// 	var t CustomActivityConfig
// 	err = db.Where("channel_id = ?", channelId).First(&t).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	db.Model(&t).Update("amount", t.Amount-1)

// 	return &item, nil
// }

func InsertCustomCount(id, channelId string) error {
	res := db.Create(&CustomMintCount{
		ChannelId: channelId,
		UserId:    id,
		Count:     0,
	})
	if res.Error != nil {
		return res.Error
	}
	return nil
}

//func CryptoDiscordId(id string) (string, error) {
//	hash := sha256.New()
//
//	_, err := hash.Write([]byte(id + "discord"))
//	if err != nil {
//		return "", err
//	}
//	sum := hash.Sum(nil)
//
//	newYearId := hex.EncodeToString(sum)
//	return newYearId[:8], nil
//}
//
//func CryptoDoDoId(id string) (string, error) {
//	hash := sha256.New()
//
//	_, err := hash.Write([]byte(id + "dodo"))
//	if err != nil {
//		return "", err
//	}
//	sum := hash.Sum(nil)
//
//	newYearId := hex.EncodeToString(sum)
//	return newYearId[:8], nil
//}
