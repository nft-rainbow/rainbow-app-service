package models

import (
	"sync"

	"github.com/mcuadros/go-defaults"
)

var (
	Cache = make(map[string]*POAPResultCountCache)
)

type POAPResult struct {
	BaseModel
	Address       string `gorm:"type:string;index" json:"address" binding:"required"`
	ConfigID      int32  `gorm:"type:integer" json:"config_id"`
	ContractRawID int32  `gorm:"type:integer" json:"contract_id" binding:"required"`
	TxID          int32  `gorm:"type:integer" json:"tx_id"`
	TokenID       string `gorm:"type:varchar(256)" json:"token_id"`
	Hash          string `gorm:"type:string" json:"hash"`
	ActivityCode  string `gorm:"type:string;index" json:"activity_id"` //TODO:  与前端一起更新为activity_code
	Status        int32  `gorm:"type:integer;index" json:"status"`
	FileURL       string `gorm:"type:string" json:"file_url"`
	ProjectorId   uint   `gorm:"type:integer" json:"projector_id"`
	AppId         uint   `gorm:"type:integer" json:"app_id"`
	SocialId      string `gorm:"type:string;index" json:"social_id"`
	SocialType    uint   `gorm:"type:integer" json:"social_type"`
}

type POAPResultQueryResult struct {
	Count int64         `json:"count"`
	Items []*POAPResult `json:"items"`
}

func FindAndCountPOAPResult(poapId string, pagination Pagination) (*POAPResultQueryResult, error) {
	defaults.SetDefaults(&pagination)

	var items []*POAPResult
	cond := &POAPResult{}
	cond.ActivityCode = poapId

	var count int64
	count, err := CountPOAPResult(poapId)
	if err != nil {
		return nil, err
	}

	if err := db.Model(&POAPResult{}).Where(cond).Offset(pagination.Offset()).Limit(pagination.Limit).Find(&items).Error; err != nil {
		return nil, err
	}

	return &POAPResultQueryResult{count, items}, nil
}

func CountPOAPResult(poapId string) (int64, error) {
	cond := &POAPResult{}
	cond.ActivityCode = poapId

	cache, err := InitCache(poapId)
	if err != nil {
		return 0, err
	}

	return cache.Count, nil
}

func CountPOAPResultBySocial(socialId, poapId string, socialType uint) (int64, error) {
	cond := &POAPResult{
		ActivityCode: poapId,
		SocialId:     socialId,
		SocialType:   socialType,
	}
	var count int64
	err := db.Model(&POAPResult{}).Where(cond).Count(&count).Error
	return count, err
}

func CountPOAPResultByAddress(address, poapId string) (int64, error) {
	cond := &POAPResult{}
	cond.ActivityCode = poapId
	cond.Address = address

	var count int64
	if err := db.Model(&POAPResult{}).Where(cond).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func FindPOAPResultById(poapId string, id int) (*POAPResult, error) {
	cond := &POAPResult{}
	resp := &POAPResult{}
	cond.ActivityCode = poapId
	cond.ID = uint(id)
	if err := db.Where(cond).Last(&resp).Error; err != nil {
		return nil, err
	}

	return resp, nil
}

func FindAndCountUnhandledPOAPResult(poapId string, offset, limit int, userAddress string) (*POAPResultQueryResult, error) {
	var items []*POAPResult
	cond := &POAPResult{}
	cond.ActivityCode = poapId
	cond.Address = userAddress

	var count int64
	if err := db.Model(&POAPResult{}).Where(cond).Count(&count).Error; err != nil {
		return nil, err
	}

	if err := db.Model(&POAPResult{}).Where(cond).Offset(offset).Limit(limit).Find(&items).Error; err != nil {
		return nil, err
	}
	return &POAPResultQueryResult{count, items}, nil
}

type POAPResultCountCache struct {
	sync.RWMutex
	Count int64 `json:"count"`
}

func InitCache(ActivityID string) (*POAPResultCountCache, error) {
	var count int64
	countCache, ok := Cache[ActivityID]
	if !ok {
		countCache = &POAPResultCountCache{}
		Cache[ActivityID] = &POAPResultCountCache{}
	}

	if countCache.Count == 0 {
		if err := db.Model(&POAPResult{}).Where(&POAPResult{ActivityCode: ActivityID}).Count(&count).Error; err != nil {
			return nil, err
		}
		countCache.Lock()
		Cache[ActivityID].Count = count
		countCache.Unlock()
	} else {
		countCache.RLock()
		count = countCache.Count
		countCache.RUnlock()
	}
	return countCache, nil
}
