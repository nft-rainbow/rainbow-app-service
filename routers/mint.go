package routers

import (
	"github.com/gin-gonic/gin"
	discordbot_errors "github.com/nft-rainbow/discordbot-service/discordbot-errors"
	"github.com/nft-rainbow/discordbot-service/models"
	"github.com/nft-rainbow/discordbot-service/services"
	"github.com/nft-rainbow/discordbot-service/utils/ginutils"
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

func handleCustomMint(c *gin.Context) {
	var req *models.MintReq
	if err := c.ShouldBind(&req); err != nil {
		ginutils.RenderRespError(c, err, discordbot_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}
	resp, err := services.CustomMint(req)
	ginutils.RenderResp(c, resp, err)
}
