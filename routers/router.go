package routers

import (
	"net/http"
	"strconv"

	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	_ "github.com/nft-rainbow/rainbow-app-service/docs"
	"github.com/nft-rainbow/rainbow-app-service/middlewares"
	"github.com/nft-rainbow/rainbow-app-service/utils/ginutils"
)

var (
	botActivityHandler *BotActivityRouteHandler
)

func Init() {
	var err error
	botActivityHandler, err = NewBotActivityRouteHandler()
	if err != nil {
		panic(err)
	}
}

func SetupRoutes(router *gin.Engine) {
	router.GET("/", indexEndpoint)

	apps := router.Group("/apps")
	apps.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	bot := apps.Group("/bot/:social_tool")

	// social := bot.Group("/:social_tool")
	botManager := bot.Group("/manager")
	botManager.Use(middlewares.JwtAuthMiddleware.MiddlewareFunc())
	{
		// socialManager.GET("/:island_id/channels", getDoDoChannelInfo)
		botManager.POST("/authcode", botActivityHandler.verifyUser)
		botManager.POST("/", botActivityHandler.insertProjectManager)
		botManager.GET("/", botActivityHandler.getProjectManager)
	}

	botActivity := bot.Group("/activity")
	botActivity.Use(middlewares.JwtAuthMiddleware.MiddlewareFunc())
	{
		// dodoProjector.GET("/", getProjectorList)
		// dodoProjector.GET("/", getDoDoCustomProjectList)
		// dodoProjector.GET("/:id", getProjector)
		botActivity.POST("", setDoDoCustomActivityConfig)
		botActivity.GET("", getDoDoCustomActivityList)
		botActivity.GET("/:id", getDoDoCustomActivity)
	}

	// discord := bot.Group("/discord")
	// discord.GET("/:guild_id/channels", getDiscordChannelInfo)
	// discordCustomProjector := discord.Group("/projector")
	// discordCustomProjector.Use(middlewares.JwtAuthMiddleware.MiddlewareFunc())
	// {
	// 	discordCustomProjector.POST("/", setDiscordCustomProjectConfig)
	// 	discordCustomProjector.GET("/", getDiscordCustomProjectList)
	// 	discordCustomProjector.GET("/:id", getDiscordCustomProject)
	// 	discordCustomProjector.POST("/activity", setDiscordCustomActivityConfig)
	// 	discordCustomProjector.GET("/activity", getDiscordCustomActivityList)
	// 	discordCustomProjector.GET("/activity/:id", getDiscordCustomActivity)
	// }

	poap := apps.Group("/poap")
	poap.POST("/h5", middlewares.IpLimitMiddleware(), poapMintByH5)
	poap.GET("/activity/:activity_id", getPOAPActivity)
	poap.GET("/activity/result/:activity_id", getPOAPResultList)
	poap.GET("/activity/result/:activity_id/:id", getPOAPResultDetail)
	poap.GET("/count/:address/:activity_id", getMintCount)
	poap.POST("/anyweb/code", collectAnywebUserCode)
	poap.Use(middlewares.JwtAuthMiddleware.MiddlewareFunc())
	{
		poap.POST("/activity/push", pushActivity)
		poap.POST("/activity/server", bindServerInfo)
		poap.GET("/activity/push/:bot", getPushes)
		poap.GET("/activity/servers/:bot", getServers)
		poap.GET("/activity/channels/discord/:guild_id", getDiscordChannels)
		poap.GET("/activity/channels/dodo/:island_id", getDoDoChannels)
		poap.GET("/activity/roles/discord/:guild_id", getDiscordRoles)
		poap.GET("/activity/roles/dodo/:island_id", getDoDoRoles)
		poap.POST("/csv", poapMintByCSV)
		poap.POST("/activity", setPOAPActivityConfig)
		poap.PUT("/activity/:activity_id", updatePOAPConfig)
		poap.POST("/activity/h5", setPOAPH5Config)
		poap.GET("/activity", getPOAPActivityList)
	}
}

func indexEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, ginutils.DataResponse("Rainbow-App-Service"))
}

type Pagination struct {
	Page  int `json:"page" form:"page"`
	Limit int `json:"limit" form:"limit"`
}

func (p Pagination) Offset() int {
	return (p.Page - 1) * p.Limit
}

func GetPagination(c *gin.Context) (*Pagination, error) {
	var pagination Pagination
	var err error
	pageStr := c.DefaultQuery("page", "1")
	sizeStr := c.DefaultQuery("limit", "10")
	if pagination.Page, err = strconv.Atoi(pageStr); err != nil {
		return nil, err
	}
	if pagination.Page < 1 {
		pagination.Page = 1
	}

	if pagination.Limit, err = strconv.Atoi(sizeStr); err != nil {
		return nil, err
	}
	if pagination.Limit < 1 {
		pagination.Limit = 10
	}
	return &pagination, nil
}

func GetIdFromJwtClaim(c *gin.Context) uint {
	return c.GetUint(middlewares.JwtIdentityKey)
}
