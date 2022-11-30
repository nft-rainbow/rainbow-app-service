package routers

import (
	"github.com/gin-gonic/gin"
	appService_errors "github.com/nft-rainbow/rainbow-app-service/appService-errors"
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/services"
	"github.com/nft-rainbow/rainbow-app-service/utils/ginutils"
	"strconv"
)

// @Tags        CustomMint
// @ID          SetDiscordCustomProjectConfig
// @Summary     Set Discord custom project
// @Description Set Discord custom project config
// @security    ApiKeyAuth
// @Produce     json
// @Param       Authorization header   string true "Bearer JWT"
// @Param       discord_custom_project_dto body     models.DiscordCustomProjectConfig true "discord_custom_project_dto"
// @Success     200           {object} string "success"
// @Failure     400           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Invalid request"
// @Failure     500           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Internal Server error"
// @Router      /custom/discord/projector/ [post]
func setDiscordCustomProjectConfig(c *gin.Context) {
	var adminConfig *models.DiscordCustomProjectConfig
	if err := c.ShouldBind(&adminConfig); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	err := services.BindDiscordProjectConfig(adminConfig, GetIdFromJwtClaim(c))
	ginutils.RenderResp(c, "success", err)
}

// @Tags        CustomMint
// @ID          SetDoDoCustomProjectConfig
// @Summary     Set DoDo custom project
// @Description Set DoDo custom project config
// @security    ApiKeyAuth
// @Produce     json
// @Param       Authorization header   string true "Bearer JWT"
// @Param       discord_custom_project_dto body     models.DoDoCustomProjectConfig true "discord_custom_project_dto"
// @Success     200           {object} string "success"
// @Failure     400           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Invalid request"
// @Failure     500           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Internal Server error"
// @Router      /custom/dodo/projector/ [post]
func setDoDoCustomProjectConfig(c *gin.Context) {
	var adminConfig *models.DoDoCustomProjectConfig
	if err := c.ShouldBind(&adminConfig); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	err := services.BindDoDoProjectConfig(adminConfig, GetIdFromJwtClaim(c))
	ginutils.RenderResp(c, "success", err)
}

// @Tags        CustomMint
// @ID          SetDiscordCustomActivityConfig
// @Summary     Set Discord custom Activity
// @Description Set Discord custom Activity config
// @security    ApiKeyAuth
// @Produce     json
// @Param       Authorization header   string true "Bearer JWT"
// @Param       discord_custom_activity_dto body     models.DiscordCustomActivityConfig true "discord_custom_activity_dto"
// @Success     200           {object} string "success"
// @Failure     400           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Invalid request"
// @Failure     500           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Internal Server error"
// @Router      /custom/discord/projector/activity [post]
func setDiscordCustomActivityConfig(c *gin.Context) {
	var config *models.DiscordCustomActivityConfig
	if err := c.ShouldBind(&config); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	err := services.DiscordCustomActivityConfig(config, GetIdFromJwtClaim(c))
	ginutils.RenderResp(c, "success", err)
}

// @Tags        CustomMint
// @ID          SetDoDoCustomActivityConfig
// @Summary     Set DoDo custom Activity
// @Description Set DoDo custom Activity config
// @security    ApiKeyAuth
// @Produce     json
// @Param       Authorization header   string true "Bearer JWT"
// @Param       dodo_custom_activity_dto body     models.DoDoCustomActivityConfig true "dodo_custom_activity_dto"
// @Success     200           {object} string "success"
// @Failure     400           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Invalid request"
// @Failure     500           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Internal Server error"
// @Router      /custom/dodo/projector/activity [post]
func setDoDoCustomActivityConfig(c *gin.Context) {
	var config *models.DoDoCustomActivityConfig
	if err := c.ShouldBind(&config); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	err := services.DoDoCustomActivityConfig(config, GetIdFromJwtClaim(c))
	ginutils.RenderResp(c, "success", err)
}

// @Tags        CustomMint
// @ID          GetDiscordCustomActivityList
// @Summary     Get Discord custom Activity list
// @Description Get Discord custom Activity list
// @security    ApiKeyAuth
// @Produce     json
// @Param       Authorization header   string true "Bearer JWT"
// @Param       page          query    integer false "page"
// @Param       limit         query    integer false "limit"
// @Success     200           {object} models.DiscordActivityQueryResult
// @Failure     400           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Invalid request"
// @Failure     500           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Internal Server error"
// @Router      /custom/discord/projector/activity [get]
func getDiscordCustomActivityList(c *gin.Context) {
	pagination, err := GetPagination(c)
	if err != nil {
		ginutils.RenderRespError(c, appService_errors.ERR_INVALID_PAGINATION)
		return
	}
	mints, err := models.FindAndCountDiscordActivity(GetIdFromJwtClaim(c), pagination.Offset(), pagination.Limit)
	ginutils.RenderResp(c, mints, err)
}

// @Tags        CustomMint
// @ID          GetDoDoCustomActivityList
// @Summary     Get DoDo custom Activity list
// @Description Get DoDo custom Activity list
// @security    ApiKeyAuth
// @Produce     json
// @Param       Authorization header   string true "Bearer JWT"
// @Param       page          query    integer false "page"
// @Param       limit         query    integer false "limit"
// @Success     200           {object} models.DoDoActivityQueryResult
// @Failure     400           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Invalid request"
// @Failure     500           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Internal Server error"
// @Router      /custom/dodo/projector/activity [get]
func getDoDoCustomActivityList(c *gin.Context) {
	pagination, err := GetPagination(c)
	if err != nil {
		ginutils.RenderRespError(c, appService_errors.ERR_INVALID_PAGINATION)
		return
	}
	mints, err := models.FindAndCountDoDoActivity(GetIdFromJwtClaim(c), pagination.Offset(), pagination.Limit)
	ginutils.RenderResp(c, mints, err)
}

// @Tags        CustomMint
// @ID          GetDiscordCustomActivityDetail
// @Summary     Get Discord custom Activity detail
// @Description Get Discord custom Activity detail info
// @security    ApiKeyAuth
// @Produce     json
// @Success     200           {object} models.DiscordCustomActivityConfig
// @Failure     400           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Invalid request"
// @Failure     500           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Internal Server error"
// @Router      /custom/discord/activity/{id} [get]
func getDiscordCustomActivity(c *gin.Context) {
	activityId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ginutils.RenderRespError(c, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}
	item, err := models.FindDiscordCustomActivityConfigById(activityId)
	ginutils.RenderResp(c, item, err)
}

// @Tags        CustomMint
// @ID          GetDoDoCustomActivityDetail
// @Summary     Get DoDo custom Activity detail
// @Description Get DoDo custom Activity detail info
// @security    ApiKeyAuth
// @Produce     json
// @Success     200           {object} models.DoDoCustomActivityConfig
// @Failure     400           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Invalid request"
// @Failure     500           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Internal Server error"
// @Router      /custom/dodo/activity/{id} [get]
func getDoDoCustomActivity(c *gin.Context) {
	activityId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ginutils.RenderRespError(c, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}
	item, err := models.FindDoDoCustomActivityConfigById(activityId)
	ginutils.RenderResp(c, item, err)
}

// @Tags        CustomMint
// @ID          GetDiscordCustomProjectList
// @Summary     Get Discord custom project list
// @Description Get Discord custom project list
// @security    ApiKeyAuth
// @Produce     json
// @Param       Authorization header   string true "Bearer JWT"
// @Param       page          query    integer false "page"
// @Param       limit         query    integer false "limit"
// @Success     200           {object} models.DiscordCustomProjectConfigQueryResult
// @Failure     400           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Invalid request"
// @Failure     500           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Internal Server error"
// @Router      /custom/discord/projector/ [get]
func getDiscordCustomProjectList(c *gin.Context){
	pagination, err := GetPagination(c)
	if err != nil {
		ginutils.RenderRespError(c, appService_errors.ERR_INVALID_PAGINATION)
		return
	}
	mints, err := models.FindAndCountDiscordCustomProjectConfig(GetIdFromJwtClaim(c), pagination.Offset(), pagination.Limit)
	ginutils.RenderResp(c, mints, err)
}

// @Tags        CustomMint
// @ID          GetDoDoCustomProjectList
// @Summary     Get DoDo custom project list
// @Description Get DoDo custom project list
// @security    ApiKeyAuth
// @Produce     json
// @Param       Authorization header   string true "Bearer JWT"
// @Param       page          query    integer false "page"
// @Param       limit         query    integer false "limit"
// @Success     200           {object} models.DoDoCustomProjectConfigQueryResult
// @Failure     400           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Invalid request"
// @Failure     500           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Internal Server error"
// @Router      /custom/dodo/projector/ [get]
func getDoDoCustomProjectList(c *gin.Context){
	pagination, err := GetPagination(c)
	if err != nil {
		ginutils.RenderRespError(c, appService_errors.ERR_INVALID_PAGINATION)
		return
	}
	mints, err := models.FindAndCountDoDoCustomProjectConfig(GetIdFromJwtClaim(c), pagination.Offset(), pagination.Limit)
	ginutils.RenderResp(c, mints, err)
}

// @Tags        CustomMint
// @ID          GetDiscordCustomProjectDetail
// @Summary     Get Discord custom project detail
// @Description Get Discord custom project detail info
// @security    ApiKeyAuth
// @Produce     json
// @Param       Authorization header   string true "Bearer JWT"
// @Param       id            path     int    true "id"
// @Success     200           {object} models.DiscordCustomProjectConfig
// @Failure     400           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Invalid request"
// @Failure     500           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Internal Server error"
// @Router      /custom/discord/projector/{id} [get]
func getDiscordCustomProject(c *gin.Context){
	ProjectorId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ginutils.RenderRespError(c, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}
	item, err := models.FindDiscordConfigById(ProjectorId)
	ginutils.RenderResp(c, item, err)
}

// @Tags        CustomMint
// @ID          GetDoDoCustomProjectDetail
// @Summary     Get DoDo custom project detail
// @Description Get DoDo custom project detail info
// @security    ApiKeyAuth
// @Produce     json
// @Param       Authorization header   string true "Bearer JWT"
// @Param       id            path     int    true "id"
// @Success     200           {object} models.DoDoCustomProjectConfig
// @Failure     400           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Invalid request"
// @Failure     500           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Internal Server error"
// @Router      /custom/discord/projector/{id} [get]
func getDoDoCustomProject(c *gin.Context){
	ProjectorId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ginutils.RenderRespError(c, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}
	item, err := models.FindDoDoConfigById(ProjectorId)
	ginutils.RenderResp(c, item, err)
}