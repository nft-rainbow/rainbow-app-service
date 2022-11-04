package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/nft-rainbow/rainbow-app-service/services"
	"github.com/nft-rainbow/rainbow-app-service/utils/ginutils"
)

func getDiscordChannelInfo(c *gin.Context) {
	guildId := c.Param("guild_id")
	resp, err := services.GetDiscordChannelInfo(guildId)
	ginutils.RenderResp(c, resp, err)
}

func getDoDoChannelInfo(c *gin.Context) {
	islandId := c.Param("island_id")
	resp, err := services.GetDoDoChannelInfo(islandId)
	ginutils.RenderResp(c, resp, err)
}
