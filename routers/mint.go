package routers

import (
	"github.com/gin-gonic/gin"
	appService_errors "github.com/nft-rainbow/rainbow-app-service/appService-errors"
	"github.com/nft-rainbow/rainbow-app-service/services"
	"github.com/nft-rainbow/rainbow-app-service/utils/ginutils"
)

type MintController struct {
	Service services.MintService
}

//	@Tags			Mints
//	@ID				MintBatchByMetaUri
//	@Summary		Batch mint by metadata uri
//	@Description	Batch mint by metadata uri
//	@security		ApiKeyAuth
//	@Produce		json
//	@Param			Authorization	header		string					true	"Bearer JWT"
//	@Param			mint_batch_dto	body		services.MintBatchDto	true	"mint_batch_dto"
//	@Success		200				{object}	models.BatchMintTask
//	@Failure		400				{object}	appService_errors.RainbowAppServiceErrorDetailInfo	"Invalid request"
//	@Failure		500				{object}	appService_errors.RainbowAppServiceErrorDetailInfo	"Internal Server error"
//	@Router			/mints/batch/by-meta-uri [post]
func (m *MintController) MintBatchByMetaUri(c *gin.Context) {
	var req services.MintBatchDto
	if err := c.ShouldBindJSON(&req); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	if err := c.ShouldBindUri(&req); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	userId := GetIdFromJwt(c)

	resp, err := m.Service.MintBatchByMetaUri(userId, &req)
	ginutils.RenderResp(c, resp, err)
}

//	@Tags			Mints
//	@ID				GetBatchMintTask
//	@Summary		Get Batch mint task
//	@Description	Get Batch mint task
//	@security		ApiKeyAuth
//	@Produce		json
//	@Param			Authorization	header		string	true	"Bearer JWT"
//	@Param			id				path		int		true	"task_id"
//	@Success		200				{object}	models.BatchMintTask
//	@Failure		400				{object}	appService_errors.RainbowAppServiceErrorDetailInfo	"Invalid request"
//	@Failure		500				{object}	appService_errors.RainbowAppServiceErrorDetailInfo	"Internal Server error"
//	@Router			/mints/batch/{id} [post]
func (m *MintController) GetBatchMintTask(c *gin.Context) {
	var id UriId
	if err := c.ShouldBindUri(&id); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	resp, err := m.Service.GetBatchMintTask(GetIdFromJwt(c), id.ID)
	ginutils.RenderResp(c, resp, err)
}
