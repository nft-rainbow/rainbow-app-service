package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/nft-rainbow/rainbow-app-service/services"
	"github.com/nft-rainbow/rainbow-app-service/utils/ginutils"
)

// @Tags        CustomMint
// @ID          GetDiscordChannelDetail
// @Summary     Discord Channel detail
// @Description Get Discord Channel detail info
// @security    ApiKeyAuth
// @Produce     json
// @Param       Authorization header   string true "Bearer JWT"
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
// @Param       Authorization header   string true "Bearer JWT"
// @Param       island_id     path     int    true "island_id"
// @Success     200           {array} model.ChannelElement
// @Failure     500           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Internal Server error"
// @Router      /custom/dodo/{island_id}/channels [get]
func getDoDoChannelInfo(c *gin.Context) {
	islandId := c.Param("island_id")
	resp, err := services.GetDoDoChannelInfo(islandId)
	ginutils.RenderResp(c, resp, err)
}
