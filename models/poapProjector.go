package models

import (
	"sync"
)

type POAPActivityConfig struct {
	BaseModel
	ContractID         int32           `gorm:"type:integer" json:"contract_id" binding:"required"`
	Amount             int32           `gorm:"type:integer" json:"amount" binding:"required"`
	Name               string          `gorm:"type:string" json:"name" binding:"required"`
	Description        string          `gorm:"type:string" json:"description" binding:"required"`
	AppId              int32           `gorm:"index" json:"app_id" binding:"required"`
	ContractType       int32           `gorm:"type:int" json:"contract_type"`
	ContractAddress    string          `gorm:"type:string" json:"contract_address"`
	Chain              int32           `gorm:"type:int" json:"chain_type"`
	MetadataURI        string          `gorm:"type:string" json:"metadata_uri" binding:"required"`
	Command            string          `gorm:"type:string" json:"command"`
	EndedTime          int64           `gorm:"type:integer" json:"end_time" binding:"required"`
	StartedTime        int64           `gorm:"type:integer" json:"start_time" binding:"required"`
	RainbowUserId      int32           `gorm:"type:integer" json:"rainbow_user_id"`
	MaxMintCount       int32           `gorm:"type:varchar(256)" json:"max_mint_count" binding:"required"`
	ActivityID         string          `gorm:"type:string;index" json:"activity_id"`
	ActivityPictureURL string          `gorm:"type:string" json:"activity_picture_url"`
	SharingContent     string          `gorm:"type:string" json:"sharing_content"`
	WhiteListInfos     []WhiteListInfo `json:"white_list_infos"`
}

type WhiteListInfo struct {
	BaseModel
	User                 string `gorm:"type:string" json:"user"`
	Count                int32  `gorm:"type:integer" json:"count"`
	POAPActivityConfigID uint
}

type H5Config struct {
	BaseModel
	ActivityId       string `gorm:"type:integer;index" json:"activity_id"`
	Link             string `gorm:"type:string" json:"link" binding:"required"`
	Title            string `gorm:"type:string" json:"title"`
	TitleSize        int32  `gorm:"type:integer" json:"title_size"`
	TitleColor       int32  `gorm:"type:integer" json:"title_color"`
	Content          string `gorm:"type:string" json:"content"`
	ContentSize      int32  `gorm:"type:integer" json:"content_size"`
	ContentColor     int32  `gorm:"type:integer" json:"content_color"`
	ClaimButtonColor string `gorm:"type:string" json:"claim_button_color"`
	ButtonWordColor  int32  `gorm:"type:string" json:"button_word_color"`
	LogoURL          string `gorm:"type:string" json:"logo_url"`
	PCPicURL         string `gorm:"type:string" json:"pc_picture_url"`
	MobilePicURL     string `gorm:"type:string" json:"mobile_picture_url"`
}

type POAPResultCountCache struct {
	sync.RWMutex
	Count int64 `json:"count"`
}

type POAPActivityQueryResult struct {
	Count int64                 `json:"count"`
	Items []*POAPActivityConfig `json:"items"`
}

type POAPResultQueryResult struct {
	Count int64         `json:"count"`
	Items []*POAPResult `json:"items"`
}

var Cache = make(map[string]*POAPResultCountCache)

func FindPOAPActivityConfig(name string, contractId int32) (*POAPActivityConfig, error) {
	var item POAPActivityConfig
	err := db.Model(&POAPActivityConfig{}).Where("name = ?", name).Where("contract_id = ?", contractId).First(&item).Error
	return &item, err
}

func FindPOAPActivityConfigById(id string) (*POAPActivityConfig, error) {
	var item POAPActivityConfig
	var cond POAPActivityConfig

	cond.ActivityID = id
	err := db.Model(&POAPActivityConfig{}).Where(cond).Find(&item).Error
	if err != nil {
		return nil, err
	}
	err = db.Preload("WhiteListInfos").Find(&item).Error
	if err != nil {
		return nil, err
	}
	return &item, err
}

func FindAndCountPOAPActivity(id uint, offset int, limit int) (*POAPActivityQueryResult, error) {
	var items []*POAPActivityConfig
	cond := &POAPActivityConfig{}
	cond.RainbowUserId = int32(id)

	var count int64
	if err := db.Model(&POAPActivityConfig{}).Where(cond).Count(&count).Error; err != nil {
		return nil, err
	}

	if err := db.Model(&POAPActivityConfig{}).Where(cond).Offset(offset).Limit(limit).
		Find(&items).Error; err != nil {
		return nil, err
	}

	return &POAPActivityQueryResult{count, items}, nil
}

func FindAndCountPOAPResult(poapId string, offset int, limit int) (*POAPResultQueryResult, error) {
	var items []*POAPResult
	cond := &POAPResult{}
	cond.ActivityID = poapId

	var count int64
	count, err := CountPOAPResult(poapId)
	if err != nil {
		return nil, err
	}

	if err := db.Model(&POAPResult{}).Where(cond).Offset(offset).Limit(limit).Find(&items).Error; err != nil {
		return nil, err
	}

	return &POAPResultQueryResult{count, items}, nil
}

func CountPOAPResult(poapId string) (int64, error) {
	cond := &POAPResult{}
	cond.ActivityID = poapId

	var count int64

	countCache, ok := Cache[poapId]
	if !ok {
		countCache = &POAPResultCountCache{}
		Cache[poapId] = &POAPResultCountCache{}
	}
	countCache.RLock()

	if countCache.Count == 0 {
		countCache.RUnlock()
		countCache.Lock()
		if err := db.Model(&POAPResult{}).Where(cond).Count(&count).Error; err != nil {
			return 0, err
		}
		Cache[poapId].Count = count
		countCache.Unlock()
	} else {
		countCache.RUnlock()
		count = countCache.Count
	}

	return count, nil
}

func FindAndCountPOAPResultByAddress(offset int, limit int, address, poapId string) (*POAPResultQueryResult, error) {
	var items []*POAPResult
	cond := &POAPResult{}
	cond.ActivityID = poapId
	cond.Address = address

	var count int64
	if err := db.Model(&POAPResult{}).Where(cond).Count(&count).Error; err != nil {
		return nil, err
	}

	if err := db.Model(&POAPResult{}).Where(cond).Offset(offset).Limit(limit).Find(&items).Error; err != nil {
		return nil, err
	}

	return &POAPResultQueryResult{count, items}, nil
}

func CountPOAPResultByAddress(address, poapId string) (int64, error) {
	cond := &POAPResult{}
	cond.ActivityID = poapId
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
	cond.ActivityID = poapId
	cond.ID = uint(id)
	if err := db.Where(cond).Last(&resp).Error; err != nil {
		return nil, err
	}

	return resp, nil
}

func FindAndCountUnhandledPOAPResult(poapId string, offset, limit int, userAddress string) (*POAPResultQueryResult, error) {
	var items []*POAPResult
	cond := &POAPResult{}
	cond.ActivityID = poapId
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
