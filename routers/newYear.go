package routers

import (
	"github.com/gin-gonic/gin"
	appService_errors "github.com/nft-rainbow/rainbow-app-service/appService-errors"
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/services"
	"github.com/nft-rainbow/rainbow-app-service/utils/ginutils"
	"strconv"
)

// @Tags        NewYear
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
// @Router      /poap/newYear/config [post]
func setNewYearConfig(c *gin.Context) {
	var config *models.NewYearConfig
	if err := c.ShouldBind(&config); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}
	resp, err := services.SetNewYearConfig(config, GetIdFromJwtClaim(c))
	ginutils.RenderResp(c, resp, err)
}

// @Tags        NewYear
// @ID          Update By Sharing
// @Summary     Update By Sharing
// @Description Update By Sharing
// @security    ApiKeyAuth
// @Produce     json
// @Param       Authorization header   string true "Bearer JWT"
// @Param       share_request body  services.ShareRequest true "share_request"
// @Success     200           {object} string "success"
// @Failure     400           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Invalid request"
// @Failure     500           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Internal Server error"
// @Router      /poap/newYear/sharer [post]
func updateBySharing(c *gin.Context){
	var req services.ShareRequest
	if err := c.ShouldBind(&req); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}
	err := services.UpdateBySharing(req)
	ginutils.RenderResp(c, "success", err)
}

// @Tags        NewYear
// @ID          GetMintCount
// @Summary     Get Common Mint Count
// @Description Get Common Mint Count
// @security    ApiKeyAuth
// @Produce     json
// @Param       Authorization header   string true "Bearer JWT"
// @Param       id            path     int    true "id"
// @Param       address       path     string true "address"
// @Success     200           {object} models.MintCount
// @Failure     400           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Invalid request"
// @Failure     500           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Internal Server error"
// @Router      /poap/count/common/{address}/{id} [get]
func getCommonMintCount(c *gin.Context) {
	address := c.Param("address")
	activityIdStr := c.Param("id")
	activityId, err := strconv.Atoi(activityIdStr)
	if err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	resp, err := models.FindMintCount(address, int32(activityId))
	ginutils.RenderResp(c, resp, err)
}

// @Tags        NewYear
// @ID          GetSpecialMintRemain
// @Summary     Get Special Mint Count Remain
// @Description Get Special Mint Count Remain
// @security    ApiKeyAuth
// @Produce     json
// @Param       Authorization header   string true "Bearer JWT"
// @Param       id            path     int    true "id"
// @Param       address       path     string true "address"
// @Success     200           {object} int
// @Failure     400           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Invalid request"
// @Failure     500           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Internal Server error"
// @Router      /poap/count/special/{address}/{id} [get]
func getSpecialMintCount(c *gin.Context) {
	address := c.Param("address")
	activityIdStr := c.Param("id")

	activityId, err := strconv.Atoi(activityIdStr)
	if err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}
	resp, err := services.GetSpecialMintCount(activityId, address)

	ginutils.RenderResp(c, resp, err)
}