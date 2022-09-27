package models

type MintResp struct {
	UserAddress string `form:"user_address" json:"user_address"`
	NFTAddress string `form:"nft_address" json:"nft_address"`
	Contract string `form:"advertise" json:"contract"`
	TokenID string `form:"token_id" json:"token_id"`
	Time string `json:"created_at"`
}

type MintReq struct {
	UserID string `json:"user_id" binding:"required"`
	ChannelID string `json:"channel_id" binding:"required"`
}

type MintResult struct {
	BaseModel
	UserID string `gorm:"type:varchar(256)" json:"user_id" binding:"required"`
	ContractID int32 `gorm:"type:integer" json:"contract_id" binding:"required"`
	TokenID string `gorm:"type:varchar(256)" json:"token_id" binding:"required"`
}

func StoreMintResult(req MintResult) error{
	res := GetDB().Create(&req)
	if res.Error != nil {
		return  res.Error
	}
	return nil
}




