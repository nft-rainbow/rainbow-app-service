package routers

import (
	"github.com/gin-gonic/gin"
	appService_errors "github.com/nft-rainbow/rainbow-app-service/appService-errors"
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/services"
	"github.com/nft-rainbow/rainbow-app-service/utils/ginutils"
	"github.com/spf13/viper"
)

// @Tags        POAP
// @ID          SetNewYearActivity
// @Summary     Set NewYear Activity
// @Description Set NewYear Activity
// @security    ApiKeyAuth
// @Produce     json
// @Param       Authorization header   string true "Bearer JWT"
// @Param       new_year_activity_config body  models.NewYearConfig true "new_year_activity_config"
// @Success     200           {object} models.NewYearConfig
// @Failure     400           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Invalid request"
// @Failure     500           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Internal Server error"
// @Router      /poap/config [post]
func setNewYearConfig(c *gin.Context) {
	var config *models.NewYearConfig
	if err := c.ShouldBind(&config); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}
	resp, err := services.SetNewYearConfig(config, GetIdFromJwtClaim(c))
	ginutils.RenderResp(c, resp, err)
}

// @Tags        POAP
// @ID          Update By Sharing
// @Summary     Update By Sharing
// @Description Update By Sharing
// @security    ApiKeyAuth
// @Produce     json
// @Param       share_request body  services.ShareRequest true "share_request"
// @Success     200           {object} string "success"
// @Failure     400           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Invalid request"
// @Failure     500           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Internal Server error"
// @Router      /poap/sharer [post]
func updateBySharing(c *gin.Context) {
	var req services.ShareRequest
	if err := c.ShouldBind(&req); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}
	err := services.UpdateBySharing(req)
	ginutils.RenderResp(c, "success", err)
}

// @Tags        POAP
// @ID          GetMintCount
// @Summary     Get Common Mint Count
// @Description Get Common Mint Count
// @security    ApiKeyAuth
// @Produce     json
// @Param       activity_id   path     string true "activity_id"
// @Param       address       path     string true "address"
// @Success     200           {object} services.MintCountResponse
// @Failure     400           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Invalid request"
// @Failure     500           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Internal Server error"
// @Router      /poap/count/{address}/{activity_id} [get]
func getMintCount(c *gin.Context) {
	var err error
	address := c.Param("address")
	poapId := c.Param("activity_id")

	var resp *services.MintCountResponse
	if poapId == viper.GetString("newYearEvent.newYearCommonId") {
		resp, err = services.GetCommonMintCount(address, poapId)
		ginutils.RenderResp(c, resp, err)
	} else if poapId == viper.GetString("newYearEvent.newYearSpecialId") {
		resp, err = services.GetSpecialMintCount(address, viper.GetString("newYearEvent.newYearCommonId"))
		ginutils.RenderResp(c, resp, err)
	} else {
		resp, err = services.GetMintCount(poapId, address)
		ginutils.RenderResp(c, resp, err)
	}
}
