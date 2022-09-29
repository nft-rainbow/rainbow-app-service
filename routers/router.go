package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nft-rainbow/discordbot-service/middlewares"
	"github.com/nft-rainbow/discordbot-service/utils/ginutils"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/", indexEndpoint)
	discord := router.Group("/discord")
	discord.POST("/login", middlewares.LoginHandler)

	discord.GET("/:guild_id/channels", getChannelInfo)

	projecter := discord.Group("/projecter")
	projecter.Use(middlewares.OpenJwtAuthMiddleware.MiddlewareFunc())
	{
		projecter.POST("/", bindAdminConfig)
		projecter.POST("/activity", activityConfig)
	}

	user := discord.Group("/user")
	user.Use(middlewares.JwtAuthMiddleware.MiddlewareFunc())
	{
		user.POST("/mint", handleCustomMint)
		user.GET("/:user_id", getUser)
		user.POST("/", bindUserAddress)
	}
}

func indexEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, ginutils.DataResponse("Rainbow-App-Service"))
}
