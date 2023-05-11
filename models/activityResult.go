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

type POAPResultFilter struct {
	Address string `form:"address" json:"address"`
}
type POAPResultQueryResult struct {
	Count int64         `json:"count"`
	Items []*POAPResult `json:"items"`
}

func FindAndCountPOAPResult(poapId string, filter POAPResultFilter, pagination Pagination) (*POAPResultQueryResult, error) {
	defaults.SetDefaults(&pagination)

	var items []*POAPResult
	cond := &POAPResult{}
	cond.ActivityCode = poapId
	cond.Address = filter.Address

	var count int64
	count, err := CountPOAPResult(poapId, &filter)
	if err != nil {
		return nil, err
	}

	if err := db.Model(&POAPResult{}).Where(cond).Order("id DESC").Offset(pagination.Offset()).Limit(pagination.Limit).Find(&items).Error; err != nil {
		return nil, err
	}

	return &POAPResultQueryResult{count, items}, nil
}

func CountPOAPResult(poapId string, filter *POAPResultFilter) (int64, error) {
	cond := &POAPResult{}
	cond.ActivityCode = poapId

	if filter == nil {
		return GetMintCountCache(poapId).GetCount(), nil
	}

	cond.Address = filter.Address

	var count int64
	err := db.Model(&POAPResult{}).Where(cond).Order("id DESC").Count(&count).Error
	return count, err
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

func GetMintSumByAddresses(poapId string, addresses ...string) (int64, error) {
	cond := &POAPResult{}
	cond.ActivityCode = poapId
	// cond.Address = address

	var count int64
	if err := db.Model(&POAPResult{}).Where(cond).Where("address in ?", addresses).Count(&count).Error; err != nil {
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

func (p *POAPResultCountCache) GetCount() int64 {
	p.RLock()
	defer p.RUnlock()
	return p.Count
}

func (p *POAPResultCountCache) SetCount(count int64) {
	p.Lock()
	defer p.Unlock()
	p.Count = count
}

func (p *POAPResultCountCache) Increase() {
	p.Lock()
	defer p.Unlock()
	p.Count += 1
}

func GetMintCountCache(ActivityID string) *POAPResultCountCache {
	countCache, ok := Cache[ActivityID]
	if !ok {
		countCache = &POAPResultCountCache{}
		Cache[ActivityID] = &POAPResultCountCache{}
	}

	if countCache.Count == 0 {
		var count int64
		if err := db.Model(&POAPResult{}).Where(&POAPResult{ActivityCode: ActivityID}).Count(&count).Error; err != nil {
			panic(err)
		}
		countCache.SetCount(count)
	}
	return countCache
}
