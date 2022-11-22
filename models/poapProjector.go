package models


type POAPProjectorConfig struct {
	BaseModel
	AppId int32 `gorm:"index" json:"app_id" binding:"required"`
	RainbowUserId int32 `gorm:"type:integer" json:"rainbow_user_id"`
	ProjectorName string `gorm:"type:string" json:"projector_name" binding:"required"`
	Description string `gorm:"type:string" json:"description" binding:"required"`
	ChainType string `gorm:"type:string" json:"chain_type" binding:"required"`
}

type POAPActivityConfig struct {
	BaseModel
	ContractID int32 `gorm:"type:integer" json:"contract_id" binding:"required"`
	Amount int32 `gorm:"type:integer" json:"amount" binding:"required"`
	Name string `gorm:"type:string" json:"name" binding:"required"`
	Description string `gorm:"type:string" json:"description" binding:"required"`
	AppId int32 `gorm:"index" json:"app_id"`
	ContractType int32 `gorm:"type:int" json:"contract_type"`
	ContractAddress string `gorm:"type:string" json:"contract_address"`
	Chain    int32   `gorm:"type:int" json:"chain_type"`
	MetadataURI string `gorm:"type:string" json:"metadata_uri" binding:"required"`
	IsCommandNeeded bool `gorm:"type:bool" json:"is_command_needed"`
	Command string `gorm:"type:string" json:"command"`
	EndedTime int64 `gorm:"type:integer" json:"end_time"`
	StartedTime int64 `gorm:"type:integer" json:"start_time"`
}

type POAPActivityQueryResult struct {
	Count int64       `json:"count"`
	Items []*POAPActivityConfig `json:"items"`
}

type POAPProjectorConfigQueryResult struct {
	Count int64       `json:"count"`
	Items []*POAPProjectorConfig `json:"items"`
}

func FindPOAPActivityConfig(name string, contractId int32) (*POAPActivityConfig, error){
	var item POAPActivityConfig
	err := db.Where("name = ?", name).Where("contract_id = ?", contractId).First(&item).Error
	return &item, err
}

func FindPOAPConfigById(id int) (*POAPProjectorConfig, error) {
	var item POAPProjectorConfig
	err := db.Where("id = ?", id).First(&item).Error
	return &item, err
}

func FindPOAPConfigByAppId(id int) (*POAPProjectorConfig, error) {
	var item POAPProjectorConfig
	err := db.Where("app_id = ?", id).First(&item).Error
	return &item, err
}

func FindPOAPConfigByUserId(id int) (*POAPProjectorConfig, error) {
	var item POAPProjectorConfig
	err := db.Where("rainbow_user_id = ?", id).First(&item).Error
	return &item, err
}

func FindPOAPActivityConfigById(id int) (*POAPActivityConfig, error){
	var item POAPActivityConfig
	err := db.Where("id = ?", id).First(&item).Error
	return &item, err
}

func FindAndCountPOAPActivity(id uint, offset int, limit int) (*POAPActivityQueryResult, error) {
	var items []*POAPActivityConfig
	cond := &POAPActivityConfig{}
	cond.AppId = int32(id)

	var count int64
	if err := db.Find(&items).Where(cond).Count(&count).Error; err != nil {
		return nil, err
	}

	if err := db.Find(&items).Where(cond).Offset(offset).Limit(limit).Error; err != nil {
		return nil, err
	}

	return &POAPActivityQueryResult{count, items}, nil
}

func FindAndCountPOAPProjectorConfig(id uint, offset int, limit int) (*POAPProjectorConfigQueryResult, error) {
	var items []*POAPProjectorConfig
	cond := &POAPProjectorConfig{}
	cond.RainbowUserId = int32(id)

	var count int64
	if err := db.Find(&items).Where(cond).Count(&count).Error; err != nil {
		return nil, err
	}

	if err := db.Find(&items).Where(cond).Offset(offset).Limit(limit).Error; err != nil {
		return nil, err
	}

	return &POAPProjectorConfigQueryResult{count, items}, nil
}