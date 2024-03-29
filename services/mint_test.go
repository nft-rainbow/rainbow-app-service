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
	models.Init()
	middlewares.InitDashboardJwtMiddleware()

	table := []struct {
		UserId            uint
		MintBatchDto      *MintBatchDto
		ExpectTaskSuccess bool
	}{
		{1, &MintBatchDto{
			SourceType: enums.SOURCE_TYPE_ADDRESS,
			AppId:      5,
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
		{1, &MintBatchDto{
			SourceType: enums.SOURCE_TYPE_PHONE,
			AppId:      5,
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
		{1, &MintBatchDto{
			SourceType: enums.SOURCE_TYPE_ADDRESS,
			AppId:      6,
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
		bmTask, err := ms.MintBatchByMetaUri(input.UserId, input.MintBatchDto)
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
