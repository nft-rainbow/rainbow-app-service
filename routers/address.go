package routers

import (
	"errors"
	"github.com/gin-gonic/gin"
	discordbot_errors "github.com/nft-rainbow/discordbot-service/discordbot-errors"
	"github.com/nft-rainbow/discordbot-service/models"
	"github.com/nft-rainbow/discordbot-service/services"
	"github.com/nft-rainbow/discordbot-service/utils/ginutils"
)

func bindAdminConfig(c *gin.Context) {
	var adminConfig *models.AdminConfig
	if err := c.ShouldBind(&adminConfig); err != nil {
		ginutils.RenderRespError(c, err, discordbot_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	err := services.BindAdminConfig(adminConfig)
	ginutils.RenderResp(c, "success", err)
}

func bindUserAddress(c *gin.Context) {
	var bindReq *models.BindCFXAddress
	if err := c.ShouldBind(&bindReq); err != nil {
		ginutils.RenderRespError(c, err, discordbot_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	err := services.BindCFXAddress(bindReq)
	ginutils.RenderResp(c, "success", err)

}

func getUserBindingAddress(c *gin.Context) {
	userID := c.Param("user_id")
	if userID == "" {
		ginutils.RenderRespError(c, errors.New("empty userID"), discordbot_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	resp, err := services.GetBindCFXAddress(userID)
	ginutils.RenderResp(c, resp, err)
}


