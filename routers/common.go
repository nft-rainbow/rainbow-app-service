package routers

import (
	"github.com/gin-gonic/gin"
	appService_errors "github.com/nft-rainbow/rainbow-app-service/appService-errors"
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/services"
	"github.com/nft-rainbow/rainbow-app-service/utils"
	"github.com/nft-rainbow/rainbow-app-service/utils/ginutils"
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

func pushActivity(c *gin.Context) {
	var req *models.PushReq
	if err := c.ShouldBind(&req); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	if req.Bot == utils.Discord {
		resp, err := services.DiscordPushActivity(req)
		ginutils.RenderResp(c, resp, err)
	}
}

func GetPushDiscord(c *gin.Context) {
	var req *models.PushReq
	if err := c.ShouldBind(&req); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	if req.Bot == utils.Discord {
		resp, err := services.DiscordPushActivity(req)
		ginutils.RenderResp(c, resp, err)
	}
}
