package models

import (
	"github.com/spf13/viper"
	"time"
)

type ShareMintCount struct {
	BaseModel
	Address    string `gorm:"type:string" json:"address"`
	Count      int32  `gorm:"type:integer" json:"count"`
	ActivityID string `gorm:"type:string;index" json:"activity_id"`
}

type ShareInfo struct {
	BaseModel
	Sharer     string `gorm:"type:string" json:"sharer"`
	Receiver   string `gorm:"type:string" json:"receiver"`
	ActivityID string `gorm:"type:string;index" json:"activity_id"`
}

type BatchBurnResult struct {
	BaseModel
	ActivityID string `gorm:"type:string;index" json:"activity_id"`
	Address    string `gorm:"type:string" json:"address"`
	Status     int32  `gorm:"type:integer" json:"status"`
	BurnID     int32  `gorm:"type:integer" json:"burn_id"`
	Hash       string `gorm:"type:string" json:"hash"`
}

type ClockTime struct {
	BaseModel
	Time       time.Time `json:"time"`
	ActivityID string    `gorm:"type:string;index" json:"activity_id"`
}

type EveryDayMintCount struct {
	BaseModel
	Address    string `gorm:"type:string" json:"address"`
	Count      int32  `gorm:"type:integer" json:"count"`
	ActivityID string `gorm:"type:string;index" json:"activity_id"`
}

func FindShareMintCount(address string, activityId string) (*ShareMintCount, error) {
	var cond ShareMintCount
	var item ShareMintCount
	cond.Address = address
	cond.ActivityID = activityId
	err := db.Where(&cond).First(&item).Error

	return &item, err
}

func FindEveryDayMintCount(address string, activityId string) (*EveryDayMintCount, error) {
	var cond EveryDayMintCount
	var item EveryDayMintCount
	cond.Address = address
	cond.ActivityID = activityId
	err := db.Where(&cond).First(&item).Error

	return &item, err
}

func UpdateMintCount(address, poapId string, updateCount int32) (*ShareMintCount, error) {
	item, err := FindShareMintCount(address, poapId)
	if err != nil {
		return nil, err
	}
	item.Count += updateCount

	db.Save(&item)

	return item, nil
}

func UpdateEveryDayMintCount(address, poapId string, updateCount int32) (*EveryDayMintCount, error) {
	item, err := FindEveryDayMintCount(address, poapId)
	if err != nil {
		return nil, err
	}
	item.Count += updateCount

	db.Save(&item)

	return item, nil
}

func FindSharingInfo(sharer, receiver, poapId string) (*ShareInfo, error) {
	var cond ShareInfo
	var item ShareInfo
	cond.Sharer = sharer
	cond.Receiver = receiver
	cond.ActivityID = poapId
	res := db.Where(&cond).First(&item)

	return &item, res.Error
}

func CountTodaySharerInfo(sharer, poapId string, now time.Time) (int64, error) {
	var cond ShareInfo
	cond.Sharer = sharer
	cond.ActivityID = poapId
	var count int64

	if viper.GetString("env") == "dev" {
		db.Model(&ShareInfo{}).Where(&cond).
			Where("updated_at > ? and updated_at < ?", now, now.Add(viper.GetDuration("testMinuteDuration")*time.Minute)).Count(&count)
	} else if viper.GetString("env") == "prod" {
		db.Model(&ShareInfo{}).Where(&cond).
			Where("updated_at > ? and updated_at < ?", now, now.Add(24*time.Hour)).Count(&count)
	}

	return count, nil
}

func CountSharerInfo(sharer, poapId string) (int64, error) {
	var cond ShareInfo
	cond.Sharer = sharer
	cond.ActivityID = poapId
	var count int64

	res := db.Find(&ShareInfo{}).Where(&cond).Count(&count)

	return count, res.Error
}

func GetClock(poapId string) (*ClockTime, error) {
	var item ClockTime
	var cond ClockTime
	cond.ActivityID = poapId

	res := db.Where("activity_id = ?", poapId).First(&item)
	return &item, res.Error
}

func CountUnhandledBurnResult(poapId string, userAddress string) (int64, error) {
	cond := &BatchBurnResult{}
	cond.ActivityID = poapId
	cond.Address = userAddress

	var count int64
	if err := db.Model(&BatchBurnResult{}).Where(cond).Not("status = ? or status = ?", 1, 2).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
