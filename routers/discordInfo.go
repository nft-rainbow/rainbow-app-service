package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/nft-rainbow/rainbow-app-service/services"
	"github.com/nft-rainbow/rainbow-app-service/utils/ginutils"
)

func getChannelInfo(c *gin.Context) {
	guildId := c.Param("guild_id")
	resp, err := services.GetChannelInfo(guildId)
	ginutils.RenderResp(c, resp, err)
}
