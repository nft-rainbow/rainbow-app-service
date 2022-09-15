package models

type ContractDeployDto struct {
	Chain                     string `form:"chain" json:"chain" binding:"required,oneof=conflux conflux_test"`
	Name                      string `form:"name" json:"name" binding:"required"`
	Symbol                    string `form:"symbol" json:"symbol" binding:"required"`
	OwnerAddress              string `form:"owner_address" json:"owner_address" binding:"required"`
	Type                      string `form:"type" json:"type" binding:"required,oneof=erc721 erc1155"`
	BaseUri                   string `form:"base_uri" json:"base_uri"`
	RoyaltiesBps              uint   `form:"royalties_bps" json:"royalties_bps"`
	RoyaltiesAddress          string `form:"royalties_address" json:"royalties_address"`
	TokensBurnableByAdmin     bool   `form:"tokens_burnable_admin" json:"tokens_burnable_admin"`
	TokensTransferableByAdmin bool   `form:"tokens_transferable_admin" json:"tokens_transferable_admin"`
	TransferCooldownTime      uint   `form:"transfer_cooldown_time" json:"transfer_cooldown_time"`
}

type Contract struct {
	BaseModel
	AppId        uint   `gorm:"index" json:"app_id"`
	ChainType    uint   `gorm:"type:uint" json:"chain_type"`
	ChainId      uint   `gorm:"type:uint;index" json:"chain_id"`
	Address      string `gorm:"type:varchar(256);index" json:"address"`
	OwnerAddress string `gorm:"type:varchar(256);index" json:"owner_address"`
	Type         uint   `gorm:"type:int" json:"type"` // 1-ERC721, 2-ERC1155
	BaseUri      string `gorm:"type:varchar(256)" json:"base_uri"`
	Name         string `gorm:"type:varchar(256)" json:"name"`
	Symbol       string `gorm:"type:varchar(256)" json:"symbol"`
	Hash         string `gorm:"type:varchar(256)" json:"hash"`
	TxId         uint   `gorm:"index" json:"tx_id"`
	Status       uint   `json:"status"` // 0-pending, 1-success, 2-failed
}

