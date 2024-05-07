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

//	@Tags			POAP
//	@ID				POAPMintByCSV
//	@Summary		POAP Mint By CSV
//	@Description	POAP Mint By CSV file
//	@security		ApiKeyAuth
//	@Produce		json
//	@Param			poap_csv_mint_dto	body		services.POAPRequest	true	"poap_csv_mint_dto"
//	@Success		200					{object}	rainbowsdk.ModelsMintTask
//	@Failure		400					{object}	appService_errors.RainbowAppServiceErrorDetailInfo	"Invalid request"
//	@Failure		500					{object}	appService_errors.RainbowAppServiceErrorDetailInfo	"Internal Server error"
//	@Router			/poap/csv [post]
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

//	@Tags			POAP
//	@ID				MintByH5
//	@Summary		Mint By H5
//	@Description	Mint By H5
//	@security		ApiKeyAuth
//	@Produce		json
//	@Param			mint_req	body		services.MintReq	true	"mint request"
//	@Success		200			{object}	models.POAPResult
//	@Failure		400			{object}	appService_errors.RainbowAppServiceErrorDetailInfo	"Invalid request"
//	@Failure		500			{object}	appService_errors.RainbowAppServiceErrorDetailInfo	"Internal Server error"
//	@Router			/poap/h5 [post]
func mintByH5(c *gin.Context) {
	var poapRequest services.MintReq
	if err := c.ShouldBind(&poapRequest); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}
	if _, err := cfxaddress.NewFromBase32(poapRequest.UserAddress); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	resp, err := activityService.H5Mint(&poapRequest)
	ginutils.RenderResp(c, resp, err)
}

//	@Tags			POAP
//	@ID				GetActivityDetail
//	@Summary		Get Activity detail
//	@Description	Get Activity detail info
//	@security		ApiKeyAuth
//	@Produce		json
//	@Param			activity_code	path		string	true	"activity_code"
//	@Success		200				{object}	models.Activity
//	@Failure		400				{object}	appService_errors.RainbowAppServiceErrorDetailInfo	"Invalid request"
//	@Failure		500				{object}	appService_errors.RainbowAppServiceErrorDetailInfo	"Internal Server error"
//	@Router			/poap/activity/{activity_code} [get]
func getActivity(c *gin.Context) {
	poapId := c.Param(ACTIVITY_CODE_KEY)
	item, err := models.FindActivityByCode(poapId)
	// if item.NeedCommand() == true {
	// 	item.Command = ""
	// }
	ginutils.RenderResp(c, item, err)
}

//	@Tags			POAP
//	@ID				GetUserActivities
//	@Summary		Get User Activity list
//	@Description	Get User Activity list
//	@security		ApiKeyAuth
//	@Produce		json
//	@Param			Authorization			header		string							true	"Bearer JWT"
//	@Param			page					query		integer							false	"page"
//	@Param			limit					query		integer							false	"limit"
//	@Param			activity_find_condition	query		models.ActivityFindCondition	false	"activity find condition"
//	@Success		200						{object}	models.ActivityQueryResult
//	@Failure		400						{object}	appService_errors.RainbowAppServiceErrorDetailInfo	"Invalid request"
//	@Failure		500						{object}	appService_errors.RainbowAppServiceErrorDetailInfo	"Internal Server error"
//	@Router			/poap/activity [get]
func getUserActivities(c *gin.Context) {
	// pagination, err := GetPagination(c)
	// if err != nil {
	// 	ginutils.RenderRespError(c, appService_errors.ERR_INVALID_PAGINATION)
	// 	return
	// }

	var cond models.ActivityFindCondition
	if err := c.ShouldBindQuery(&cond); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	mints, err := models.FindAndCountActivity(GetIdFromJwt(c), cond)
	ginutils.RenderResp(c, mints, err)
}

//	@Tags			POAP
//	@ID				AddActivity
//	@Summary		Add Activity
//	@Description	Add Activity
//	@security		ApiKeyAuth
//	@Produce		json
//	@Param			Authorization		header		string						true	"Bearer JWT"
//	@Param			activity_request	body		models.ActivityInsertPart	true	"activity config"
//	@Success		200					{object}	models.Activity
//	@Failure		400					{object}	appService_errors.RainbowAppServiceErrorDetailInfo	"Invalid request"
//	@Failure		500					{object}	appService_errors.RainbowAppServiceErrorDetailInfo	"Internal Server error"
//	@Router			/poap/activity [post]
func addActivity(c *gin.Context) {
	var config *models.ActivityInsertPart
	if err := c.ShouldBind(&config); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	resp, err := activityService.InsertActivity(config, GetIdFromJwt(c))
	ginutils.RenderResp(c, resp, err)
}

//	@Tags			POAP
//	@ID				AddActivityNftConfigs
//	@Summary		Add Activity NFT configs
//	@Description	Add Activity NFT configs
//	@security		ApiKeyAuth
//	@Produce		json
//	@Param			Authorization	header		string							true	"Bearer JWT"
//	@Param			nft_configs		body		[]models.NftConfigUpdatePart	true	"activity nft configs"
//	@Success		200				{array}		models.NFTConfig
//	@Failure		400				{object}	appService_errors.RainbowAppServiceErrorDetailInfo	"Invalid request"
//	@Failure		500				{object}	appService_errors.RainbowAppServiceErrorDetailInfo	"Internal Server error"
//	@Router			/poap/activity/{activity_code}/nftconfigs [post]
func addActivityNftConfigs(c *gin.Context) {
	var nftConfigs []*models.NftConfigUpdatePart
	if err := c.ShouldBind(&nftConfigs); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	activityCode := c.Param(ACTIVITY_CODE_KEY)

	resp, err := activityService.AddActivityNftConfigs(activityCode, nftConfigs, GetIdFromJwt(c))
	ginutils.RenderResp(c, resp, err)
}

//	@Tags			POAP
//	@ID				SetActivityH5Config
//	@Summary		Set H5 Config
//	@Description	Set H5 Page Config
//	@security		ApiKeyAuth
//	@Produce		json
//	@Param			Authorization	header		string			true	"Bearer JWT"
//	@Param			poap_h5_config	body		models.H5Config	true	"poap_h5_config"
//	@Success		200				{object}	models.H5Config
//	@Failure		400				{object}	appService_errors.RainbowAppServiceErrorDetailInfo	"Invalid request"
//	@Failure		500				{object}	appService_errors.RainbowAppServiceErrorDetailInfo	"Internal Server error"
//	@Router			/poap/activity/h5 [post]
func setActivityH5Config(c *gin.Context) {
	var config *models.H5Config
	if err := c.ShouldBind(&config); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	resp, err := activityService.POAPH5Config(config)
	ginutils.RenderResp(c, resp, err)
}

//	@Tags			POAP
//	@ID				UpdateActivity
//	@Summary		Update Activity
//	@Description	Update Activity
//	@security		ApiKeyAuth
//	@Produce		json
//	@Param			Authorization			header		string							true	"Bearer JWT"
//	@Param			activity_id				path		string							true	"activity_id"
//	@Param			update_activity_request	body		models.ActivityUpdateBasePart	true	"update activity request"
//	@Success		200						{object}	models.Activity
//	@Failure		400						{object}	appService_errors.RainbowAppServiceErrorDetailInfo	"Invalid request"
//	@Failure		500						{object}	appService_errors.RainbowAppServiceErrorDetailInfo	"Internal Server error"
//	@Router			/poap/activity/{activity_code} [put]
func updateActivityBase(c *gin.Context) {
	var config models.ActivityUpdateBasePart
	if err := c.ShouldBind(&config); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}
	activityCode := c.Param(ACTIVITY_CODE_KEY)

	resp, err := activityService.UpdateActivityBase(GetIdFromJwt(c), activityCode, &config)
	ginutils.RenderResp(c, resp, err)
}

//	@Tags			POAP
//	@ID				UpdateActivityNftConfig
//	@Summary		Update Activity NFT Config
//	@Description	Update Activity NFT Config
//	@security		ApiKeyAuth
//	@Produce		json
//	@Param			Authorization				header		string						true	"Bearer JWT"
//	@Param			nft_config_id				path		uint						true	"nft_config_id"
//	@Param			update_nft_config_request	body		models.NftConfigUpdatePart	true	"update Nft config request"
//	@Success		200							{object}	models.NFTConfig
//	@Failure		400							{object}	appService_errors.RainbowAppServiceErrorDetailInfo	"Invalid request"
//	@Failure		500							{object}	appService_errors.RainbowAppServiceErrorDetailInfo	"Internal Server error"
//	@Router			/poap/activity/nftconfig/{nft_config_id} [put]
func updateActivityNftConfig(c *gin.Context) {
	req := struct {
		models.NftConfigUpdatePart
		NftConfigId uint `uri:"nft_config_id"`
	}{}

	if err := c.ShouldBindJSON(&req); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	if err := c.ShouldBindUri(&req); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	resp, err := activityService.UpdateNftConfig(GetIdFromJwt(c), req.NftConfigId, &req.NftConfigUpdatePart)
	ginutils.RenderResp(c, resp, err)
}

//	@Tags			POAP
//	@ID				DeleteActivityNftConfig
//	@Summary		Delete Activity NFT Config
//	@Description	Delete Activity NFT Config
//	@security		ApiKeyAuth
//	@Produce		json
//	@Param			Authorization	header		string	true	"Bearer JWT"
//	@Param			nft_config_id	path		uint	true	"nft_config_id"
//	@Success		200				{object}	models.NFTConfig
//	@Failure		400				{object}	appService_errors.RainbowAppServiceErrorDetailInfo	"Invalid request"
//	@Failure		500				{object}	appService_errors.RainbowAppServiceErrorDetailInfo	"Internal Server error"
//	@Router			/post/activity/nftconfig/{nft_config_id} [delete]
func deleteActivityNftConfig(c *gin.Context) {
	req := struct {
		NftConfigId uint `uri:"nft_config_id"`
	}{}

	if err := c.ShouldBindJSON(&req); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	if err := c.ShouldBindUri(&req); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	err := activityService.DeleteActivityNftConfig(GetIdFromJwt(c), req.NftConfigId)
	ginutils.RenderResp(c, ginutils.CommonSuccessMessage, err)
}

//	@Tags			POAP
//	@ID				GetMintResultList
//	@Summary		Get Mint Result list
//	@Description	Get Mint Result list
//	@security		ApiKeyAuth
//	@Produce		json
//	@Param			page			query		integer					false	"page"
//	@Param			limit			query		integer					false	"limit"
//	@Param			filter			query		models.POAPResultFilter	false	"filter"
//	@Param			activity_code	path		string					true	"activity_code"
//	@Success		200				{object}	models.POAPResultQueryResult
//	@Failure		400				{object}	appService_errors.RainbowAppServiceErrorDetailInfo	"Invalid request"
//	@Failure		500				{object}	appService_errors.RainbowAppServiceErrorDetailInfo	"Internal Server error"
//	@Router			/poap/activity/result/{activity_code} [get]
func getMintResultList(c *gin.Context) {
	var pagination models.Pagination
	if err := c.ShouldBindQuery(&pagination); err != nil {
		ginutils.RenderRespError(c, appService_errors.ERR_INVALID_PAGINATION)
		return
	}

	var filter models.POAPResultFilter
	if err := c.ShouldBindQuery(&filter); err != nil {
		ginutils.RenderRespError(c, appService_errors.ERR_INVALID_PAGINATION)
		return
	}

	poapId := c.Param(ACTIVITY_CODE_KEY)
	if poapId == "" {
		ginutils.RenderRespError(c, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	mints, err := models.FindAndCountPOAPResult(poapId, filter, pagination)
	ginutils.RenderResp(c, mints, err)
}

//	@Tags			POAP
//	@ID				GetMintResult
//	@Summary		Get Mint Result
//	@Description	Get Mint Result
//	@security		ApiKeyAuth
//	@Produce		json
//	@Param			activity_code	path		string	true	"activity_code"
//	@Param			id				path		int		true	"id"
//	@Success		200				{object}	models.POAPResult
//	@Failure		400				{object}	appService_errors.RainbowAppServiceErrorDetailInfo	"Invalid request"
//	@Failure		500				{object}	appService_errors.RainbowAppServiceErrorDetailInfo	"Internal Server error"
//	@Router			/poap/activity/result/{activity_code}/{id} [get]
func getMintResultDetail(c *gin.Context) {
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

//	@Tags			POAP
//	@ID				GetMintCount
//	@Summary		Get Mint Count
//	@Description	Get Mint Count
//	@security		ApiKeyAuth
//	@Produce		json
//	@Param			activity_code	path		string	true	"activity_code"
//	@Param			address			path		string	true	"address"
//	@Success		200				{object}	int
//	@Failure		400				{object}	appService_errors.RainbowAppServiceErrorDetailInfo	"Invalid request"
//	@Failure		500				{object}	appService_errors.RainbowAppServiceErrorDetailInfo	"Internal Server error"
//	@Router			/poap/count/{address}/{activity_code} [get]
func getMintCount(c *gin.Context) {
	var err error
	address := c.Param("address")
	activityCode := c.Param(ACTIVITY_CODE_KEY)

	var resp int32
	resp, err = activityService.GetMintCount(activityCode, address)
	ginutils.RenderResp(c, resp, err)
}

//	@Tags			POAP
//	@ID				AddWalletUser
//	@Summary		Add user wallet profile, so backend can get user phone from anyweb
//	@Description	Add user wallet profile, so backend can get user phone from anyweb
//	@security		ApiKeyAuth
//	@Produce		json
//	@Param			add_wallet_user_request	body		services.AddWalletUserReq	true	"add wallet user request"
//	@Success		200						{object}	ginutils.CommonMessage
//	@Failure		400						{object}	appService_errors.RainbowAppServiceErrorDetailInfo	"Invalid request"
//	@Failure		500						{object}	appService_errors.RainbowAppServiceErrorDetailInfo	"Internal Server error"
//	@Router			/poap/wallet/user [POST]
func addWalletUser(c *gin.Context) {
	var config services.AddWalletUserReq
	if err := c.ShouldBind(&config); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	err := walletService.InsertUser(config)
	ginutils.RenderResp(c, ginutils.CommonSuccessMessage, err)
}

//	@Tags			POAP
//	@ID				AddTokenReserves
//	@Summary		Add reserve tokenids for activity
//	@Description	Add reserve tokenids for activity, they will be skiped when mint by h5
//	@security		ApiKeyAuth
//	@Produce		json
//	@Param			token_reserves	body		[]models.TokenReserve	true	"token reserves"
//	@Success		200				{object}	ginutils.CommonMessage
//	@Failure		400				{object}	appService_errors.RainbowAppServiceErrorDetailInfo	"Invalid request"
//	@Failure		500				{object}	appService_errors.RainbowAppServiceErrorDetailInfo	"Internal Server error"
//	@Router			/poap/activity/token-reserve [POST]
func addTokenReserves(c *gin.Context) {
	var trs []*models.TokenReserve
	if err := c.ShouldBindJSON(&trs); err != nil {
		ginutils.RenderRespError(c, err, appService_errors.ERR_INVALID_REQUEST_COMMON)
		return
	}

	err := models.GetDB().Save(trs).Error
	ginutils.RenderResp(c, ginutils.CommonSuccessMessage, err)
}

//	@Tags			POAP
//	@ID				GetTokenReserves
//	@Summary		Get reserve tokenids of activity
//	@Description	Get reserve tokenids of activity
//	@security		ApiKeyAuth
//	@Produce		json
//	@Param			activity_code	path		string	true	"activity code"
//	@Success		200				{array}		models.TokenReserve
//	@Failure		400				{object}	appService_errors.RainbowAppServiceErrorDetailInfo	"Invalid request"
//	@Failure		500				{object}	appService_errors.RainbowAppServiceErrorDetailInfo	"Internal Server error"
//	@Router			/poap/activity/token-reserve/{activity_code} [GET]
func getTokenReserves(c *gin.Context) {
	activityCode := c.Param("activity_code")
	result, err := models.GetResrverTokenIdsByActivityCode(activityCode)
	ginutils.RenderResp(c, result, err)
}
