package routers

import (
	"github.com/gin-gonic/gin"
	appService_errors "github.com/nft-rainbow/rainbow-app-service/appService-errors"
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/services"
	"github.com/nft-rainbow/rainbow-app-service/utils"
	"github.com/nft-rainbow/rainbow-app-service/utils/ginutils"
	"strconv"
)

// @Tags        CustomMint
// @ID          GetDiscordChannelDetail
// @Summary     Discord Channel detail
// @Description Get Discord Channel detail info
// @security    ApiKeyAuth
// @Produce     json
// @Param       guild_id      path     int    true "guild_id"
// @Success     200           {array} discordgo.Channel
// @Failure     500           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Internal Server error"
// @Router      /custom/discord/{guild_id}/channels [get]
func getDiscordChannelInfo(c *gin.Context) {
	guildId := c.Param("guild_id")
	resp, err := services.GetDiscordChannelInfo(guildId)

	ginutils.RenderResp(c, resp, err)
}

// @Tags        CustomMint
// @ID          GetDoDoChannelDetail
// @Summary     DoDo Channel detail
// @Description Get DoDo Channel detail info
// @security    ApiKeyAuth
// @Produce     json
// @Param       island_id     path     int    true "island_id"
// @Success     200           {array} model.ChannelElement
// @Failure     500           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Internal Server error"
// @Router      /custom/dodo/{island_id}/channels [get]
func getDoDoChannelInfo(c *gin.Context) {
	islandId := c.Param("island_id")
	resp, err := services.GetDoDoChannelInfo(islandId)
	ginutils.RenderResp(c, resp, err)
}

// @Tags        POAP
// @ID          PushActivity
// @Summary     Push Activity
// @Description Push Activity Info
// @security    ApiKeyAuth
// @Produce     json
// @Param       push_req     body     services.PushReq    true "push_req"
// @Success     200           {object} string "success"
// @Failure     500           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Internal Server error"
// @Router      /poap/activity/push [post]
func pushActivity(c *gin.Context) {
	var req *services.PushReq
	if err := c.ShouldBind(&req); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	req.RainbowUserId = int32(GetIdFromJwtClaim(c))

	if req.Bot == utils.Discord {
		_, err := services.DiscordPushActivity(req)
		ginutils.RenderResp(c, "success", err)
	} else if req.Bot == utils.DoDo {
		_, err := services.DoDoPushActivity(req)
		ginutils.RenderResp(c, "success", err)
	} else {
		ginutils.RenderRespError(c, nil, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}
}

// @Tags        POAP
// @ID          GetPush
// @Summary     Get Pushes List
// @Description Get Pushes List
// @security    ApiKeyAuth
// @Produce     json
// @Param       app_id        path     string   true "app_id"
// @Param       bot           path     string   true "bot"
// @Success     200           {object} models.PushInfoQueryResult
// @Failure     500           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Internal Server error"
// @Router      /poap/activity/push/{app_id}/{bot} [get]
func getPushes(c *gin.Context) {
	pagination, err := GetPagination(c)
	if err != nil {
		ginutils.RenderRespError(c, appService_errors.ERR_INVALID_PAGINATION)
		return
	}
	appIdStr := c.Param("app_id")
	appId, err := strconv.Atoi(appIdStr)
	if err != nil {
		ginutils.RenderRespError(c, appService_errors.ERR_INVALID_PAGINATION)
		return
	}

	botStr := c.Param("bot")
	var bot uint
	if botStr == "discord" {
		bot = utils.Discord
	} else {
		bot = utils.DoDo
	}
	resp, err := models.FindAndCountPushInfo(pagination.Offset(), pagination.Limit, appId, int(GetIdFromJwtClaim(c)), bot)
	ginutils.RenderResp(c, resp, err)
}
