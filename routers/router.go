package routers

import (
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/nft-rainbow/rainbow-app-service/docs"
	"github.com/nft-rainbow/rainbow-app-service/middlewares"
	"github.com/nft-rainbow/rainbow-app-service/utils/ginutils"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/", indexEndpoint)

	apps := router.Group("/apps")
	apps.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	custom := apps.Group("/custom")
	poap := apps.Group("/poap")

	discord := custom.Group("/discord")
	dodo := custom.Group("/dodo")
	discord.GET("/:guild_id/channels", getDiscordChannelInfo)
	dodo.GET("/:island_id/channels", getDoDoChannelInfo)

	dodoCustomProject := dodo.Group("/projector")
	dodoCustomProject.Use(middlewares.JwtAuthMiddleware.MiddlewareFunc())
	{
		dodoCustomProject.POST("/", setDoDoCustomProjectConfig)
		dodoCustomProject.GET("/", getDoDoCustomProjectList)
		dodoCustomProject.GET("/:id", getDoDoCustomProject)
		dodoCustomProject.POST("/activity", setDoDoCustomActivityConfig)
		dodoCustomProject.GET("/activity", getDoDoCustomActivityList)
		dodoCustomProject.GET("/activity/:id", getDoDoCustomActivity)
	}

	discordCustomProjector := discord.Group("/projector")
	discordCustomProjector.Use(middlewares.JwtAuthMiddleware.MiddlewareFunc())
	{
		discordCustomProjector.POST("/", setDiscordCustomProjectConfig)
		discordCustomProjector.GET("/", getDiscordCustomProjectList)
		discordCustomProjector.GET("/:id", getDiscordCustomProject)
		discordCustomProjector.POST("/activity", setDiscordCustomActivityConfig)
		discordCustomProjector.GET("/activity", getDiscordCustomActivityList)
		discordCustomProjector.GET("/activity/:id", getDiscordCustomActivity)
	}

	poap.POST("/csv", poapMintByCSV)
	poap.POST("/h5", poapMintByH5)
	poap.GET("/activity/:id", getPOAPActivity)
	poap.GET("/activity/result/:activity_id", getPOAPAResultList)
	poap.GET("/activity/result/:activity_id/:id", getPOAPAResult)
	poap.GET("/count/:address/:activity_id", getMintCount)
	poap.POST("/sharer", updateBySharing)
	poap.Use(middlewares.JwtAuthMiddleware.MiddlewareFunc())
	{
		poap.POST("/activity", setPOAPActivityConfig)
		poap.POST("/activity/h5", setPOAPH5Config)
		poap.GET("/activity", getPOAPActivityList)
		poap.POST("/config", setNewYearConfig)
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