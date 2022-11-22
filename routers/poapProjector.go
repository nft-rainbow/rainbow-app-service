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

func POAPProjectorConfig(c *gin.Context) {
	var adminConfig *models.POAPProjectorConfig
	if err := c.ShouldBind(&adminConfig); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	err := services.POAPProjectorConfig(adminConfig, GetIdFromJwtClaim(c))
	ginutils.RenderResp(c, "success", err)
}

func POAPCSVMint(c *gin.Context) {
	poapRequest := services.POAPRequest{}
	str := c.PostFormMap("data")
	poapRequest.Name = str["name"]
	tmp, _ := strconv.Atoi(str["contract_id"])
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

func POAPH5Mint(c *gin.Context) {
	var poapRequest *services.POAPRequest
	if err := c.ShouldBind(&poapRequest); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	resp, err := services.HandlePOAPH5Mint(poapRequest)

	ginutils.RenderResp(c, resp, err)
}

func getPOAPProjector(c *gin.Context){
	ProjectorId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ginutils.RenderRespError(c, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}
	item, err := models.FindPOAPConfigById(ProjectorId)
	ginutils.RenderResp(c, item, err)
}

func getPOAPProjectorList(c *gin.Context){
	pagination, err := GetPagination(c)
	if err != nil {
		ginutils.RenderRespError(c, appService_errors.ERR_INVALID_PAGINATION)
		return
	}
	mints, err := models.FindAndCountPOAPProjectorConfig(GetIdFromJwtClaim(c), pagination.Offset(), pagination.Limit)
	ginutils.RenderResp(c, mints, err)
}

func getPOAPActivity(c *gin.Context) {
	activityId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ginutils.RenderRespError(c, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}
	item, err := models.FindPOAPActivityConfigById(activityId)
	ginutils.RenderResp(c, item, err)
}

func getPOAPActivityList(c *gin.Context) {
	pagination, err := GetPagination(c)
	if err != nil {
		ginutils.RenderRespError(c, appService_errors.ERR_INVALID_PAGINATION)
		return
	}
	mints, err := models.FindAndCountPOAPActivity(GetIdFromJwtClaim(c), pagination.Offset(), pagination.Limit)
	ginutils.RenderResp(c, mints, err)
}

func POAPActivityConfig(c *gin.Context) {
	var config *models.POAPActivityConfig
	if err := c.ShouldBind(&config); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	err := services.POAPActivityConfig(config, GetIdFromJwtClaim(c))
	ginutils.RenderResp(c, "success", err)
}