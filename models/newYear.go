package models

import (
	"github.com/spf13/viper"
	"time"
)

type NewYearConfig struct {
	BaseModel
	Amount int32 `gorm:"type:integer" json:"amount" binding:"required"`
	Name string `gorm:"type:string" json:"name" binding:"required"`
	Description string `gorm:"type:string" json:"description" binding:"required"`
	AppId int32 `gorm:"index" json:"app_id" binding:"required"`
	Chain    int32   `gorm:"type:int" json:"chain_type"`
	EndedTime int64 `gorm:"type:integer" json:"end_time" binding:"required"`
	StartedTime int64 `gorm:"type:integer" json:"start_time" binding:"required"`
	RainbowUserId int32 `gorm:"type:integer" json:"rainbow_user_id"`
	ContractType int32 `gorm:"type:integer" json:"contract_type"`
	ContractAddress string `gorm:"type:string" json:"contract_address"`
	ContractID int32 `gorm:"type:integer" json:"contract_id" binding:"required"`
	MaxMintCount int32 `gorm:"type:varchar(256)" json:"max_mint_count" binding:"required"`
	ActivityPictureURL string `gorm:"type:string" json:"activity_picture_url"`
	SharingContent string `gorm:"type:string" json:"sharing_content"`
	ActivityID string `gorm:"type:string" json:"activity_id"`
	ContractInfos []NFTContractInfo `json:"nft_contract_infos"`
}

type ShareMintCount struct {
	BaseModel
	Address string `gorm:"type:string" json:"address"`
	Count int32 `gorm:"type:integer" json:"count"`
	ActivityID string `gorm:"type:string" json:"activity_id"`
}

type NFTContractInfo struct {
	BaseModel
	MetadataURI string `gorm:"type:string" json:"metadata_uri" binding:"required"`
	Probability float32 `gorm:"type:varchar(256)" json:"probability" binding:"required"`
	TokenID    string `gorm:"type:string" json:"token_id" binding:"required"`
	NewYearConfigID uint
}

type ShareInfo struct {
	BaseModel
	Sharer string `gorm:"type:string" json:"sharer"`
	Receiver string `gorm:"type:string" json:"receiver"`
	ActivityID string `gorm:"type:string" json:"activity_id"`
}

type ClockTime struct {
	BaseModel
	Time time.Time `json:"time"`
	ActivityID string `gorm:"type:string" json:"activity_id"`
}

type NewYearConfigQueryResult struct {
	Count int64       `json:"count"`
	Items []*NewYearConfig `json:"items"`
}

type NewYearSpecialConfigQueryResult struct {
	Count int64       `json:"count"`
	Items []*NewYearConfig `json:"items"`
}

type EveryDayMintCount struct {
	BaseModel
	Address string `gorm:"type:string" json:"address"`
	Count int32 `gorm:"type:integer" json:"count"`
	ActivityID string `gorm:"type:string" json:"activity_id"`
}

func FindNewYearConfigById(id string) (*NewYearConfig, error){
	var item NewYearConfig
	err := db.Where("activity_id = ?", id).First(&item).Error
	if err != nil {
		return nil, err
	}

	err = db.Preload("ContractInfos").Find(&item).Error
	if err != nil {
		return nil, err
	}
	return &item, err
}

func FindShareMintCount(address string, activityId string) (*ShareMintCount, error){
	var cond ShareMintCount
	var item ShareMintCount
	cond.Address = address
	cond.ActivityID = activityId
	err := db.Where(&cond).Last(&item).Error

	return &item, err
}

func FindEveryDayMintCount(address string, activityId string)(*EveryDayMintCount, error) {
	var cond EveryDayMintCount
	var item EveryDayMintCount
	cond.Address = address
	cond.ActivityID = activityId
	err := db.Where(&cond).Last(&item).Error

	return &item, err
}

func UpdateMintCount(address, poapId string, updateCount int32) (*ShareMintCount, error){
	item, err := FindShareMintCount(address, poapId)
	if err != nil {
		return nil, err
	}
	item.Count += updateCount

	db.Save(&item)

	return item, nil
}

func UpdateEveryDayMintCount(address, poapId string, updateCount int32) (*EveryDayMintCount, error){
	item, err := FindEveryDayMintCount(address, poapId)
	if err != nil {
		return nil, err
	}
	item.Count += updateCount

	db.Save(&item)

	return item, nil
}

func FindSharingInfo(sharer, receiver, poapId string) (*ShareInfo, error){
	var cond ShareInfo
	var item ShareInfo
	cond.Sharer = sharer
	cond.Receiver = receiver
	cond.ActivityID = poapId
	res := db.Where(&cond).First(&item)

	return &item, res.Error
}

func CountTodaySharerInfo(sharer, poapId string, now time.Time)(int64, error) {
	var item []*ShareInfo
	var cond ShareInfo
	cond.Sharer = sharer
	cond.ActivityID = poapId
	var count int64
	if viper.GetString("env") == "dev" {
		db.Find(&item).Where(&cond).
			Where("updated_at > ? and updated_at < ?", now, now.Add(viper.GetDuration("testMinuteDuration") * time.Minute)).Count(&count)
	}else if viper.GetString("env") == "prod"{
		db.Find(&item).Where(&cond).
			Where("updated_at > ? and updated_at < ?", now, now.Add(24 * time.Hour)).Count(&count)
	}

	return count, nil
}

func CountSharerInfo(sharer, poapId string)(int64, error) {
	var cond ShareInfo
	cond.Sharer = sharer
	cond.ActivityID = poapId
	var count int64

	res := db.Find(&ShareInfo{}).Where(&cond).Count(&count)

	return count, res.Error
}

func GetClock(poapId string) (*ClockTime, error){
	var item ClockTime
	var cond ClockTime
	cond.ActivityID = poapId

	res := db.Where("activity_id = ?", poapId).First(&item)
	return &item, res.Error
}