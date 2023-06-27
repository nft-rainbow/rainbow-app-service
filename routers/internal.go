package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/nft-rainbow/rainbow-app-service/utils/ginutils"
)

type InternalController struct{}

func (i *InternalController) GetUserInfo(c *gin.Context) {
	userId := GetIdFromJwt(c)
	ginutils.RenderRespOK(c, userId)
}
