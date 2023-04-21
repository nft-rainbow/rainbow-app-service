package routers

import (
	"net/http"
	"strconv"

	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	_ "github.com/nft-rainbow/rainbow-app-service/docs"
	"github.com/nft-rainbow/rainbow-app-service/middlewares"
	"github.com/nft-rainbow/rainbow-app-service/services"
	"github.com/nft-rainbow/rainbow-app-service/utils/ginutils"
)

var (
	botServerController *BotServerController
	activityService     *services.ActivityService
	walletService       *services.WalletService
)

func Init() {
	var err error
	botServerController, err = NewBotServerController()
	if err != nil {
		panic(err)
	}
	walletService = services.NewWalletService()
	activityService = services.GetActivityService()
}

func SetupRoutes(router *gin.Engine) {
	router.GET("/", indexEndpoint)

	apps := router.Group("/apps")
	apps.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	bot := apps.Group("/bot")
	bot.GET("/invite_url", botServerController.GetInviteUrl)

	botServer := bot.Group("/server")
	botServer.Use(middlewares.JwtAuthMiddleware.MiddlewareFunc())
	{
		botServer.GET("/channels", botServerController.GetChannels)
		botServer.GET("/roles", botServerController.GetRoles)

		botServer.GET("/authcode", botServerController.getAuthCode)
		botServer.POST("", botServerController.insertBotServer)
		botServer.GET("", botServerController.GetBotServers)
		botServer.GET("/:id", botServerController.GetBotServer)

		botServer.POST("/:id/pushinfo", botServerController.AddPushInfo)
		botServer.PUT("/pushinfo/:id", botServerController.UpdatePushInfo)
		botServer.POST("/push/:id", botServerController.Push)

		botServer.GET("/activities", botServerController.GetActivitiesOfUserBotServers)
	}

	// botActivity := bot.Group("/activity")
	// botActivity.Use(middlewares.JwtAuthMiddleware.MiddlewareFunc())
	// {
	// dodoProjector.GET("/", getProjectorList)
	// dodoProjector.GET("/", getDoDoCustomProjectList)
	// dodoProjector.GET("/:id", getProjector)
	// botActivity.POST("", setDoDoCustomActivityConfig)
	// botActivity.GET("", getDoDoCustomActivityList)
	// botActivity.GET("/:id", getDoDoCustomActivity)
	// }

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
	poap.GET("/activity/:activity_code", getActivity)
	poap.GET("/activity/result/:activity_code", getMintResultList)
	poap.GET("/activity/result/:activity_code/:id", getMintResultDetail)
	poap.GET("/count/:address/:activity_code", getMintCount)
	poap.POST("/wallet/user", addWalletUser)
	poap.Use(middlewares.JwtAuthMiddleware.MiddlewareFunc())
	{
		// poap.POST("/activity/push", pushActivity)
		// poap.POST("/activity/server", bindServerInfo)
		// poap.GET("/activity/push/:bot", getPushes)
		// poap.GET("/activity/servers/:bot", getServers)
		// poap.GET("/activity/channels/discord/:guild_id", getDiscordChannels)
		// poap.GET("/activity/channels/dodo/:island_id", getDoDoChannels)
		// poap.GET("/activity/roles/discord/:guild_id", getDiscordRoles)
		// poap.GET("/activity/roles/dodo/:island_id", getDoDoRoles)
		// poap.POST("/csv", poapMintByCSV)

		poap.POST("/activity", addActivity)
		poap.PUT("/activity/:activity_code", updateActivity)
		poap.POST("/activity/h5", setActivityH5Config)
		poap.GET("/activity", getUserActivities)
	}
}

func indexEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, ginutils.DataResponse("Rainbow-App-Service"))
}

type Pagination struct {
	Page  int `json:"page" form:"page" default:"1"`
	Limit int `json:"limit" form:"limit" default:"10"`
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
