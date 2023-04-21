package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mcuadros/go-defaults"
	appService_errors "github.com/nft-rainbow/rainbow-app-service/appService-errors"
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/models/enums"
	"github.com/nft-rainbow/rainbow-app-service/services"
	"github.com/nft-rainbow/rainbow-app-service/utils/ginutils"
)

type BotServerController struct {
	botService *services.BotServerService
}

type (
	IdUintReq struct {
		Id uint `uri:"id" form:"id" binding:"required"`
	}

	ServerIdReq struct {
		ServerId string `uri:"server_id" form:"server_id" binding:"required"`
	}

	SocialAndServerIdReq struct {
		ServerIdReq
		services.SocialToolQueryReq
	}
)

func NewBotServerController() (*BotServerController, error) {
	service, err := services.NewBotServerService()
	if err != nil {
		return nil, err
	}
	return &BotServerController{
		botService: service,
	}, nil
}

// @Tags        Bot
// @ID          GetAuthCode
// @Summary     get bot server roles
// @Description get bot server roles
// @security    ApiKeyAuth
// @Produce     json
// @Param       get_authcode_options query    services.VerifyBotServerReq true "get authcode options"
// @Param       social_tool          query    string                      true "social tool" Enums(dodo,discord)
// @Success     200                  {object} ginutils.CommonMessage
// @Failure     400                  {object} appService_errors.RainbowAppServiceErrorDetailInfo "Invalid request"
// @Failure     500                  {object} appService_errors.RainbowAppServiceErrorDetailInfo "Internal Server error"
// @Router      /bot/server/authcode [get]
func (b *BotServerController) getAuthCode(c *gin.Context) {
	var verifyUserReq services.VerifyBotServerReq
	if err := c.ShouldBindQuery(&verifyUserReq); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	socialTool, err := enums.ParseSocialToolType(verifyUserReq.SocialTool)
	if err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	err = b.botService.GetAuthcode(*socialTool, verifyUserReq.ServerId)
	ginutils.RenderResp(c, ginutils.CommonSuccessMessage, err)
}

// @Tags        Bot
// @ID          InsertBotServer
// @Summary     insert bot server
// @Description insert bot server
// @security    ApiKeyAuth
// @Produce     json
// @Param       insert_bot_server_request body     services.InsertBotServerReq true "insert_bot_server_request"
// @Success     200                       {object} models.BotServer
// @Failure     400                       {object} appService_errors.RainbowAppServiceErrorDetailInfo "Invalid request"
// @Failure     500                       {object} appService_errors.RainbowAppServiceErrorDetailInfo "Internal Server error"
// @Router      /bot/server [post]
func (b *BotServerController) insertBotServer(c *gin.Context) {
	var req services.InsertBotServerReq
	if err := c.ShouldBindJSON(&req); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	userId := GetIdFromJwtClaim(c)
	botServer, err := b.botService.InsertBotServer(userId, req)
	ginutils.RenderResp(c, botServer, err)
}

// @Tags        Bot
// @ID          GetBotServers
// @Summary     get bot servers
// @Description get bot servers
// @security    ApiKeyAuth
// @Produce     json
// @Param       get_bot_servers_params query    services.GetBotServersReq true "get_bot_servers_params"
// @Param       social_tool            query    string                    true "social tool" Enums(dodo,discord)
// @Success     200                    {object} models.FindBotServersResult
// @Failure     400                    {object} appService_errors.RainbowAppServiceErrorDetailInfo "Invalid request"
// @Failure     500                    {object} appService_errors.RainbowAppServiceErrorDetailInfo "Internal Server error"
// @Router      /bot/server [get]
func (b *BotServerController) GetBotServers(c *gin.Context) {
	var queryParams services.GetBotServersReq
	if err := c.ShouldBindQuery(&queryParams); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	userId := GetIdFromJwtClaim(c)
	p, err := b.botService.GetBotServers(userId, &queryParams)
	ginutils.RenderResp(c, p, err)
}

// @Tags        Bot
// @ID          GetBotServer
// @Summary     get bot server
// @Description get bot server
// @security    ApiKeyAuth
// @Produce     json
// @Param       get_bot_server_param path     uint true "get_bot_server_param"
// @Success     200                  {object} models.BotServer
// @Failure     400                  {object} appService_errors.RainbowAppServiceErrorDetailInfo "Invalid request"
// @Failure     500                  {object} appService_errors.RainbowAppServiceErrorDetailInfo "Internal Server error"
// @Router      /bot/server/{id} [get]
func (b *BotServerController) GetBotServer(c *gin.Context) {
	userId := GetIdFromJwtClaim(c)

	var req IdUintReq
	if err := c.ShouldBindUri(&req); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	p, err := b.botService.GetBotServer(userId, req.Id)
	ginutils.RenderResp(c, p, err)
}

// @Tags        Bot
// @ID          AddPushInfo
// @Summary     add push info
// @Description add push info
// @security    ApiKeyAuth
// @Produce     json
// @Param       push_info_req body     services.PushInfoReq true "add push info request"
// @Param       id            path     uint                 true "bot_server ID"
// @Success     200           {object} models.PushInfo
// @Failure     400           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Invalid request"
// @Failure     500           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Internal Server error"
// @Router      /bot/server/{id}/pushinfo [post]
func (b *BotServerController) AddPushInfo(c *gin.Context) {
	var req services.PushInfoReq
	if err := c.ShouldBindJSON(&req); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}
	var uriParams IdUintReq
	if err := c.ShouldBindUri(&uriParams); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	userId := GetIdFromJwtClaim(c)
	res, err := b.botService.AddPushInfo(userId, uriParams.Id, req)
	ginutils.RenderResp(c, res, err)
}

// @Tags        Bot
// @ID          Push
// @Summary     push notification to social tool server
// @Description push notification to social tool server
// @security    ApiKeyAuth
// @Produce     json
// @Param       id  path     uint true "push_info ID"
// @Success     200 {object} ginutils.CommonMessage
// @Failure     400 {object} appService_errors.RainbowAppServiceErrorDetailInfo "Invalid request"
// @Failure     500 {object} appService_errors.RainbowAppServiceErrorDetailInfo "Internal Server error"
// @Router      /bot/server/push/{id} [post]
func (b *BotServerController) Push(c *gin.Context) {
	var uriParams IdUintReq
	if err := c.ShouldBindUri(&uriParams); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	userId := GetIdFromJwtClaim(c)
	err := b.botService.Push(userId, uriParams.Id)
	ginutils.RenderResp(c, ginutils.CommonSuccessMessage, err)
}

// @Tags        Bot
// @ID          UpdatePushInfo
// @Summary     update push info
// @Description update push info
// @security    ApiKeyAuth
// @Produce     json
// @Param       push_info_req body     services.PushInfoReq true "update push info request"
// @Param       id            path     uint                 true "push_info ID"
// @Success     200           {object} models.PushInfo
// @Failure     400           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Invalid request"
// @Failure     500           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Internal Server error"
// @Router      /bot/server/pushinfo/{id} [put]
func (b *BotServerController) UpdatePushInfo(c *gin.Context) {
	var req services.PushInfoReq
	if err := c.ShouldBindJSON(&req); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}
	var uriParams IdUintReq
	if err := c.ShouldBindUri(&uriParams); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	userId := GetIdFromJwtClaim(c)
	res, err := b.botService.UpdatePushInfo(userId, uriParams.Id, req)
	ginutils.RenderResp(c, res, err)

}

// @Tags        Bot
// @ID          GetActivitiesOfUserBotServers
// @Summary     get activites of user bot servers
// @Description get activites of user bot servers
// @security    ApiKeyAuth
// @Produce     json
// @Param       find_bot_server_activities_req query    models.FindBotServerActivitiesCond true "find bot server activities params"
// @Param       social_tool                    query    string                             true "social tool" Enums(dodo,discord)
// @Success     200                            {object} models.FindBotServerActivitiesResult
// @Failure     400                            {object} appService_errors.RainbowAppServiceErrorDetailInfo "Invalid request"
// @Failure     500                            {object} appService_errors.RainbowAppServiceErrorDetailInfo "Internal Server error"
// @Router      /bot/server/activities [get]
func (b *BotServerController) GetActivitiesOfUserBotServers(c *gin.Context) {
	var req models.FindBotServerActivitiesCond
	if err := c.ShouldBindQuery(&req); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}
	defaults.SetDefaults(&req)

	userId := GetIdFromJwtClaim(c)
	result, err := b.botService.GetActivitiesOfBotServers(userId, &req)
	ginutils.RenderResp(c, result, err)

}

// @Tags        Bot
// @ID          GetBotServerChannels
// @Summary     get bot server channels
// @Description get bot server channels
// @security    ApiKeyAuth
// @Produce     json
// @Param       get_channels_options query    SocialAndServerIdReq true "get channels options"
// @Param       social_tool          query    string               true "social tool" Enums(dodo,discord)
// @Success     200                  {array}  services.Channel
// @Failure     400                  {object} appService_errors.RainbowAppServiceErrorDetailInfo "Invalid request"
// @Failure     500                  {object} appService_errors.RainbowAppServiceErrorDetailInfo "Internal Server error"
// @Router      /bot/server/channels [get]
func (b *BotServerController) GetChannels(c *gin.Context) {
	var req SocialAndServerIdReq
	if err := c.ShouldBindQuery(&req); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	socialTool, err := enums.ParseSocialToolType(req.SocialTool)
	if err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	channels, err := b.botService.GetChannels(*socialTool, req.ServerId)
	ginutils.RenderResp(c, channels, err)
}

// @Tags        Bot
// @ID          GetBotServerRoles
// @Summary     get bot server roles
// @Description get bot server roles
// @security    ApiKeyAuth
// @Produce     json
// @Param       get_roles_options query    SocialAndServerIdReq true "get roles options"
// @Param       social_tool       query    string               true "social tool" Enums(dodo,discord)
// @Success     200               {array}  services.Role
// @Failure     400               {object} appService_errors.RainbowAppServiceErrorDetailInfo "Invalid request"
// @Failure     500               {object} appService_errors.RainbowAppServiceErrorDetailInfo "Internal Server error"
// @Router      /bot/server/roles [get]
func (b *BotServerController) GetRoles(c *gin.Context) {
	var req SocialAndServerIdReq
	if err := c.ShouldBindQuery(&req); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	socialTool, err := enums.ParseSocialToolType(req.SocialTool)
	if err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	roles, err := b.botService.GetRoles(*socialTool, req.ServerId)
	ginutils.RenderResp(c, roles, err)
}

// @Tags        Bot
// @ID          GetInviteUrl
// @Summary     get invite url
// @Description get invite url
// @security    ApiKeyAuth
// @Produce     json
// @Param       social_tool query    string true "social tool" Enums(dodo,discord)
// @Success     200         {string} string
// @Failure     400         {object} appService_errors.RainbowAppServiceErrorDetailInfo "Invalid request"
// @Failure     500         {object} appService_errors.RainbowAppServiceErrorDetailInfo "Internal Server error"
// @Router      /bot/invite_url [get]
func (b *BotServerController) GetInviteUrl(c *gin.Context) {
	var req services.SocialToolQueryReq
	if err := c.ShouldBindQuery(&req); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	socialTool, err := enums.ParseSocialToolType(req.SocialTool)
	if err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	url := b.botService.GetInviteUrl(*socialTool)
	ginutils.RenderRespOK(c, url)
}
