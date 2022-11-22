package routers

import (
	"github.com/gin-gonic/gin"
	appService_errors "github.com/nft-rainbow/rainbow-app-service/appService-errors"
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/services"
	"github.com/nft-rainbow/rainbow-app-service/utils/ginutils"
	"strconv"
)

func bindDiscordCustomProjectorConfig(c *gin.Context) {
	var adminConfig *models.DiscordCustomProjectorConfig
	if err := c.ShouldBind(&adminConfig); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	err := services.BindDiscordProjectorConfig(adminConfig, GetIdFromJwtClaim(c))
	ginutils.RenderResp(c, "success", err)
}

func bindDoDoCustomProjectorConfig(c *gin.Context) {
	var adminConfig *models.DoDoCustomProjectorConfig
	if err := c.ShouldBind(&adminConfig); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	err := services.BindDoDoProjectorConfig(adminConfig, GetIdFromJwtClaim(c))
	ginutils.RenderResp(c, "success", err)
}

func discordCustomActivityConfig(c *gin.Context) {
	var config *models.DiscordCustomActivityConfig
	if err := c.ShouldBind(&config); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	err := services.DiscordCustomActivityConfig(config, GetIdFromJwtClaim(c))
	ginutils.RenderResp(c, "success", err)
}

func dodoCustomActivityConfig(c *gin.Context) {
	var config *models.DoDoCustomActivityConfig
	if err := c.ShouldBind(&config); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	err := services.DoDoCustomActivityConfig(config, GetIdFromJwtClaim(c))
	ginutils.RenderResp(c, "success", err)
}

func getDiscordCustomActivityList(c *gin.Context) {
	pagination, err := GetPagination(c)
	if err != nil {
		ginutils.RenderRespError(c, appService_errors.ERR_INVALID_PAGINATION)
		return
	}
	mints, err := models.FindAndCountDiscordActivity(GetIdFromJwtClaim(c), pagination.Offset(), pagination.Limit)
	ginutils.RenderResp(c, mints, err)
}

func getDoDoCustomActivityList(c *gin.Context) {
	pagination, err := GetPagination(c)
	if err != nil {
		ginutils.RenderRespError(c, appService_errors.ERR_INVALID_PAGINATION)
		return
	}
	mints, err := models.FindAndCountDoDoActivity(GetIdFromJwtClaim(c), pagination.Offset(), pagination.Limit)
	ginutils.RenderResp(c, mints, err)
}

func getDiscordCustomActivity(c *gin.Context) {
	activityId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ginutils.RenderRespError(c, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}
	item, err := models.FindBindingDiscordCustomActivityConfigById(activityId)
	ginutils.RenderResp(c, item, err)
}

func getDoDoCustomActivity(c *gin.Context) {
	activityId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ginutils.RenderRespError(c, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}
	item, err := models.FindBindingDoDoCustomActivityConfigById(activityId)
	ginutils.RenderResp(c, item, err)
}

func getDiscordCustomProjectorList(c *gin.Context){
	pagination, err := GetPagination(c)
	if err != nil {
		ginutils.RenderRespError(c, appService_errors.ERR_INVALID_PAGINATION)
		return
	}
	mints, err := models.FindAndCountDiscordCustomProjectorConfig(GetIdFromJwtClaim(c), pagination.Offset(), pagination.Limit)
	ginutils.RenderResp(c, mints, err)
}

func getDoDoCustomProjectorList(c *gin.Context){
	pagination, err := GetPagination(c)
	if err != nil {
		ginutils.RenderRespError(c, appService_errors.ERR_INVALID_PAGINATION)
		return
	}
	mints, err := models.FindAndCountDoDoCustomProjectorConfig(GetIdFromJwtClaim(c), pagination.Offset(), pagination.Limit)
	ginutils.RenderResp(c, mints, err)
}

func getDiscordCustomProjector(c *gin.Context){
	ProjectorId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ginutils.RenderRespError(c, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}
	item, err := models.FindBindingDiscordConfigById(ProjectorId)
	ginutils.RenderResp(c, item, err)
}

func getDoDoCustomProjector(c *gin.Context){
	ProjectorId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ginutils.RenderRespError(c, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}
	item, err := models.FindBindingDoDoConfigById(ProjectorId)
	ginutils.RenderResp(c, item, err)
}