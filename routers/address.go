package routers

import (
	"github.com/gin-gonic/gin"
	appService_errors "github.com/nft-rainbow/rainbow-app-service/appService-errors"
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/services"
	"github.com/nft-rainbow/rainbow-app-service/utils/ginutils"
)

func bindProjectorConfig(c *gin.Context) {
	var adminConfig *models.AdminConfig
	if err := c.ShouldBind(&adminConfig); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	err := services.BindProjectorConfig(adminConfig, GetIdFromJwtClaim(c))
	ginutils.RenderResp(c, "success", err)
}

