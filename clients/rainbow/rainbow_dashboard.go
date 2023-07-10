package rainbow

import (
	"fmt"
	"net/http"

	"github.com/nft-rainbow/rainbow-app-service/config"
	"github.com/nft-rainbow/rainbow-app-service/middlewares"
	"github.com/nft-rainbow/rainbow-app-service/models/enums"
	"github.com/nft-rainbow/rainbow-app-service/utils"
	"github.com/pkg/errors"
)

type RainbowApiDashboardClient struct {
	baseUri string
}

func NewRainbowApiDashboardClient(baseUri ...string) *RainbowApiDashboardClient {
	if len(baseUri) == 0 {
		baseUri = append(baseUri, config.GetConfig().RainbowDashboardAPI)
	}

	return &RainbowApiDashboardClient{baseUri[0]}
}

type (
	ContractInfoDtoWithoutType struct {
		Chain           enums.Chain `form:"chain" json:"chain" binding:"required,oneof=conflux conflux_test"`
		ContractAddress string      `form:"contract_address" json:"contract_address" binding:"required"`
	}

	MintItemDto struct {
		MintToAddress string `form:"mint_to_address" json:"mint_to_address" binding:"required"`
		TokenId       string `form:"token_id" json:"token_id"`
		Amount        *uint  `form:"amount" json:"amount"`
		MetadataUri   string `form:"metadata_uri" json:"metadata_uri"`
	}

	CustomMintBatchDto struct {
		ContractInfoDtoWithoutType
		MintItems []*MintItemDto `form:"mint_items" json:"mint_items" binding:"required,dive"`
	}
)

func (r *RainbowApiDashboardClient) MintBatchByMetauri(userId uint, appId uint, req *CustomMintBatchDto) ([]uint, error) {
	// 1. gen token
	token, err := middlewares.GenerateRainbowJWT(userId)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to gen auth token")
	}
	// 2. send request
	url := fmt.Sprintf("%v/dashboard/apps/%d/nft/batch/by-meta-uri", r.baseUri, appId)
	headers := map[string]string{
		"Authorization": token,
	}

	mintTaskIds, err := utils.SendHttp[any, []uint](http.MethodPost, url, req, headers)
	if err != nil {
		return nil, err
	}

	return *mintTaskIds, nil
}
