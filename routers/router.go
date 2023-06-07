package routers

import (
	"net/http"

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
	poap.POST("/h5", middlewares.IpLimitMiddleware(), mintByH5)
	poap.GET("/activity/:activity_code", getActivity)
	poap.GET("/activity/result/:activity_code", getMintResultList)
	poap.GET("/activity/result/:activity_code/:id", getMintResultDetail)
	poap.GET("/count/:address/:activity_code", getMintCount)
	poap.POST("/wallet/user", addWalletUser)
	poap.Use(middlewares.JwtAuthMiddleware.MiddlewareFunc())
	{
		poap.POST("/activity", addActivity)
		poap.POST("/activity/:activity_code/nftconfigs", addActivityNftConfigs)
		poap.PUT("/activity/:activity_code/base", updateActivityBase)
		poap.PUT("/activity/nftconfig/:nft_config_id", updateActivityNftConfig)
		poap.DELETE("/activity/nftconfig/:nft_config_id", deleteActivityNftConfig)
		poap.POST("/activity/h5", setActivityH5Config)
		poap.GET("/activity", getUserActivities)

		poap.POST("/activity/token-reserve", addTokenReserves)
		poap.GET("/activity/token-reserve/:activity_code", getTokenReserves)
	}

	certi := apps.Group("/certis")
	{
		certiCtrl := NewCertiController()
		certi.POST("/strategy/type/:certificate_type", certiCtrl.InsertCertificateStrategy)

		certi.GET("/strategy/:id/certificates", certiCtrl.GetCertificates)
		certi.POST("/strategy/:id/certificates", certiCtrl.InsertCertificates)
		certi.DELETE("/strategy/:id/certificates", certiCtrl.DeleteCertificates)

		certi.GET("contract_certificate/:certificate_id/snapshot", certiCtrl.GetSnapshots)
		certi.POST("contract_certificate/:certificate_id/snapshot/run", certiCtrl.TriggerObtainSnapshot)
	}
}

func indexEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, ginutils.DataResponse("Rainbow-App-Service"))
}

// type Pagination struct {
// 	Page  int `json:"page" form:"page" default:"1"`
// 	Limit int `json:"limit" form:"limit" default:"10"`
// }

// func (p Pagination) Offset() int {
// 	return (p.Page - 1) * p.Limit
// }

// func GetPagination(c *gin.Context) (*Pagination, error) {
// 	var pagination Pagination
// 	var err error
// 	pageStr := c.DefaultQuery("page", "1")
// 	sizeStr := c.DefaultQuery("limit", "10")
// 	if pagination.Page, err = strconv.Atoi(pageStr); err != nil {
// 		return nil, err
// 	}
// 	if pagination.Page < 1 {
// 		pagination.Page = 1
// 	}

// 	if pagination.Limit, err = strconv.Atoi(sizeStr); err != nil {
// 		return nil, err
// 	}
// 	if pagination.Limit < 1 {
// 		pagination.Limit = 10
// 	}
// 	return &pagination, nil
// }

func GetIdFromJwtClaim(c *gin.Context) uint {
	return c.GetUint(middlewares.JwtIdentityKey)
}
