package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/nft-rainbow/discordbot-service/utils/ginutils"
	"net/http"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/", indexEndpoint)
	admin := router.Group("/channel")
	admin.POST("/admin", bindAdminConfig)
	admin.POST("/contract", customMintConfig)
	admin.POST("/easy", easyMintConfig)

	user := router.Group("/user")
	user.POST("/address", bindUserAddress)
	user.GET("/address/:user_id", getUserBindingAddress)
	user.POST("/mint/custom", handleCustomMint)
	user.POST("/mint/easy", handleEasyMint)
}

func indexEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, ginutils.DataResponse("Discord-Bot-Service"))
}