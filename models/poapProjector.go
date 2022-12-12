package models

type POAPActivityConfig struct {
	BaseModel
	ContractID int32 `gorm:"type:integer" json:"contract_id" binding:"required"`
	Amount int32 `gorm:"type:integer" json:"amount" binding:"required"`
	Name string `gorm:"type:string" json:"name" binding:"required"`
	Description string `gorm:"type:string" json:"description" binding:"required"`
	AppId int32 `gorm:"index" json:"app_id" binding:"required"`
	ContractType int32 `gorm:"type:int" json:"contract_type"`
	ContractAddress string `gorm:"type:string" json:"contract_address"`
	Chain    int32   `gorm:"type:int" json:"chain_type"`
	MetadataURI string `gorm:"type:string" json:"metadata_uri" binding:"required"`
	IsCommandNeeded bool `gorm:"type:bool" json:"is_command_needed"`
	Command string `gorm:"type:string" json:"command"`
	EndedTime int64 `gorm:"type:integer" json:"end_time"`
	StartedTime int64 `gorm:"type:integer" json:"start_time"`
	RainbowUserId int32 `gorm:"type:integer" json:"rainbow_user_id"`
}

type H5Config struct {
	BaseModel
	ActivityId int32 `gorm:"type:integer" json:"activity_id"`
	Link string `gorm:"type:string" json:"link" binding:"required"`
	Title string `gorm:"type:string" json:"title"`
	TitleSize int32 `gorm:"type:integer" json:"title_size"`
	TitleColor int32 `gorm:"type:integer" json:"title_color"`
	Content string `gorm:"type:string" json:"content"`
	ContentSize int32 `gorm:"type:integer" json:"content_size"`
	ContentColor int32 `gorm:"type:integer" json:"content_color"`
	ClaimButtonColor string `gorm:"type:string" json:"claim_button_color"`
	ButtonWordColor int32 `gorm:"type:string" json:"button_word_color"`
	LogoURL string `gorm:"type:string" json:"logo_url"`
	PCPicURL string `gorm:"type:string" json:"pc_picture_url"`
	MobilePicURL string `gorm:"type:string" json:"mobile_picture_url"`
}

type POAPActivityQueryResult struct {
	Count int64       `json:"count"`
	Items []*POAPActivityConfig `json:"items"`
}

type POAPResultQueryResult struct {
	Count int64       `json:"count"`
	Items []*POAPResult `json:"items"`
}

func FindPOAPActivityConfig(name string, contractId int32) (*POAPActivityConfig, error){
	var item POAPActivityConfig
	err := db.Where("name = ?", name).Where("contract_id = ?", contractId).First(&item).Error
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
	cond.RainbowUserId = int32(id)

	var count int64
	if err := db.Find(&items).Where(cond).Count(&count).Error; err != nil {
		return nil, err
	}

	if err := db.Find(&items).Where(cond).Offset(offset).Limit(limit).Error; err != nil {
		return nil, err
	}

	return &POAPActivityQueryResult{count, items}, nil
}

func FindAndCountPOAPResult(activityId, offset int, limit int) (*POAPResultQueryResult, error) {
	var items []*POAPResult
	cond := &POAPResult{}
	cond.ActivityID = int32(activityId)

	var count int64
	if err := db.Find(&items).Where(cond).Count(&count).Error; err != nil {
		return nil, err
	}

	if err := db.Find(&items).Where(cond).Offset(offset).Limit(limit).Error; err != nil {
		return nil, err
	}

	return &POAPResultQueryResult{count, items}, nil
}

func FindPOAPResultById(activityId, id int) (*POAPResult, error) {
	cond := &POAPResult{}
	resp := &POAPResult{}
	cond.ActivityID = int32(activityId)
	cond.ID = uint(id)
	if err := db.Where(cond).First(&resp).Error; err != nil {
		return nil, err
	}

	return resp, nil
}