package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/nft-rainbow/discordbot-service/middlewares"
	"github.com/nft-rainbow/discordbot-service/utils/ginutils"
	"net/http"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/", indexEndpoint)
	discord := router.Group("/discord")
	discord.POST("/login", middlewares.LoginHandler)

	admin := discord.Group("/channel")
	admin.Use(middlewares.OpenJwtAuthMiddleware.MiddlewareFunc())
	{
		admin.POST("/config/user", bindAdminConfig)
		admin.GET("/:guild_id/channels", getChannelInfo)
		admin.POST("/config/event", customMintConfig)

	}

	user := discord.Group("/user")
	user.Use(middlewares.JwtAuthMiddleware.MiddlewareFunc())
	{
		user.POST("/mint", handleCustomMint)
		user.GET("/address/:user_id", getUserBindingAddress)
		user.POST("/address", bindUserAddress)
	}
}

func indexEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, ginutils.DataResponse("Rainbow-App-Service"))
}