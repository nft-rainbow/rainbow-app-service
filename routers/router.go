package routers

import (
	"net/http"
	"strconv"

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
		projecter.GET("/", getProjecterList)
		projecter.GET("/:id", getProjecter)
		projecter.POST("/activity", activityConfig)
		projecter.GET("/activity", getActivityList)
		projecter.GET("/activity/:id", getActivity)
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