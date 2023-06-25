package models

import (
	"sync"

	"github.com/mcuadros/go-defaults"
	"github.com/nft-rainbow/rainbow-app-service/models/enums"
	"gorm.io/gorm"
)

var (
	Cache = make(map[string]*POAPResultCountCache)
)

type POAPResult struct {
	BaseModel
	Address       string                  `gorm:"type:string;index" json:"address" binding:"required"`
	ConfigID      int32                   `gorm:"type:integer" json:"config_id"`
	ContractRawID int32                   `gorm:"type:integer" json:"contract_id" binding:"required"`
	TxID          int32                   `gorm:"type:integer" json:"tx_id"`
	TokenID       string                  `gorm:"type:varchar(256)" json:"token_id"`
	Hash          string                  `gorm:"type:string" json:"hash"`
	ActivityCode  string                  `gorm:"type:string;index" json:"activity_id"` //TODO:  与前端一起更新为activity_code
	Status        enums.TransactionStatus `gorm:"type:integer;index" json:"status"`
	FileURL       string                  `gorm:"type:string" json:"file_url"`
	ProjectorId   uint                    `gorm:"type:integer" json:"projector_id"`
	AppId         uint                    `gorm:"type:integer" json:"app_id"`
	SocialId      string                  `gorm:"type:string;index" json:"social_id"`
	SocialType    uint                    `gorm:"type:integer" json:"social_type"`
}

type POAPResultFilter struct {
	Address  string                    `form:"address" json:"address"`
	Statuses []enums.TransactionStatus `form:"statuses" json:"statuses"`
}

func (p *POAPResultFilter) ToWhere() *gorm.DB {
	sql := db
	if p.Address != "" {
		sql = sql.Where("address = ?", p.Address)
	}
	if len(p.Statuses) > 0 {
		sql = sql.Where("status in (?)", p.Statuses)
	}
	return sql
}

type POAPResultQueryResult struct {
	Count int64         `json:"count"`
	Items []*POAPResult `json:"items"`
}

func FindAndCountPOAPResult(poapId string, filter POAPResultFilter, pagination Pagination) (*POAPResultQueryResult, error) {
	defaults.SetDefaults(&pagination)

	var items []*POAPResult
	var count int64

	if err := db.Model(&POAPResult{}).Where("activity_code=?", poapId).Where(filter.ToWhere()).Count(&count).Order("id DESC").Offset(pagination.Offset()).Limit(pagination.Limit).Find(&items).Error; err != nil {
		return nil, err
	}

	return &POAPResultQueryResult{count, items}, nil
}

func CountPOAPResult(poapId string, filter *POAPResultFilter) (int64, error) {
	if filter == nil {
		return GetMintCountCache(poapId).GetCount(), nil
	}

	var count int64
	err := db.Model(&POAPResult{}).Where("activity_code=?", poapId).Where(filter.ToWhere()).Order("id DESC").Count(&count).Error
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
