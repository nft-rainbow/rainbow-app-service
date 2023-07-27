package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/mcuadros/go-defaults"
	appService_errors "github.com/nft-rainbow/rainbow-app-service/appService-errors"
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/utils/ginutils"
)

func Pagination() gin.HandlerFunc {
	return func(c *gin.Context) {
		var pagination models.Pagination
		defaults.SetDefaults(&pagination)

		if err := c.ShouldBindQuery(&pagination); err != nil {
			ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_PAGINATION)
			return
		}

		c.Set("page", pagination.Page)
		c.Set("limit", pagination.Limit)
		c.Set("offset", pagination.Offset())
		c.Next()
	}
}
