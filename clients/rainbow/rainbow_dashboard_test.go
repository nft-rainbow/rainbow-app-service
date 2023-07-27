package rainbow

import (
	"testing"

	"github.com/nft-rainbow/rainbow-app-service/config"
	"github.com/nft-rainbow/rainbow-app-service/middlewares"
	"github.com/nft-rainbow/rainbow-app-service/models/enums"
	"github.com/stretchr/testify/assert"
)

func TestMintbatchByMetauri(t *testing.T) {
	config.InitByFile("../../config.yaml")
	middlewares.InitDashboardJwtMiddleware()
	client := NewRainbowApiDashboardClient("http://localhost:8080")
	taskIds, err := client.MintBatchByMetauri(1, 5, &CustomMintBatchDto{
		ContractInfoDtoWithoutType: ContractInfoDtoWithoutType{
			Chain:           enums.CHAIN_CONFLUX_TEST,
			ContractAddress: "cfxtest:acbwua8x80xfr5nf9g60xgn3b4m0we7ddu01n695wn",
		},
		MintItems: []*MintItemDto{
			{
				MintToAddress: "cfxtest:acbwua8x80xfr5nf9g60xgn3b4m0we7ddu01n695wn",
				MetadataUri:   "http://xxx/xxx.json",
			},
		},
	})
	assert.NoError(t, err)
	assert.Equal(t, 1, len(taskIds))
}
