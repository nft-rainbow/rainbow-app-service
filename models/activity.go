package models

import (
	"errors"

	"github.com/nft-rainbow/rainbow-app-service/utils"
)

type (
	NFTConfig struct {
		BaseModel
		ImageURL             string               `gorm:"type:string" json:"image_url"`
		Name                 string               `gorm:"type:string" json:"name"`
		Probability          float32              `gorm:"type:float" json:"probability"`
		MetadataAttributes   []*MetadataAttribute `json:"metadata_attributes"`
		POAPActivityConfigID uint
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
		User                 string `gorm:"type:string" json:"user"`
		Count                int32  `gorm:"type:integer" json:"count"`
		POAPActivityConfigID uint
	}
)

type (
	ActivityQueryResult struct {
		Count int64                 `json:"count"`
		Items []*POAPActivityConfig `json:"items"`
	}

	ActivityFindCondition struct {
		Name            string  `form:"name"`
		ActivityId      string  `form:"activity_id"`
		ContractAddress *string `form:"contract_address"`
	}
)

type (
	ActivityReq struct {
		AppId                  uint            `gorm:"index" json:"app_id" binding:"required"`
		Amount                 int32           `gorm:"type:integer" json:"amount" binding:"required"`
		Name                   string          `gorm:"type:string" json:"name" binding:"required"`
		Description            string          `gorm:"type:string" json:"description" binding:"required"`
		AppName                string          `gorm:"string" json:"app_name" binding:"required"`
		ActivityType           uint            `gorm:"type:uint" json:"activity_type" binding:"required"`
		Command                string          `gorm:"type:string" json:"command,omitempty"`
		IsPhoneWhiteListOpened bool            `gorm:"type:bool;default:false" json:"is_phone_white_list_opened"`
		EndedTime              int64           `gorm:"type:integer" json:"end_time" default:"-1"`
		StartedTime            int64           `gorm:"type:integer" json:"start_time" default:"-1"`
		MaxMintCount           int32           `gorm:"type:varchar(256)" json:"max_mint_count" binding:"required"`
		WhiteListInfos         []WhiteListInfo `json:"white_list_infos"`
		NFTConfigs             []NFTConfig     `json:"nft_configs"`
		MetadataUri            string          `gorm:"type:string" json:"metadata_uri"`
		ActivityPictureURL     string          `gorm:"type:string" json:"activity_picture_url"`
		ContractRawID          *int32          `gorm:"type:string" json:"contract_raw_id"` //rainbow-api contract id
	}

	POAPActivityConfig struct {
		BaseModel
		ActivityReq
		ActivityCode      string            `gorm:"type:string;index" json:"activity_code"`
		RainbowUserId     uint              `gorm:"type:integer" json:"rainbow_user_id"`
		ActivityPosterURL string            `gorm:"type:string" json:"activity_poster_url"`
		Contract          *ActivityContract `gorm:"-" json:"contract,omitempty"`
	}
)

func (p *POAPActivityConfig) NeedCommand() bool {
	return p.Command == ""
}

func (p *POAPActivityConfig) LoadBindedContract() error {
	if p.ContractRawID == nil {
		return nil
	}

	err := GetDB().Model(&ActivityContract{}).Where("contract_id=?", p.ContractRawID).First(&p.Contract).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *POAPActivityConfig) IsContractBinded() bool {
	return p.ContractRawID != nil
}

// check if mintable by user
func (p *POAPActivityConfig) VerifyMintable() error {
	if !p.IsContractBinded() {
		return errors.New("not bind contract")
	}
	// FIXME: 设置了地址白名单后，只能空投，不能页面领; v2会变更逻辑
	if len(p.WhiteListInfos) != 0 {
		return errors.New("the activity has opened the white list, could not mint by user")
	}

	switch p.ActivityType {
	case utils.BLIND_BOX:
		if len(p.NFTConfigs) == 0 {
			return errors.New("missing nft configs for blind box activity")
		}
	}
	return nil
}

func CompleteActivities(ps ...*POAPActivityConfig) error {
	for _, p := range ps {
		if err := p.LoadBindedContract(); err != nil {
			return err
		}
	}
	return nil
}

func DoAndCompleteActivity(f func() (*POAPActivityConfig, error)) (*POAPActivityConfig, error) {
	activity, err := f()
	if err != nil {
		return nil, err
	}
	if err := CompleteActivities(activity); err != nil {
		return nil, err
	}
	return activity, nil
}

func FindActivity(name string, contractId int32) (*POAPActivityConfig, error) {
	return DoAndCompleteActivity(func() (*POAPActivityConfig, error) {
		var item POAPActivityConfig
		err := db.Model(&POAPActivityConfig{}).Where("name = ?", name).Where("contract_id = ?", contractId).First(&item).Error
		return &item, err
	})
}

func FindActivityByCode(activityCode string) (*POAPActivityConfig, error) {
	return DoAndCompleteActivity(func() (*POAPActivityConfig, error) {
		if activityCode == "" {
			return nil, errors.New("activity_id is required")
		}

		var item POAPActivityConfig
		item.ActivityCode = activityCode
		err := db.Model(&POAPActivityConfig{}).Where(&item).Find(&item).Error
		if err != nil {
			return nil, err
		}
		err = db.Preload("WhiteListInfos").Preload("NFTConfigs").Preload("NFTConfigs.MetadataAttributes").Find(&item).Error
		if err != nil {
			return nil, err
		}

		return &item, err
	})
}

func FindAndCountActivity(ranbowUserId uint, offset int, limit int, _cond ActivityFindCondition) (*ActivityQueryResult, error) {
	var items []*POAPActivityConfig
	cond := &POAPActivityConfig{}
	cond.RainbowUserId = ranbowUserId
	cond.Name = _cond.Name
	cond.ActivityCode = _cond.ActivityId

	clause := db.Debug().Model(&POAPActivityConfig{}).Preload("WhiteListInfos").Preload("NFTConfigs").Preload("NFTConfigs.MetadataAttributes").Where(cond)

	// _cond.ContractAddress 如果不为空，查找Contract, 拿到 contract_raw_id
	if _cond.ContractAddress != nil {
		contract, err := FirstContract(ActivityContract{ContractAddress: *_cond.ContractAddress})
		if err != nil {
			return nil, err
		}
		clause = clause.Where("contract_raw_id", contract.ContractRawID)
	}

	var count int64
	if err := clause.Count(&count).Error; err != nil {
		return nil, err
	}

	if err := clause.Order("id DESC").Offset(offset).Limit(limit).
		Find(&items).Error; err != nil {
		return nil, err
	}

	if err := CompleteActivities(items...); err != nil {
		return nil, err
	}

	return &ActivityQueryResult{count, items}, nil
}
