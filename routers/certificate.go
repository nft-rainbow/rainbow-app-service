package routers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	appService_errors "github.com/nft-rainbow/rainbow-app-service/appService-errors"
	"github.com/nft-rainbow/rainbow-app-service/models/certificate"
	"github.com/nft-rainbow/rainbow-app-service/services"
	"github.com/nft-rainbow/rainbow-app-service/utils/ginutils"
)

type CertiController struct {
	snapshotService *services.SnapshotService
}

func NewCertiController() *CertiController {
	return &CertiController{
		snapshotService: &services.SnapshotService{},
	}
}

func (ctrl *CertiController) InsertCertificateStrategy(c *gin.Context) {
	var req services.InsertCertificateStrategyReq[any]
	if err := c.ShouldBindJSON(&req); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	cs, err := services.InsertCertificateStrategy(&req)
	ginutils.RenderResp(c, cs, err)
}

func (ctrl *CertiController) GetCertificates(c *gin.Context) {
	csIdStr := c.Param("id")
	csId, err := strconv.Atoi(csIdStr)
	if err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	certificateReuslus, err := services.GetCertificates(uint(csId), c.GetInt("offset"), c.GetInt("limit"))
	ginutils.RenderResp(c, certificateReuslus, err)
}

func (ctrl *CertiController) InsertCertificates(c *gin.Context) {
	csIdStr := c.Param("id")
	csId, err := strconv.Atoi(csIdStr)
	if err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	var items []any
	if err := c.ShouldBindJSON(&items); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_PAGINATION)
		return
	}

	err = func() error {
		cs, err := certificate.FindCertificateStrategyById(uint(csId))
		if err != nil {
			return err
		}
		return cs.InsertCertificates(items)
	}()
	ginutils.RenderResp(c, ginutils.CommonSuccessMessage, err)
}

func (ctrl *CertiController) GetSnapshots(c *gin.Context) {
	cIdStr := c.Param("certificate_id")
	cId, err := strconv.Atoi(cIdStr)
	if err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	resp, err := ctrl.snapshotService.GetContractSnapshots(uint(cId), c.GetInt("offset"), c.GetInt("limit"))
	ginutils.RenderResp(c, resp, err)
}

func (ctrl *CertiController) TriggerRunSnapshot(c *gin.Context) {
	cIdStr := c.Param("certificate_id")
	cId, err := strconv.Atoi(cIdStr)
	if err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	cc, err := certificate.FindContractCertificateById(uint(cId))
	if err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	err = ctrl.snapshotService.Start(cc)
	ginutils.RenderResp(c, ginutils.CommonSuccessMessage, err)
}
