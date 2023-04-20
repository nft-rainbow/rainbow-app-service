package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mcuadros/go-defaults"
	appService_errors "github.com/nft-rainbow/rainbow-app-service/appService-errors"
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/services"
	"github.com/nft-rainbow/rainbow-app-service/utils/ginutils"
)

type BotServer struct {
	botService *services.BotServerService
}

type (
	SocialToolReq struct {
		SocialTool models.SocialToolType `uri:"social_tool" form:"social_tool" binding:"required"`
	}

	SocialToolMaybeReq struct {
		SocialTool *models.SocialToolType `uri:"social_tool" form:"social_tool"`
	}

	IdUintReq struct {
		Id uint `uri:"id" form:"id" binding:"required"`
	}

	ServerIdReq struct {
		ServerId string `uri:"server_id" form:"server_id" binding:"required"`
	}

	SocialAndIdUriReq struct {
		SocialToolReq
		IdUintReq
	}

	SocialAndServerIdReq struct {
		ServerIdReq
		SocialTool models.SocialToolType `form:"social_tool" binding:"required"`
	}
)

func NewBotServer() (*BotServer, error) {
	service, err := services.NewBotServerService()
	if err != nil {
		return nil, err
	}
	return &BotServer{
		botService: service,
	}, nil
}

func (b *BotServer) verifyBotServer(c *gin.Context) {
	var verifyUserReq services.VerifySocialServerReq
	if err := c.ShouldBind(&verifyUserReq); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	err := b.botService.VerifyBotServer(verifyUserReq.SocialTool, verifyUserReq.ServerId)
	ginutils.RenderResp(c, services.Success, err)
}

func (b *BotServer) insertBotServer(c *gin.Context) {
	var req services.InsertSocialServerReq
	if err := c.ShouldBind(&req); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	userId := GetIdFromJwtClaim(c)
	err := b.botService.InsertBotServer(userId, req)
	ginutils.RenderResp(c, gin.H{"result": "ok"}, err)
}

func (b *BotServer) GetBotServers(c *gin.Context) {
	var queryParams SocialToolMaybeReq
	if err := c.ShouldBindQuery(&queryParams); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	userId := GetIdFromJwtClaim(c)
	p, err := b.botService.GetBotServers(userId, queryParams.SocialTool)
	ginutils.RenderResp(c, p, err)
}

func (b *BotServer) GetActivitiesOfUserBotServers(c *gin.Context) {
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

func (b *BotServer) GetBotServer(c *gin.Context) {
	userId := GetIdFromJwtClaim(c)

	var req IdUintReq
	if err := c.ShouldBindUri(&req); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	p, err := b.botService.GetBotServer(userId, req.Id)
	ginutils.RenderResp(c, p, err)
}

func (b *BotServer) AddActivity(c *gin.Context) {
	var req services.PushInfoReq
	if err := c.ShouldBind(&req); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}
	var uriParams IdUintReq
	if err := c.ShouldBindUri(&uriParams); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	userId := GetIdFromJwtClaim(c)
	res, err := b.botService.AddActivity(userId, uriParams.Id, req)
	ginutils.RenderResp(c, res, err)
}

func (b *BotServer) Push(c *gin.Context) {
	var uriParams IdUintReq
	if err := c.ShouldBindUri(&uriParams); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	userId := GetIdFromJwtClaim(c)
	err := b.botService.Push(userId, uriParams.Id)
	ginutils.RenderResp(c, services.Success, err)
}

func (b *BotServer) UpdateActivity(c *gin.Context) {
	var req services.PushInfoReq
	if err := c.ShouldBind(&req); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}
	var uriParams IdUintReq
	if err := c.ShouldBindUri(&uriParams); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	userId := GetIdFromJwtClaim(c)
	res, err := b.botService.UpdateActivity(userId, uriParams.Id, req)
	ginutils.RenderResp(c, res, err)

}

func (b *BotServer) GetChannels(c *gin.Context) {
	var req SocialAndServerIdReq
	if err := c.ShouldBindQuery(&req); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	channels, err := b.botService.GetChannels(req.SocialTool, req.ServerId)
	ginutils.RenderResp(c, channels, err)
}

func (b *BotServer) GetRoles(c *gin.Context) {
	var req SocialAndServerIdReq
	if err := c.ShouldBindQuery(&req); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	roles, err := b.botService.GetRoles(req.SocialTool, req.ServerId)
	ginutils.RenderResp(c, roles, err)
}

func (b *BotServer) GetInviteUrl(c *gin.Context) {
	var req SocialToolReq
	if err := c.ShouldBindQuery(&req); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}
	
}

func mustGetSocialToolFromPath(c *gin.Context) *models.SocialToolType {
	return mustGetSocialToolType(c.Param("social_tool"))
}

func mustGetSocialToolType(name string) *models.SocialToolType {
	if t, ok := models.ParseSocialToolType(name); ok {
		return t
	}
	panic("unsupport")
}

func getUriServerId(c *gin.Context) string {
	return c.Param("server_id")
}

func getUriId(c *gin.Context) string {
	return c.Param("id")
}
