package routers

import (
	"encoding/csv"
	"github.com/gin-gonic/gin"
	appService_errors "github.com/nft-rainbow/rainbow-app-service/appService-errors"
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/services"
	"github.com/nft-rainbow/rainbow-app-service/utils/ginutils"
	"strconv"
)

const FileKey = "list"

// @Tags        POAP
// @ID          POAPMintByCSV
// @Summary     POAP Mint By CSV
// @Description POAP Mint By CSV file
// @security    ApiKeyAuth
// @Produce     json
// @Accept      mpfd
// @Param       Authorization header   string true "Bearer JWT"
// @Param       list          formData file   true "uploaded csv file"
// @Param       name          formData string   true "activity name"
// @Param       contract_id   formData string   true "contract id"
// @Success     200           {array} rainbowsdk.ModelsMintTask
// @Failure     400           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Invalid request"
// @Failure     500           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Internal Server error"
// @Router      /poap/csv [post]
func poapMintByCSV(c *gin.Context) {
	poapRequest := services.POAPRequest{}
	poapRequest.Name = c.PostForm("name")
	tmp, _ := strconv.Atoi(c.PostForm("contract_id"))
	poapRequest.ContractID = int32(tmp)

	file, err := c.FormFile(FileKey)
	if err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}
	files, _ := file.Open()
	defer files.Close()

	content, err := csv.NewReader(files).ReadAll()
	if err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	resp, err := services.HandlePOAPCSVMint(content, &poapRequest)
	if err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INTERNAL_SERVER_COMMON)
		return
	}

	ginutils.RenderResp(c, resp, err)
}

// @Tags        POAP
// @ID          POAPMintByH5
// @Summary     POAP Mint By H5
// @Description POAP Mint By H5
// @security    ApiKeyAuth
// @Produce     json
// @Param       Authorization header   string true "Bearer JWT"
// @Param       poap_h5_mint_dto body  services.POAPRequest true "poap_h5_mint_dto"
// @Success     200           {object} rainbowsdk.ModelsMintTask
// @Failure     400           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Invalid request"
// @Failure     500           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Internal Server error"
// @Router      /poap/h5 [post]
func poapMintByH5(c *gin.Context) {
	var poapRequest *services.POAPRequest
	if err := c.ShouldBind(&poapRequest); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	resp, err := services.HandlePOAPH5Mint(poapRequest)

	ginutils.RenderResp(c, resp, err)
}

// @Tags        POAP
// @ID          GetPOAPActivityDetail
// @Summary     Get POAP Activity detail
// @Description Get POAP Activity detail info
// @security    ApiKeyAuth
// @Produce     json
// @Param       Authorization header   string true "Bearer JWT"
// @Param       id            path     int    true "id"
// @Success     200           {object} models.POAPActivityConfig
// @Failure     400           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Invalid request"
// @Failure     500           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Internal Server error"
// @Router      /poap/activity/{id} [get]
func getPOAPActivity(c *gin.Context) {
	activityId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ginutils.RenderRespError(c, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}
	item, err := models.FindPOAPActivityConfigById(activityId)
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
// @Success     200           {object} models.POAPActivityQueryResult
// @Failure     400           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Invalid request"
// @Failure     500           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Internal Server error"
// @Router      /poap/activity [get]
func getPOAPActivityList(c *gin.Context) {
	pagination, err := GetPagination(c)
	if err != nil {
		ginutils.RenderRespError(c, appService_errors.ERR_INVALID_PAGINATION)
		return
	}
	mints, err := models.FindAndCountPOAPActivity(GetIdFromJwtClaim(c), pagination.Offset(), pagination.Limit)
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
// @Success     200           {object} string "success"
// @Failure     400           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Invalid request"
// @Failure     500           {object} appService_errors.RainbowAppServiceErrorDetailInfo "Internal Server error"
// @Router      /poap/activity [post]
func setPOAPActivityConfig(c *gin.Context) {
	var config *models.POAPActivityConfig
	if err := c.ShouldBind(&config); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	err := services.POAPActivityConfig(config, GetIdFromJwtClaim(c))
	ginutils.RenderResp(c, "success", err)
}