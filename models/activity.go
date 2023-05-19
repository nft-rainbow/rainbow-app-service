package models

import (
	"encoding/json"

	"time"

	"github.com/mcuadros/go-defaults"
	. "github.com/nft-rainbow/rainbow-app-service/appService-errors"
	"github.com/nft-rainbow/rainbow-app-service/models/enums"
	"github.com/pkg/errors"
	"gorm.io/datatypes"
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
		Name              string                 `form:"name"`
		ActivityId        string                 `form:"activity_id"`
		ContractAddress   *string                `form:"contract_address"`
		ExcludeNoContract bool                   `form:"exclude_no_contract"`
		ActivityStatus    []enums.ActivityStatus `form:"activity_status"`
	}
)

type UpdateActivityReq struct {
	Amount                 int32           `gorm:"type:integer" json:"amount" binding:"required"`
	Name                   string          `gorm:"type:string" json:"name" binding:"required"`
	Description            string          `gorm:"type:string" json:"description" binding:"required"`
	Command                string          `gorm:"type:string" json:"command,omitempty"`
	IsPhoneWhiteListOpened bool            `gorm:"type:bool;default:false" json:"is_phone_white_list_opened"`
	IsTokenIdOrdered       *bool           `gorm:"type:bool" json:"is_token_id_ordered" default:"true"`
	EndedTime              int64           `gorm:"type:integer" json:"end_time" default:"-1"`
	StartedTime            int64           `gorm:"type:integer" json:"start_time" default:"-1"`
	MaxMintCount           int32           `gorm:"type:varchar(256)" json:"max_mint_count" binding:"required"`
	WhiteListInfos         []WhiteListInfo `json:"white_list_infos"`
	NFTConfigs             []NFTConfig     `json:"nft_configs"`
	MetadataUri            string          `gorm:"type:string" json:"metadata_uri"` //支持模版 如 http://xx/{id}.json, 铸造时 MetadataUri 优先，若为空则根据nftconfig创建metadata
	ActivityPictureURL     string          `gorm:"type:string" json:"activity_picture_url"`
	ContractRawID          *int32          `gorm:"type:string" json:"contract_id"`
	SupportWallets         datatypes.JSON  `json:"support_wallets,omitempty" swaggertype:"array,string"` //default value: ["anyweb","cellar"]
}

func (u *UpdateActivityReq) SetDefaults() error {

	var wallets []enums.WalletType
	if len(u.SupportWallets) > 0 {
		if err := json.Unmarshal(u.SupportWallets, &wallets); err != nil {
			return err
		}
	}

	if len(wallets) == 0 {
		j, _ := json.Marshal([]enums.WalletType{enums.WALLET_ANYWEB, enums.WALLET_CELLAR})
		u.SupportWallets = j
	}

	if u.IsTokenIdOrdered == nil {
		t := true
		u.IsTokenIdOrdered = &t
	}

	defaults.SetDefaults(u)
	return nil
}

type (
	ActivityReq struct {
		UpdateActivityReq
		AppId        uint               `gorm:"index" json:"app_id" binding:"required"`
		AppName      string             `gorm:"string" json:"app_name"`
		ActivityType enums.ActivityType `gorm:"type:uint" json:"activity_type" binding:"required"`
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
		return errors.Wrap(ERR_BUSNISS_ACTIVITY_CONFIG_WRONG, "not bind contract")
	}
	// FIXME: 设置了地址白名单后，只能空投，不能页面领; v2会变更逻辑
	if len(a.WhiteListInfos) != 0 {
		return errors.New("the activity has opened the white list, could not mint by user")
	}

	if a.StartedTime != -1 && time.Now().Unix() < a.StartedTime {
		return ERR_BUSINESS_TIME_EARLY
	}

	if a.EndedTime != -1 && time.Now().Unix() > a.EndedTime {
		return ERR_BUSINESS_TIME_EXPIRED
	}

	if len(a.NFTConfigs) == 0 {
		return errors.Wrap(ERR_BUSNISS_ACTIVITY_CONFIG_WRONG, "missing nft config")
	}

	if a.Amount != -1 {
		resp, err := CountPOAPResult(a.ActivityCode, nil)
		if err != nil {
			return err
		}
		if int32(resp) >= a.Amount {
			return ERR_BUSINESS_ACTIVITY_MAX_AMOUNT_ARRIVED
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
		err := db.Debug().Preload("WhiteListInfos").Preload("NFTConfigs").Preload("NFTConfigs.MetadataAttributes").Where(&item).First(&item).Error
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

	clause := db.Debug().Model(&Activity{}).Preload("WhiteListInfos").Preload("NFTConfigs").Preload("NFTConfigs.MetadataAttributes").Where(cond)

	if _cond.ExcludeNoContract {
		clause = clause.Where("contract_raw_id is not null")
	}

	if len(_cond.ActivityStatus) > 0 {
		orClause := db
		now := time.Now().Unix()
		for _, item := range _cond.ActivityStatus {
			switch item {
			case enums.ACTIVITY_STATUS_UNSTART:
				orClause = orClause.Or(db.Where("started_time > ?", now))
			case enums.ACTIVITY_STATUS_ONGOING:
				orClause = orClause.Or(db.
					Where(db.Where("results.minted_count < activities.max_mint_count").Or("activities.max_mint_count = -1")).
					Where(db.Where("started_time <? or started_time=-1", now).Where("ended_time >? or ended_time=-1", now)))
			case enums.ACTIVITY_SINGLE_END:
				orClause = orClause.Or(db.
					Where("results.minted_count>=activities.max_mint_count and activities.max_mint_count!=-1").
					Or("activities.ended_time<? and ended_time!=-1", now))
			}
		}

		clause = clause.
			Joins("left join (select activity_code,count(*) as minted_count from poap_results group by activity_code) as results on activities.activity_code=results.activity_code").
			Where(orClause)
	}

	// 未开始 starttime>now
	// 进行中 (starttime<now || starttime==-1) and (endtime>now || endtime==-1) && (results.count<max_mint_count || max_mint_count==-1)
	// 已结束 (endtime<now and endedtime!=-1) || (results.count>=max_mint_count && max_mint_count!=-1)

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
