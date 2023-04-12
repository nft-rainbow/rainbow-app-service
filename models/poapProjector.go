package models

import (
	"errors"
	"sync"
)

type Activity struct {
	BaseModel
	ContractID             *int32          `gorm:"type:integer" json:"contract_id"`
	Amount                 int32           `gorm:"type:integer" json:"amount" binding:"required"`
	Name                   string          `gorm:"type:string" json:"name" binding:"required"`
	Description            string          `gorm:"type:string" json:"description" binding:"required"`
	AppId                  uint            `gorm:"index" json:"app_id" binding:"required"`
	AppName                string          `gorm:"string" json:"app_name" binding:"required"`
	ContractType           int32           `gorm:"type:int" json:"contract_type"`
	ContractAddress        *string         `gorm:"type:string" json:"contract_address"`
	ChainId                int32           `gorm:"type:int" json:"chain_id"`
	ChainType              int32           `gorm:"type:int" json:"chain_type"`
	ActivityType           uint            `gorm:"type:uint" json:"activity_type" binding:"required"`
	Command                string          `gorm:"type:string" json:"command,omitempty"`
	IsCommand              bool            `gorm:"type:bool" json:"is_command"`
	IsPhoneWhiteListOpened bool            `gorm:"type:bool;default:false" json:"is_phone_white_list_opened"`
	EndedTime              int64           `gorm:"type:integer" json:"end_time"`
	StartedTime            int64           `gorm:"type:integer" json:"start_time"`
	RainbowUserId          uint            `gorm:"type:integer" json:"rainbow_user_id"`
	MaxMintCount           int32           `gorm:"type:varchar(256)" json:"max_mint_count" binding:"required"`
	ActivityID             string          `gorm:"type:string;index" json:"activity_id"`
	ActivityPictureURL     string          `gorm:"type:string" json:"activity_picture_url"`
	ActivityPosterURL      string          `gorm:"type:string" json:"activity_poster_url"`
	WhiteListInfos         []WhiteListInfo `json:"white_list_infos"`
	NFTConfigs             []NFTConfig     `json:"nft_configs"`
	MetadataUri            string          `gorm:"type:string" json:"metadata_uri"`
	PushInfoID             *uint
}

func (p *Activity) CheckContractValid() error {
	if p.ContractID == nil {
		return errors.New("contract id is empty")
	}
	if p.ContractAddress == nil {
		return errors.New("contract address is empty")
	}
	return nil
}

func (p *Activity) CheckActivityValid() error {
	if p.ActivityID == "" {
		return errors.New("activity id is empty")
	}
	return nil
}

func (p *Activity) CheckContractAndActivityValid() error {
	if err := p.CheckContractValid(); err != nil {
		return err
	}
	return p.CheckActivityValid()
}

type NFTConfig struct {
	BaseModel
	ImageURL           string               `gorm:"type:string" json:"image_url"`
	Name               string               `gorm:"type:string" json:"name"`
	Probability        float32              `gorm:"type:float" json:"probability"`
	MetadataAttributes []*MetadataAttribute `json:"metadata_attributes"`
	ActivityID         uint
}

type Metadata struct {
	Name         string `gorm:"type:string" json:"name"`
	Description  string `gorm:"type:string" json:"description"`
	ExternalLink string `gorm:"type:string" json:"external_link"`
}

type MetadataAttribute struct {
	BaseModel
	TraitType   string `gorm:"type:varchar(256)"  json:"trait_type"`
	DisplayType string `gorm:"type:varchar(256)"  json:"display_type,omitempty"`
	Value       string `gorm:"type:varchar(256)"  json:"value"`
	NFTConfigID uint
}

type WhiteListInfo struct {
	BaseModel
	User       string `gorm:"type:string" json:"user"`
	Count      int32  `gorm:"type:integer" json:"count"`
	ActivityID uint
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
	Count int64       `json:"count"`
	Items []*Activity `json:"items"`
}

type POAPResultQueryResult struct {
	Count int64         `json:"count"`
	Items []*POAPResult `json:"items"`
}

type POAPActivityFindCondition struct {
	Name     string `form:"name"`
	Activity string `form:"activity_id"`
	Contract string `form:"contract_address"`
}

var Cache = make(map[string]*POAPResultCountCache)

func FindPOAPActivityConfig(name string, contractId int32) (*Activity, error) {
	var item Activity
	err := db.Model(&Activity{}).Where("name = ?", name).Where("contract_id = ?", contractId).First(&item).Error
	return &item, err
}

func FindPOAPActivityConfigById(id string) (*Activity, error) {
	var item Activity
	var cond Activity

	cond.ActivityID = id
	err := db.Model(&Activity{}).Where(cond).Find(&item).Error
	if err != nil {
		return nil, err
	}
	err = db.Preload("WhiteListInfos").Preload("NFTConfigs").Preload("NFTConfigs.MetadataAttributes").Find(&item).Error
	if err != nil {
		return nil, err
	}

	return &item, err
}

func FindAndCountPOAPActivity(ranbowUserId uint, offset int, limit int, _cond POAPActivityFindCondition) (*POAPActivityQueryResult, error) {
	var items []*Activity
	cond := &Activity{}
	cond.RainbowUserId = ranbowUserId
	cond.Name = _cond.Name
	cond.ActivityID = _cond.Activity

	if _cond.Contract != "" {
		cond.ContractAddress = &_cond.Contract
	}

	var count int64
	if err := db.Debug().Model(&Activity{}).Preload("WhiteListInfos").Preload("NFTConfigs").Preload("NFTConfigs.MetadataAttributes").Where(cond). /*.Where("activity_id=? and contract_address=?", activity, contract)*/ Count(&count).Error; err != nil {
		return nil, err
	}

	if err := db.Debug().Model(&Activity{}).Preload("WhiteListInfos").Preload("NFTConfigs").Preload("NFTConfigs.MetadataAttributes").Where(cond). /*Where("activity_id=? and contract_address=?", activity, contract).*/ Order("id DESC").Offset(offset).Limit(limit).
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

	cache, err := InitCache(poapId)
	if err != nil {
		return 0, err
	}

	return cache.Count, nil
}

func CountPOAPResultBySocial(socialId, poapId string, socialType uint) (int64, error) {
	cond := &POAPResult{
		ActivityID: poapId,
		SocialId:   socialId,
		SocialType: socialType,
	}
	var count int64
	err := db.Model(&POAPResult{}).Where(cond).Count(&count).Error
	return count, err
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

func InitCache(ActivityID string) (*POAPResultCountCache, error) {
	var count int64
	countCache, ok := Cache[ActivityID]
	if !ok {
		countCache = &POAPResultCountCache{}
		Cache[ActivityID] = &POAPResultCountCache{}
	}

	if countCache.Count == 0 {
		if err := db.Model(&POAPResult{}).Where(&POAPResult{ActivityID: ActivityID}).Count(&count).Error; err != nil {
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
