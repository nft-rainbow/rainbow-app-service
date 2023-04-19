package routers

import (
	"strconv"

	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/gin-gonic/gin"
	appService_errors "github.com/nft-rainbow/rainbow-app-service/appService-errors"
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/services"
	"github.com/nft-rainbow/rainbow-app-service/utils/ginutils"
)

// @Tags        POAP
// @ID          POAPMintByCSV
// @Summary     POAP Mint By CSV
// @Description POAP Mint By CSV file
// @security    ApiKeyAuth
// @Produce     json
// @Param       poap_csv_mint_dto body  services.POAPRequest true "poap_csv_mint_dto"
// @Success     200           {object} rainbowsdk.ModelsMintTask
// @Failure     400           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Invalid request"
// @Failure     500           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Internal Server error"
// @Router      /poap/csv [post]
// func poapMintByCSV(c *gin.Context) {
// 	poapRequest := services.MintReq{}
// 	if err := c.ShouldBind(&poapRequest); err != nil {
// 		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
// 		return
// 	}

// 	resp, err := services.HandlePOAPCSVMint(&poapRequest)
// 	if err != nil {
// 		ginutils.RenderRespError(c, err, appService_errors.ERR_INTERNAL_SERVER_COMMON)
// 		return
// 	}

// 	ginutils.RenderResp(c, resp, err)
// }

// @Tags        POAP
// @ID          POAPMintByH5
// @Summary     POAP Mint By H5
// @Description POAP Mint By H5
// @security    ApiKeyAuth
// @Produce     json
// @Param       poap_h5_mint_dto body  services.POAPRequest true "poap_h5_mint_dto"
// @Success     200           {object} models.POAPResult
// @Failure     400           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Invalid request"
// @Failure     500           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Internal Server error"
// @Router      /poap/h5 [post]
func poapMintByH5(c *gin.Context) {
	var poapRequest *services.MintReq
	if err := c.ShouldBind(&poapRequest); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}
	if _, err := cfxaddress.NewFromBase32(poapRequest.UserAddress); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	resp, err := activityService.HandlePOAPH5Mint(poapRequest)
	ginutils.RenderResp(c, resp, err)
}

// @Tags        POAP
// @ID          GetPOAPActivityDetail
// @Summary     Get POAP Activity detail
// @Description Get POAP Activity detail info
// @security    ApiKeyAuth
// @Produce     json
// @Param       activity_id       path     string    true "activity_id"
// @Success     200           {object} models.POAPActivityConfig
// @Failure     400           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Invalid request"
// @Failure     500           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Internal Server error"
// @Router      /poap/activity/{activity_id} [get]
func getActivity(c *gin.Context) {
	poapId := c.Param(ACTIVITY_CODE_KEY)
	item, err := models.FindActivityByCode(poapId)
	// if item.NeedCommand() == true {
	// 	item.Command = ""
	// }
	ginutils.RenderResp(c, item, err)
}

// @Tags        POAP
// @ID          GetPOAPActivityList
// @Summary     Get POAP Activity list
// @Description Get POAP Activity list
// @security    ApiKeyAuth
// @Produce     json
// @Param       Authorization header   string true "Bearer JWT"
// @Param       page          query    integer false "page"
// @Param       limit         query    integer false "limit"
// @Param       contract_address          query    string false "contract_address"
// @Param       activity_id         query    string false "activity_id"
// @Param       name         query    string false "name"
// @Success     200           {object} models.POAPActivityQueryResult
// @Failure     400           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Invalid request"
// @Failure     500           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Internal Server error"
// @Router      /poap/activity [get]
func getActivities(c *gin.Context) {
	pagination, err := GetPagination(c)
	if err != nil {
		ginutils.RenderRespError(c, appService_errors.ERR_INVALID_PAGINATION)
		return
	}

	var cond models.ActivityFindCondition
	if err := c.ShouldBindQuery(&cond); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	mints, err := models.FindAndCountActivity(GetIdFromJwtClaim(c), pagination.Offset(), pagination.Limit, cond)
	ginutils.RenderResp(c, mints, err)
}

// @Tags        POAP
// @ID          SetPOAPActivity
// @Summary     Set POAP Activity
// @Description Set POAP Activity
// @security    ApiKeyAuth
// @Produce     json
// @Param       Authorization header   string true "Bearer JWT"
// @Param       poap_activity_config body  models.POAPActivityConfig true "poap_activity_config"
// @Success     200           {object} models.POAPActivityConfig
// @Failure     400           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Invalid request"
// @Failure     500           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Internal Server error"
// @Router      /poap/activity [post]
func addActivity(c *gin.Context) {
	var config *models.ActivityReq
	if err := c.ShouldBind(&config); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	resp, err := activityService.InsertActivity(config, GetIdFromJwtClaim(c))
	ginutils.RenderResp(c, resp, err)
}

// @Tags        POAP
// @ID          SetH5Config
// @Summary     Set H5 Config
// @Description Set H5 Page Config
// @security    ApiKeyAuth
// @Produce     json
// @Param       Authorization header   string true "Bearer JWT"
// @Param       poap_h5_config body  models.H5Config true "poap_h5_config"
// @Success     200           {object} models.H5Config
// @Failure     400           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Invalid request"
// @Failure     500           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Internal Server error"
// @Router      /poap/activity/h5 [post]
func setActivityH5Config(c *gin.Context) {
	var config *models.H5Config
	if err := c.ShouldBind(&config); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	resp, err := activityService.POAPH5Config(config)
	ginutils.RenderResp(c, resp, err)
}

// @Tags        POAP
// @ID          UpdatePOAPActivity
// @Summary     Update POAP Activity
// @Description Update POAP Activity
// @security    ApiKeyAuth
// @Produce     json
// @Param       Authorization header   string true "Bearer JWT"
// @Param       activity_id       path     string    true "activity_id"
// @Param       poap_activity_config body  models.POAPActivityConfig true "poap_activity_config"
// @Success     200           {object} models.POAPActivityConfig
// @Failure     400           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Invalid request"
// @Failure     500           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Internal Server error"
// @Router      /poap/activity/{activity_id} [put]
func updateActivity(c *gin.Context) {
	var config models.UpdateActivityReq
	if err := c.ShouldBind(&config); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}
	activityCode := c.Param(ACTIVITY_CODE_KEY)
	// config.ActivityID = poapId

	resp, err := activityService.UpdateActivity(activityCode, &config)
	ginutils.RenderResp(c, resp, err)
}

// @Tags        POAP
// @ID          GetPOAPResultList
// @Summary     Get POAP Result list
// @Description Get POAP Result list
// @security    ApiKeyAuth
// @Produce     json
// @Param       page          query    integer false "page"
// @Param       limit         query    integer false "limit"
// @Param       activity_id       path     string    true "activity_id"
// @Success     200           {object} models.POAPResultQueryResult
// @Failure     400           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Invalid request"
// @Failure     500           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Internal Server error"
// @Router      /poap/activity/result/{activity_id} [get]
func getPOAPResultList(c *gin.Context) {
	pagination, err := GetPagination(c)
	if err != nil {
		ginutils.RenderRespError(c, appService_errors.ERR_INVALID_PAGINATION)
		return
	}
	poapId := c.Param(ACTIVITY_CODE_KEY)
	if poapId == "" {
		ginutils.RenderRespError(c, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}
	mints, err := models.FindAndCountPOAPResult(poapId, pagination.Offset(), pagination.Limit)
	ginutils.RenderResp(c, mints, err)
}

// @Tags        POAP
// @ID          GetPOAPResult
// @Summary     Get POAP Result
// @Description Get POAP Result
// @security    ApiKeyAuth
// @Produce     json
// @Param       activity_id   path     string    true "activity_id"
// @Param       id            path     int    true "id"
// @Success     200           {object} models.POAPResult
// @Failure     400           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Invalid request"
// @Failure     500           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Internal Server error"
// @Router      /poap/activity/result/{activity_id}/{id} [get]
func getPOAPResultDetail(c *gin.Context) {
	poapId := c.Param(ACTIVITY_CODE_KEY)
	if poapId == "" {
		ginutils.RenderRespError(c, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ginutils.RenderRespError(c, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}
	resp, err := models.FindPOAPResultById(poapId, id)
	ginutils.RenderResp(c, resp, err)
}

// @Tags        POAP
// @ID          GetMintCount
// @Summary     Get Common Mint Count
// @Description Get Common Mint Count
// @security    ApiKeyAuth
// @Produce     json
// @Param       activity_id   path     string true "activity_id"
// @Param       address       path     string true "address"
// @Success     200           {object} int
// @Failure     400           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Invalid request"
// @Failure     500           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Internal Server error"
// @Router      /poap/count/{address}/{activity_id} [get]
func getMintCount(c *gin.Context) {
	var err error
	address := c.Param("address")
	activityCode := c.Param(ACTIVITY_CODE_KEY)

	var resp *int32
	resp, err = activityService.GetMintCount(activityCode, address)
	ginutils.RenderResp(c, resp, err)
}

// @Tags        POAP
// @ID          CollectUserAnywebCode
// @Summary     Collect user code so backend can get user phone from anyweb
// @Description Collect user code so backend can get user phone from anyweb
// @security    ApiKeyAuth
// @Produce     json
// @Success     200           {object} int
// @Failure     400           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Invalid request"
// @Failure     500           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Internal Server error"
// @Router      /poap/anyweb/code [POST]
func addWalletUser(c *gin.Context) {
	var config services.AddWalletUserReq
	if err := c.ShouldBind(&config); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	err := walletService.InsertUser(config)

	ginutils.RenderResp(c, map[string]string{"message": "success"}, err)
}
