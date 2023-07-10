package services

import (
	"fmt"
	"testing"
	"time"

	"github.com/nft-rainbow/rainbow-app-service/clients/rainbow"
	"github.com/nft-rainbow/rainbow-app-service/config"
	"github.com/nft-rainbow/rainbow-app-service/logger"
	"github.com/nft-rainbow/rainbow-app-service/middlewares"
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/models/enums"
	"github.com/stretchr/testify/assert"
)

func TestMintBatch(t *testing.T) {
	config.Init()
	logger.Init()
	models.ConnectDB()
	middlewares.InitDashboardJwtMiddleware()

	table := []struct {
		UserId            uint
		AppId             uint
		MintBatchDto      *MintBatchDto
		ExpectTaskSuccess bool
	}{
		{1, 5, &MintBatchDto{
			SourceType: enums.SOURCE_TYPE_ADDRESS,
			ContractInfoDtoWithoutType: rainbow.ContractInfoDtoWithoutType{
				Chain:           enums.CHAIN_CONFLUX_TEST,
				ContractAddress: "cfxtest:acbwua8x80xfr5nf9g60xgn3b4m0we7ddu01n695wn",
			},
			MintItems: []*MintItemDto{
				{
					MintTo:      "cfxtest:acbwua8x80xfr5nf9g60xgn3b4m0we7ddu01n695wn",
					MetadataUri: "http://xxx/xxx.json",
				},
			},
		}, true},
		{1, 5, &MintBatchDto{
			SourceType: enums.SOURCE_TYPE_PHONE,
			ContractInfoDtoWithoutType: rainbow.ContractInfoDtoWithoutType{
				Chain:           enums.CHAIN_CONFLUX_TEST,
				ContractAddress: "cfxtest:acbwua8x80xfr5nf9g60xgn3b4m0we7ddu01n695wn",
			},
			MintItems: []*MintItemDto{
				{
					MintTo:      "13983211056",
					MetadataUri: "http://xxx/xxx.json",
				},
			},
		}, true},
		{1, 6, &MintBatchDto{
			SourceType: enums.SOURCE_TYPE_ADDRESS,
			ContractInfoDtoWithoutType: rainbow.ContractInfoDtoWithoutType{
				Chain:           enums.CHAIN_CONFLUX_TEST,
				ContractAddress: "cfxtest:acbwua8x80xfr5nf9g60xgn3b4m0we7ddu01n695wn",
			},
			MintItems: []*MintItemDto{
				{
					MintTo:      "cfxtest:acbwua8x80xfr5nf9g60xgn3b4m0we7ddu01n695wn",
					MetadataUri: "http://xxx/xxx.json",
				},
			},
		}, false},
	}

	ms := &MintService{}
	for _, input := range table {
		bmTask, err := ms.MintBatchByMetaUri(input.UserId, input.AppId, input.MintBatchDto)
		assert.NoError(t, err)
		assert.NotNil(t, bmTask)

		tick := time.NewTicker(time.Second * 2)
		for {
			<-tick.C
			err = models.GetDB().Debug().Model(&models.BatchMintTask{}).Where("id=?", bmTask.ID).First(&bmTask).Error
			assert.NoError(t, err)
			if bmTask.IsFinalized() {
				tick.Stop()
				break
			}
		}

		assert.Equal(t, input.ExpectTaskSuccess, bmTask.IsSuceess())
		fmt.Printf("batch mint result: %+v", *bmTask)
	}

}
