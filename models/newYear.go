package models

type NewYearConfig struct {
	BaseModel
	Amount int32 `gorm:"type:integer" json:"amount" binding:"required"`
	Name string `gorm:"type:string" json:"name" binding:"required"`
	Description string `gorm:"type:string" json:"description" binding:"required"`
	AppId int32 `gorm:"index" json:"app_id" binding:"required"`
	Chain    int32   `gorm:"type:int" json:"chain_type"`
	EndedTime int64 `gorm:"type:integer" json:"end_time" binding:"required"`
	StartedTime int64 `gorm:"type:integer" json:"start_time" binding:"required"`
	RainbowUserId int32 `gorm:"type:integer" json:"rainbow_user_id"`
	MaxMintCount uint `gorm:"type:varchar(256)" json:"max_mint_count" binding:"required"`
	ActivityPictureURL string `gorm:"type:string" json:"activity_picture_url"`
	SharingContent string `gorm:"type:string" json:"sharing_content"`
	ContractInfos []NFTContractInfo `json:"nft_contract_infos"`
}

type MintCount struct {
	BaseModel
	Address string `gorm:"type:string" json:"address"`
	Count int32 `gorm:"type:integer" json:"count"`
	ActivityID uint `gorm:"type:integer" json:"activity_id"`
}

type NFTContractInfo struct {
	BaseModel
	MetadataURI string `gorm:"type:string" json:"metadata_uri" binding:"required"`
	Probability float32 `gorm:"type:varchar(256)" json:"probability" binding:"required"`
	ContractType int32 `gorm:"type:integer" json:"contract_type"`
	ContractAddress string `gorm:"type:string" json:"contract_address"`
	ContractID int32 `gorm:"type:integer" json:"contract_id" binding:"required"`
	NewYearConfigID uint
}

type NewYearConfigQueryResult struct {
	Count int64       `json:"count"`
	Items []*NewYearConfig `json:"items"`
}

type NewYearSpecialConfigQueryResult struct {
	Count int64       `json:"count"`
	Items []*NewYearConfig `json:"items"`
}

func FindNewYearConfigById(id int) (*NewYearConfig, error){
	var item NewYearConfig
	err := db.Where("id = ?", id).First(&item).Error
	if err != nil {
		return nil, err
	}

	err = db.Preload("ContractInfos").Find(&item).Error
	if err != nil {
		return nil, err
	}
	return &item, err
}

func FindMintCount(address string, activityId int32) (*MintCount, error){
	var count int64
	var cond MintCount
	cond.Address = address
	cond.ActivityID = uint(activityId)
	cond.Count = 0
	if err := db.Find(&cond).Count(&count).Error; err != nil {
		return nil, err
	}

	if count == 0 {
		item := &MintCount{
			Address: address,
			Count: 0,
			ActivityID: uint(activityId),
		}
		db.Create(item)
		return item, nil
	}
	var item MintCount
	err := db.Where("activity_id = ?", activityId).Where("address = ?", address).Last(&item).Error
	return &item, err
}

func UpdateMintCount(address string, activityId, updateCount int32) (*MintCount, error){
	item, err := FindMintCount(address, activityId)
	if err != nil {
		return nil, err
	}
	item.Count += updateCount

	db.Save(&item)

	return item, nil
}