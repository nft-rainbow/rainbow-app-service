package routers

import (
	"github.com/gin-gonic/gin"
	appService_errors "github.com/nft-rainbow/rainbow-app-service/appService-errors"
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/services"
	"github.com/nft-rainbow/rainbow-app-service/utils/ginutils"
)

type BotActivityRouteHandler struct {
	Dodo        *services.BotActivityService
	Discord     *services.BotActivityService
	botServices map[models.SocialToolType]*services.BotActivityService
}

func NewBotActivityRouteHandler() (*BotActivityRouteHandler, error) {
	dodo, err := services.NewBotActivityService(models.SOCIAL_TOOL_DODO)
	if err != nil {
		return nil, err
	}
	handler := &BotActivityRouteHandler{
		Dodo: dodo,
		botServices: map[models.SocialToolType]*services.BotActivityService{
			models.SOCIAL_TOOL_DODO: dodo,
		},
	}

	return handler, nil
}

func (b *BotActivityRouteHandler) geBotService(socialToolType models.SocialToolType) *services.BotActivityService {
	if v, ok := b.botServices[socialToolType]; ok {
		return v
	}
	panic("unsupported social tool type")
}

func (b *BotActivityRouteHandler) verifyUser(c *gin.Context) {
	var verifyUserReq VerifySocialUserReq
	if err := c.ShouldBind(&verifyUserReq); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	socialTool := mustGetSocialToolFromPath(c)
	verifyResp, err := b.geBotService(*socialTool).VerifyUser(models.SocialToolUser{
		SocialTool:   *socialTool,
		UserSocialId: verifyUserReq.UserSocialId,
	})
	ginutils.RenderResp(c, verifyResp, err)
}

func (b *BotActivityRouteHandler) insertProjectManager(c *gin.Context) {
	var req services.InsertProjectorReq
	if err := c.ShouldBind(&req); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}
	userId := GetIdFromJwtClaim(c)
	socialTool := mustGetSocialToolFromPath(c)
	err := b.geBotService(*socialTool).InsertProjectManager(userId, *socialTool, req)
	ginutils.RenderResp(c, gin.H{"result": "ok"}, err)
}

func (b *BotActivityRouteHandler) getProjectManager(c *gin.Context) {
	userId := GetIdFromJwtClaim(c)
	socialTool := mustGetSocialToolFromPath(c)
	p, err := b.geBotService(*socialTool).GetProjectManager(userId, *socialTool)
	ginutils.RenderResp(c, p, err)
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
