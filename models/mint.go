package models

type CustomMintResp struct {
	UserAddress string `form:"user_address" json:"user_address"`
	NFTAddress  string `form:"nft_address" json:"nft_address"`
	Contract    string `form:"contract" json:"contract"`
	TokenID     string `form:"token_id" json:"token_id"`
	Time        string `json:"created_at"`
}

type CustomMintReq struct {
	UserID    string `json:"user_id" binding:"required"`
	ChannelID string `json:"channel_id" binding:"required"`
}

type CustomMintResult struct {
	BaseModel
	UserID     string `gorm:"type:varchar(256)" json:"user_id"`
	ContractID int32  `gorm:"type:integer" json:"contract_id"`
	TokenID    string `gorm:"type:varchar(256)" json:"token_id"`
	Hash       string `gorm:"type:string" json:"hash"`
	Status     int32  `gorm:"type:integer" json:"status"`
}

type POAPResult struct {
	BaseModel
	Address     string `gorm:"type:string;index" json:"address" binding:"required"`
	ConfigID    int32  `gorm:"type:integer" json:"config_id"`
	ContractID  int32  `gorm:"type:integer" json:"contract_id" binding:"required"`
	TxID        int32  `gorm:"type:integer" json:"tx_id"`
	TokenID     string `gorm:"type:varchar(256)" json:"token_id"`
	Hash        string `gorm:"type:string" json:"hash"`
	ActivityID  string `gorm:"type:string;index" json:"activity_id"`
	Status      int32  `gorm:"type:integer;index" json:"status"`
	FileURL     string `gorm:"type:string" json:"file_url"`
	ProjectorId uint   `gorm:"type:integer" json:"projector_id"`
	AppId       uint   `gorm:"type:integer" json:"app_id"`
	SocialId    string `gorm:"type:string;index" json:"social_id"`
	SocialType  uint   `gorm:"type:integer" json:"social_type"`
}

func StoreCustomMintResult(req CustomMintResult) error {
	res := GetDB().Create(&req)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
