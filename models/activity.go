package models

import (
	"errors"
	"time"

	"github.com/mcuadros/go-defaults"
	"github.com/nft-rainbow/rainbow-app-service/models/enums"
)

type (
	NFTConfig struct {
		BaseModel
		ImageURL           string               `gorm:"type:string" json:"image_url"`
		Name               string               `gorm:"type:string" json:"name"`
		Probability        float32              `gorm:"type:float" json:"probability"`
		MetadataAttributes []*MetadataAttribute `json:"metadata_attributes"`
		ActivityID         uint
	}

	MetadataAttribute struct {
		BaseModel
		TraitType   string `gorm:"type:varchar(256)"  json:"trait_type"`
		DisplayType string `gorm:"type:varchar(256)"  json:"display_type,omitempty"`
		Value       string `gorm:"type:varchar(256)"  json:"value"`
		NFTConfigID uint
	}

	WhiteListInfo struct {
		BaseModel
		User       string `gorm:"type:string" json:"user"`
		Count      int32  `gorm:"type:integer" json:"count"`
		ActivityID uint
	}
)

type (
	ActivityQueryResult struct {
		Count int64       `json:"count"`
		Items []*Activity `json:"items"`
	}

	ActivityFindCondition struct {
		Pagination
		Name            string  `form:"name"`
		ActivityId      string  `form:"activity_id"`
		ContractAddress *string `form:"contract_address"`
	}
)

type (
	UpdateActivityReq struct {
		Amount                 int32           `gorm:"type:integer" json:"amount" binding:"required"`
		Name                   string          `gorm:"type:string" json:"name" binding:"required"`
		Description            string          `gorm:"type:string" json:"description" binding:"required"`
		Command                string          `gorm:"type:string" json:"command,omitempty"`
		IsPhoneWhiteListOpened bool            `gorm:"type:bool;default:false" json:"is_phone_white_list_opened"`
		EndedTime              int64           `gorm:"type:integer" json:"end_time" default:"-1"`
		StartedTime            int64           `gorm:"type:integer" json:"start_time" default:"-1"`
		MaxMintCount           int32           `gorm:"type:varchar(256)" json:"max_mint_count" binding:"required"`
		WhiteListInfos         []WhiteListInfo `json:"white_list_infos"`
		NFTConfigs             []NFTConfig     `json:"nft_configs"`
		MetadataUri            string          `gorm:"type:string" json:"metadata_uri"` //与contract base_uri 的区别：base_uri需拼接tokenid后缀；
		ActivityPictureURL     string          `gorm:"type:string" json:"activity_picture_url"`
		ContractRawID          *int32          `gorm:"type:string" json:"contract_id"`
	}
	ActivityReq struct {
		UpdateActivityReq
		AppId            uint               `gorm:"index" json:"app_id" binding:"required"`
		AppName          string             `gorm:"string" json:"app_name"`
		ActivityType     enums.ActivityType `gorm:"type:uint" json:"activity_type" binding:"required"`
		IsTokenIdOrdered *bool              `gorm:"type:bool" json:"is_token_id_ordered" default:"true"`
	}

	Activity struct {
		BaseModel
		ActivityReq
		ActivityCode      string    `gorm:"type:string;index" json:"activity_id"` //TODO: 与前端统一调整为activity_code
		RainbowUserId     uint      `gorm:"type:integer" json:"rainbow_user_id"`
		ActivityPosterURL string    `gorm:"type:string" json:"activity_poster_url"`
		Contract          *Contract `gorm:"-" json:"contract,omitempty"`
	}
)

func (a *Activity) NeedCommand() bool {
	return a.Command == ""
}

func (a *Activity) LoadBindedContract() error {
	if a.ContractRawID == nil {
		return nil
	}

	err := GetDB().Model(&Contract{}).Where("contract_raw_id=?", a.ContractRawID).First(&a.Contract).Error
	if err != nil {
		return err
	}
	return nil
}

func (a *Activity) IsContractBinded() bool {
	return a.ContractRawID != nil
}

// check if mintable by user
func (a *Activity) VerifyMintable() error {
	if !a.IsContractBinded() {
		return errors.New("not bind contract")
	}
	// FIXME: 设置了地址白名单后，只能空投，不能页面领; v2会变更逻辑
	if len(a.WhiteListInfos) != 0 {
		return errors.New("the activity has opened the white list, could not mint by user")
	}

	if a.StartedTime != -1 && time.Now().Unix() < a.StartedTime {
		return errors.New("the activity has not been started")
	}

	if a.EndedTime != -1 && time.Now().Unix() > a.EndedTime {
		return errors.New("the activity has been expired")
	}

	switch a.ActivityType {
	case enums.ACTIVITY_BLINDBOX:
		if len(a.NFTConfigs) == 0 {
			return errors.New("missing nft configs for blind box activity")
		}
	}

	if a.Amount != -1 {
		resp, err := CountPOAPResult(a.ActivityCode, nil)
		if err != nil {
			return err
		}
		if int32(resp) >= a.Amount {
			return errors.New("the mint amount has exceeded the limit")
		}
	}

	return nil
}

func CompleteActivities(ps ...*Activity) error {
	for _, p := range ps {
		if err := p.LoadBindedContract(); err != nil {
			return err
		}
	}
	return nil
}

func DoAndCompleteActivity(f func() (*Activity, error)) (*Activity, error) {
	activity, err := f()
	if err != nil {
		return nil, err
	}
	if err := CompleteActivities(activity); err != nil {
		return nil, err
	}
	return activity, nil
}

func FindActivity(name string, contractId int32) (*Activity, error) {
	return DoAndCompleteActivity(func() (*Activity, error) {
		var item Activity
		err := db.Model(&Activity{}).Where("name = ?", name).Where("contract_id = ?", contractId).First(&item).Error
		return &item, err
	})
}

func FindActivityByCode(activityCode string) (*Activity, error) {
	return DoAndCompleteActivity(func() (*Activity, error) {
		if activityCode == "" {
			return nil, errors.New("activity_code is required")
		}

		var item Activity
		item.ActivityCode = activityCode
		err := db.Preload("WhiteListInfos").Preload("NFTConfigs").Preload("NFTConfigs.MetadataAttributes").Where(&item).First(&item).Error
		if err != nil {
			return nil, err
		}

		return &item, err
	})
}

func FindAndCountActivity(ranbowUserId uint, _cond ActivityFindCondition) (*ActivityQueryResult, error) {
	defaults.SetDefaults(&_cond)

	var items []*Activity
	cond := &Activity{}
	cond.RainbowUserId = ranbowUserId
	cond.Name = _cond.Name
	cond.ActivityCode = _cond.ActivityId

	clause := db.Model(&Activity{}).Preload("WhiteListInfos").Preload("NFTConfigs").Preload("NFTConfigs.MetadataAttributes").Where(cond)

	// _cond.ContractAddress 如果不为空，查找Contract, 拿到 contract_raw_id
	if _cond.ContractAddress != nil {
		contract, err := FirstContract(Contract{ContractAddress: *_cond.ContractAddress})
		if err != nil {
			return nil, err
		}
		clause = clause.Where("contract_raw_id", contract.ContractRawID)
	}

	var count int64
	if err := clause.Count(&count).Error; err != nil {
		return nil, err
	}

	if err := clause.Order("id DESC").Offset(_cond.Offset()).Limit(_cond.Limit).Find(&items).Error; err != nil {
		return nil, err
	}

	if err := CompleteActivities(items...); err != nil {
		return nil, err
	}

	return &ActivityQueryResult{count, items}, nil
}
