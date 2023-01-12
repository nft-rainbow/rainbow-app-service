package models

import "gorm.io/gorm"

type CustomMintResp struct {
	UserAddress string `form:"user_address" json:"user_address"`
	NFTAddress string `form:"nft_address" json:"nft_address"`
	Contract string `form:"advertise" json:"contract"`
	TokenID string `form:"token_id" json:"token_id"`
	Time string `json:"created_at"`
}

type CustomMintReq struct {
	UserID string `json:"user_id" binding:"required"`
	ChannelID string `json:"channel_id" binding:"required"`
}

type CustomMintResult struct {
	BaseModel
	UserID string `gorm:"type:varchar(256)" json:"user_id"`
	ContractID int32 `gorm:"type:integer" json:"contract_id"`
	TokenID string `gorm:"type:varchar(256)" json:"token_id"`
	Hash string `gorm:"type:string" json:"hash"`
}

type POAPResult struct {
	BaseModel
	Address string `gorm:"type:string" json:"address" binding:"required"`
	ActivityID int32 `gorm:"type:integer" json:"activity_id"`
	ContractID int32 `gorm:"type:integer" json:"contract_id" binding:"required"`
	TxID int32 `gorm:"type:integer" json:"tx_id"`
	TokenID string `gorm:"type:varchar(256)" json:"token_id"`
	Hash string `gorm:"type:string" json:"hash"`
}

func StoreCustomMintResult(req CustomMintResult) error{
	res := GetDB().Create(&req)
	if res.Error != nil {
		return  res.Error
	}
	return nil
}

func StorePOAPResult(req POAPResult) (*gorm.DB, error){
	res := GetDB().Create(&req)
	if res.Error != nil {
		return nil, res.Error
	}
	return res, nil
}