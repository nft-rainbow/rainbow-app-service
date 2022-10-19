package routers

import (
	"github.com/gin-gonic/gin"
	discordbot_errors "github.com/nft-rainbow/discordbot-service/discordbot-errors"
	"github.com/nft-rainbow/discordbot-service/models"
	"github.com/nft-rainbow/discordbot-service/services"
	"github.com/nft-rainbow/discordbot-service/utils/ginutils"
	"strconv"
)

func activityConfig(c *gin.Context) {
	var config *models.ActivityConfig
	if err := c.ShouldBind(&config); err != nil {
		ginutils.RenderRespError(c, err, discordbot_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	err := services.ActivityConfig(config)
	ginutils.RenderResp(c, "success", err)
}

func getActivityList(c *gin.Context) {
	pagination, err := GetPagination(c)
	if err != nil {
		ginutils.RenderRespError(c, discordbot_errors.ERR_INVALID_PAGINATION)
		return
	}
	mints, err := models.FindAndCountActivity(GetIdFromJwtClaim(c), pagination.Offset(), pagination.Limit)
	ginutils.RenderResp(c, mints, err)
}

func getActivity(c *gin.Context) {
	activityId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ginutils.RenderRespError(c, discordbot_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}
	item, err := models.FindBindingActivityConfigById(activityId)
	ginutils.RenderResp(c, item, err)
}

func getProjecterList(c *gin.Context){
	pagination, err := GetPagination(c)
	if err != nil {
		ginutils.RenderRespError(c, discordbot_errors.ERR_INVALID_PAGINATION)
		return
	}
	mints, err := models.FindAndCountAdminConfig(GetIdFromJwtClaim(c), pagination.Offset(), pagination.Limit)
	ginutils.RenderResp(c, mints, err)
}

func getProjecter(c *gin.Context){
	ProjecterId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ginutils.RenderRespError(c, discordbot_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}
	item, err := models.FindBindingConfigById(ProjecterId)
	ginutils.RenderResp(c, item, err)
}