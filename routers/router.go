package routers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nft-rainbow/rainbow-app-service/middlewares"
	"github.com/nft-rainbow/rainbow-app-service/utils/ginutils"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/", indexEndpoint)
	discord := router.Group("/discord")
	dodo := router.Group("/dodo")

	discord.GET("/:guild_id/channels", getDiscordChannelInfo)
	dodo.GET("/:island_id/channels", getDoDoChannelInfo)

	dodoProject := dodo.Group("/projector")
	dodoProject.Use(middlewares.OpenJwtAuthMiddleware.MiddlewareFunc())
	{
		dodoProject.POST("/", bindDoDoProjectorConfig)
		dodoProject.GET("/", getDoDoProjectorList)
		dodoProject.GET("/:id", getDoDoProjector)
		dodoProject.POST("/activity", dodoActivityConfig)
		dodoProject.GET("/activity", getDoDoActivityList)
		dodoProject.GET("/activity/:id", getDoDoActivity)
	}

	discordProjector := discord.Group("/projector")
	discordProjector.Use(middlewares.OpenJwtAuthMiddleware.MiddlewareFunc())
	{
		discordProjector.POST("/", bindDiscordProjectorConfig)
		discordProjector.GET("/", getDiscordProjectorList)
		discordProjector.GET("/:id", getDiscordProjector)
		discordProjector.POST("/activity", discordActivityConfig)
		discordProjector.GET("/activity", getDiscordActivityList)
		discordProjector.GET("/activity/:id", getDiscordActivity)
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
	return c.GetUint(middlewares.OpenJwtIdentityKey)
}