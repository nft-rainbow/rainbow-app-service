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
	router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	custom := router.Group("/custom")
	poap := router.Group("/poap")

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

	poap.Use(middlewares.JwtAuthMiddleware.MiddlewareFunc())
	{
		poap.POST("/activity", setPOAPActivityConfig)
		poap.POST("/activity/h5", setPOAPH5Config)
		poap.POST("/csv", poapMintByCSV)
		poap.POST("/h5", poapMintByH5)
		poap.GET("/activity", getPOAPActivityList)
		poap.GET("/activity/:id", getPOAPActivity)
		poap.GET("/activity/result/:activity_id", getPOAPAResultList)
		poap.GET("/activity/result/:activity_id/:id", getPOAPAResult)
	}

	poapNewYear := poap.Group("/newYear")
	poapNewYear.Use(middlewares.JwtAuthMiddleware.MiddlewareFunc())
	{
		poapNewYear.POST("/config", setNewYearConfig)
		poapNewYear.POST("/common", newYearCommonMint)
		poapNewYear.POST("/special/:common_id", newYearSpecialMint)
		poapNewYear.POST("/sharer", updateBySharing)
		poapNewYear.GET("/count/common/:address/:id", getCommonMintCount)
		poapNewYear.GET("/count/special/:address/:id", getSpecialMintCount)
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