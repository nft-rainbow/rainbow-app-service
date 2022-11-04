package routers

import (
	"github.com/gin-gonic/gin"
	appService_errors "github.com/nft-rainbow/rainbow-app-service/appService-errors"
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/services"
	"github.com/nft-rainbow/rainbow-app-service/utils/ginutils"
	"strconv"
)

func discordActivityConfig(c *gin.Context) {
	var config *models.DiscordActivityConfig
	if err := c.ShouldBind(&config); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	err := services.DiscordActivityConfig(config)
	ginutils.RenderResp(c, "success", err)
}

func dodoActivityConfig(c *gin.Context) {
	var config *models.DoDoActivityConfig
	if err := c.ShouldBind(&config); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	err := services.DoDoActivityConfig(config)
	ginutils.RenderResp(c, "success", err)
}

func getDiscordActivityList(c *gin.Context) {
	pagination, err := GetPagination(c)
	if err != nil {
		ginutils.RenderRespError(c, appService_errors.ERR_INVALID_PAGINATION)
		return
	}
	mints, err := models.FindAndCountDiscordActivity(GetIdFromJwtClaim(c), pagination.Offset(), pagination.Limit)
	ginutils.RenderResp(c, mints, err)
}

func getDoDoActivityList(c *gin.Context) {
	pagination, err := GetPagination(c)
	if err != nil {
		ginutils.RenderRespError(c, appService_errors.ERR_INVALID_PAGINATION)
		return
	}
	mints, err := models.FindAndCountDoDoActivity(GetIdFromJwtClaim(c), pagination.Offset(), pagination.Limit)
	ginutils.RenderResp(c, mints, err)
}

func getDiscordActivity(c *gin.Context) {
	activityId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ginutils.RenderRespError(c, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}
	item, err := models.FindBindingDiscordActivityConfigById(activityId)
	ginutils.RenderResp(c, item, err)
}

func getDoDoActivity(c *gin.Context) {
	activityId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ginutils.RenderRespError(c, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}
	item, err := models.FindBindingDoDoActivityConfigById(activityId)
	ginutils.RenderResp(c, item, err)
}

func getDiscordProjectorList(c *gin.Context){
	pagination, err := GetPagination(c)
	if err != nil {
		ginutils.RenderRespError(c, appService_errors.ERR_INVALID_PAGINATION)
		return
	}
	mints, err := models.FindAndCountDiscordAdminConfig(GetIdFromJwtClaim(c), pagination.Offset(), pagination.Limit)
	ginutils.RenderResp(c, mints, err)
}

func getDoDoProjectorList(c *gin.Context){
	pagination, err := GetPagination(c)
	if err != nil {
		ginutils.RenderRespError(c, appService_errors.ERR_INVALID_PAGINATION)
		return
	}
	mints, err := models.FindAndCountDoDoAdminConfig(GetIdFromJwtClaim(c), pagination.Offset(), pagination.Limit)
	ginutils.RenderResp(c, mints, err)
}


func getDiscordProjector(c *gin.Context){
	ProjectorId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ginutils.RenderRespError(c, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}
	item, err := models.FindBindingDiscordConfigById(ProjectorId)
	ginutils.RenderResp(c, item, err)
}

func getDoDoProjector(c *gin.Context){
	ProjectorId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ginutils.RenderRespError(c, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}
	item, err := models.FindBindingDoDoConfigById(ProjectorId)
	ginutils.RenderResp(c, item, err)
}