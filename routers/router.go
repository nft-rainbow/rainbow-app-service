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
	custom := router.Group("/custom")
	poap := router.Group("/poap")

	discord := custom.Group("/discord")
	dodo := custom.Group("/dodo")
	discord.GET("/:guild_id/channels", getDiscordChannelInfo)
	dodo.GET("/:island_id/channels", getDoDoChannelInfo)

	dodoCustomProject := dodo.Group("/projector")
	dodoCustomProject.Use(middlewares.JwtAuthMiddleware.MiddlewareFunc())
	{
		dodoCustomProject.POST("/", bindDoDoCustomProjectorConfig)
		dodoCustomProject.GET("/", getDoDoCustomProjectorList)
		dodoCustomProject.GET("/:id", getDoDoCustomProjector)
		dodoCustomProject.POST("/activity", dodoCustomActivityConfig)
		dodoCustomProject.GET("/activity", getDoDoCustomActivityList)
		dodoCustomProject.GET("/activity/:id", getDoDoCustomActivity)
	}

	discordCustomProjector := discord.Group("/projector")
	discordCustomProjector.Use(middlewares.JwtAuthMiddleware.MiddlewareFunc())
	{
		discordCustomProjector.POST("/", bindDiscordCustomProjectorConfig)
		discordCustomProjector.GET("/", getDiscordCustomProjectorList)
		discordCustomProjector.GET("/:id", getDiscordCustomProjector)
		discordCustomProjector.POST("/activity", discordCustomActivityConfig)
		discordCustomProjector.GET("/activity", getDiscordCustomActivityList)
		discordCustomProjector.GET("/activity/:id", getDiscordCustomActivity)
	}

	poapProjector := poap.Group("/projector")
	poapProjector.Use(middlewares.JwtAuthMiddleware.MiddlewareFunc())
	{
		poapProjector.POST("/activity", POAPActivityConfig)
		poapProjector.POST("/mint/csv", POAPCSVMint)
		poapProjector.POST("/mint/h5", POAPH5Mint)
		poapProjector.POST("/", POAPProjectorConfig)
		poapProjector.GET("/activity", getPOAPActivityList)
		poapProjector.GET("/activity/:id", getPOAPActivity)
		poapProjector.GET("/", getPOAPProjectorList)
		poapProjector.GET("/:id", getPOAPProjector)
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