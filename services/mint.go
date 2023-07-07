package services

import (
	"github.com/nft-rainbow/rainbow-app-service/models/enums"
	"github.com/pkg/errors"
)

type MintItemDto struct {
	MintTo      string `form:"mint_to" json:"mint_to" binding:"required"`
	TokenId     string `form:"token_id" json:"token_id"`
	Amount      *uint  `form:"amount" json:"amount"`
	MetadataUri string `form:"metadata_uri" json:"metadata_uri"`
}

type CustomMintBatchDto struct {
	SourceType      enums.SourceType
	Chain           string         `form:"chain" json:"chain" binding:"required,oneof=conflux conflux_test"`
	ContractAddress string         `form:"contract_address" json:"contract_address" binding:"required"`
	MintItems       []*MintItemDto `form:"mint_items" json:"mint_items" binding:"required,dive"`
}

type MintBatchResp struct {
	Status      enums.BatchMintStatus
	Error       string
	MintTaskIds []uint `json:"mint"`
}

// 1. map source to address
// 2. call rainbow-api
func MintBatch(req *CustomMintBatchDto) error {
	var sources []string
	for _, item := range req.MintItems {
		sources = append(sources, item.MintTo)
	}

	// exists, unexist, err := (&AddressFinder{req.SourceType}).Find(sources)
	// if len(unexist) > 0 {
	// 	// create wallet
	// }
	return nil

}

func MintBatchFromCerti() {

}

func MintBatchViaRainbowApi(req *CustomMintBatchDto) error {
	if req.SourceType != enums.SOURCE_TYPE_ADDRESS {
		return errors.New("source type must be address")
	}
	// call rainbow-api
	return nil
}
