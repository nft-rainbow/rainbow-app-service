package routers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	appService_errors "github.com/nft-rainbow/rainbow-app-service/appService-errors"
	"github.com/nft-rainbow/rainbow-app-service/models/certificate"
	"github.com/nft-rainbow/rainbow-app-service/models/enums"
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
	certiTypeInStr := c.Param("certificate_type")
	certiType, err := enums.ParseCertificateType(certiTypeInStr)
	if err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	var items []any
	if err := c.ShouldBindJSON(&items); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	req := services.InsertCertificateStrategyReq[any]{
		CertificateType: *certiType,
		Items:           items,
	}

	cs, err := services.InsertCertificateStrategy(&req)
	ginutils.RenderResp(c, cs, err)
}

//	@Tags			Certi
//	@ID				GetCertificates
//	@Summary		Get Certificates
//	@Description	Get Certificates
//	@security		ApiKeyAuth
//	@Produce		json
//	@Param			id	path		int	true	"certificate_strategy_id"
//	@Success		200	{object}	certificate.CertificatesQueryResult[CompositedCertificate]
//	@Failure		400	{object}	appService_errors.RainbowAppServiceErrorDetailInfo	"Invalid request"
//	@Failure		500	{object}	appService_errors.RainbowAppServiceErrorDetailInfo	"Internal Server error"
//	@Router			/certis/strategy/{id}/certificates [get]
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

//	@Tags			Certi
//	@ID				InsertCertificates
//	@Summary		Insert Certificates
//	@Description	Insert Certificates
//	@security		ApiKeyAuth
//	@Produce		json
//	@Param			id				path		int						true	"certificate_strategy_id"
//	@Param			certificates	body		[]CompositedCertificate	true	"certificate"
//	@Success		200				{object}	ginutils.CommonMessage
//	@Failure		400				{object}	appService_errors.RainbowAppServiceErrorDetailInfo	"Invalid request"
//	@Failure		500				{object}	appService_errors.RainbowAppServiceErrorDetailInfo	"Internal Server error"
//	@Router			/certis/strategy/{id}/certificates [post]
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

//	@Tags			Certi
//	@ID				DeleteCertificates
//	@Summary		Delete Certificates
//	@Description	Delete Certificates
//	@security		ApiKeyAuth
//	@Produce		json
//	@Param			id				path		int		true	"certificate_strategy_id"
//	@Param			certificate_ids	body		[]uint	true	"certificate_ids"
//	@Success		200				{object}	ginutils.CommonMessage
//	@Failure		400				{object}	appService_errors.RainbowAppServiceErrorDetailInfo	"Invalid request"
//	@Failure		500				{object}	appService_errors.RainbowAppServiceErrorDetailInfo	"Internal Server error"
//	@Router			/certis/strategy/{id}/certificates [delete]
func (ctrl *CertiController) DeleteCertificates(c *gin.Context) {
	csIdStr := c.Param("id")
	csId, err := strconv.Atoi(csIdStr)
	if err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	var certificateIds []uint
	if err := c.ShouldBindJSON(&certificateIds); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_PAGINATION)
		return
	}

	err = func() error {
		cs, err := certificate.FindCertificateStrategyById(uint(csId))
		if err != nil {
			return err
		}
		return cs.DeleteCertificates(certificateIds)
	}()
	ginutils.RenderResp(c, ginutils.CommonSuccessMessage, err)
}

//	@Tags			Certi
//	@ID				GetSnapshots
//	@Summary		Get Snapshots
//	@Description	Get Snapshots
//	@security		ApiKeyAuth
//	@Produce		json
//	@Param			certificate_id	path		int	true	"certificate_id"
//	@Success		200				{object}	services.ContractSnapshotResp
//	@Failure		400				{object}	appService_errors.RainbowAppServiceErrorDetailInfo	"Invalid request"
//	@Failure		500				{object}	appService_errors.RainbowAppServiceErrorDetailInfo	"Internal Server error"
//	@Router			/certis/contract_certificate/{certificate_id}/snapshot [get]
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

//	@Tags			Certi
//	@ID				TriggerObtainSnapshot
//	@Summary		Trigger Obtain Snapshot
//	@Description	Trigger Obtain Snapshot
//	@security		ApiKeyAuth
//	@Produce		json
//	@Param			certificate_id	path		int	true	"certificate_id"
//	@Success		200				{object}	ginutils.CommonMessage
//	@Failure		400				{object}	appService_errors.RainbowAppServiceErrorDetailInfo	"Invalid request"
//	@Failure		500				{object}	appService_errors.RainbowAppServiceErrorDetailInfo	"Internal Server error"
//	@Router			/certis/contract_certificate/{certificate_id}/snapshot/run [post]
func (ctrl *CertiController) TriggerObtainSnapshot(c *gin.Context) {
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

// ============================= only for gen swagger =============================

type CompositedCertificate struct {
	certificate.AddressCertificateInsertPart
	certificate.PhoneCertificateInsertPart
	certificate.DodoCertificateInsertPart
	certificate.ContractCertificateInsertPart
	certificate.GaslessCertificateInsertPart
}

//	@Tags			Certi
//	@ID				InsertAddressCertificateStrategy
//	@Summary		Insert address Certificate Strategy
//	@Description	Insert address Certificate Strategy
//	@security		ApiKeyAuth
//	@Produce		json
//	@Param			insert_certificate_strategy_req	body		[]certificate.AddressCertificateInsertPart	true	"insert_certificate_strategy_req"
//	@Success		200								{object}	certificate.CertificateStrategy
//	@Failure		400								{object}	appService_errors.RainbowAppServiceErrorDetailInfo	"Invalid request"
//	@Failure		500								{object}	appService_errors.RainbowAppServiceErrorDetailInfo	"Internal Server error"
//	@Router			/certis/strategy/type/address [post]
func (ctrl *CertiController) insertAddressCertificateStrategy(c *gin.Context) {}

//	@Tags			Certi
//	@ID				InsertPhoneCertificateStrategy
//	@Summary		Insert phone Certificate Strategy
//	@Description	Insert phone Certificate Strategy
//	@security		ApiKeyAuth
//	@Produce		json
//	@Param			insert_certificate_strategy_req	body		[]certificate.PhoneCertificateInsertPart	true	"insert_certificate_strategy_req"
//	@Success		200								{object}	certificate.CertificateStrategy
//	@Failure		400								{object}	appService_errors.RainbowAppServiceErrorDetailInfo	"Invalid request"
//	@Failure		500								{object}	appService_errors.RainbowAppServiceErrorDetailInfo	"Internal Server error"
//	@Router			/certis/strategy/type/phone [post]
func (ctrl *CertiController) insertPhoneCertificateStrategy(c *gin.Context) {}

//	@Tags			Certi
//	@ID				InsertDodoCertificateStrategy
//	@Summary		Insert dodo Certificate Strategy
//	@Description	Insert dodo Certificate Strategy
//	@security		ApiKeyAuth
//	@Produce		json
//	@Param			insert_certificate_strategy_req	body		[]certificate.DodoCertificateInsertPart	true	"insert_certificate_strategy_req"
//	@Success		200								{object}	certificate.CertificateStrategy
//	@Failure		400								{object}	appService_errors.RainbowAppServiceErrorDetailInfo	"Invalid request"
//	@Failure		500								{object}	appService_errors.RainbowAppServiceErrorDetailInfo	"Internal Server error"
//	@Router			/certis/strategy/type/dodo [post]
func (ctrl *CertiController) insertDodoCertificateStrategy(c *gin.Context) {}

//	@Tags			Certi
//	@ID				InsertContractCertificateStrategy
//	@Summary		Insert contract Certificate Strategy
//	@Description	Insert contract Certificate Strategy
//	@security		ApiKeyAuth
//	@Produce		json
//	@Param			insert_certificate_strategy_req	body		[]certificate.ContractCertificateInsertPart	true	"insert_certificate_strategy_req"
//	@Success		200								{object}	certificate.CertificateStrategy
//	@Failure		400								{object}	appService_errors.RainbowAppServiceErrorDetailInfo	"Invalid request"
//	@Failure		500								{object}	appService_errors.RainbowAppServiceErrorDetailInfo	"Internal Server error"
//	@Router			/certis/strategy/type/contract [post]
func (ctrl *CertiController) insertContractCertificateStrategy(c *gin.Context) {}

//	@Tags			Certi
//	@ID				InsertGaslessCertificateStrategy
//	@Summary		Insert gasless Certificate Strategy
//	@Description	Insert gasless Certificate Strategy
//	@security		ApiKeyAuth
//	@Produce		json
//	@Param			insert_certificate_strategy_req	body		[]certificate.GaslessCertificateInsertPart	true	"insert_certificate_strategy_req"
//	@Success		200								{object}	certificate.CertificateStrategy
//	@Failure		400								{object}	appService_errors.RainbowAppServiceErrorDetailInfo	"Invalid request"
//	@Failure		500								{object}	appService_errors.RainbowAppServiceErrorDetailInfo	"Internal Server error"
//	@Router			/certis/strategy/type/gasless [post]
func (ctrl *CertiController) insertGaslessCertificateStrategy(c *gin.Context) {}
