package ginutils

import (
	"fmt"
	"net/http"

	appService_errors "github.com/nft-rainbow/rainbow-app-service/appService-errors"
	"github.com/pkg/errors"

	"github.com/gin-gonic/gin"
)

type CommonMessage struct {
	Message string `json:"message"`
}

var (
	CommonSuccessMessage = CommonMessage{Message: "Success"}
)

func DataResponse(data interface{}) interface{} {
	return data
}

func ErrorResponse(code int, err error) gin.H {
	return gin.H{
		"code":    code,
		"message": err.Error(),
	}
}

func RenderResp(c *gin.Context, data interface{}, err error) {
	if err != nil {
		RenderRespError(c, err)
		return
	}
	RenderRespOK(c, data)
}

func RenderRespOK(c *gin.Context, data interface{}, httpStatusCode ...int) {
	statusCode := http.StatusOK

	if len(httpStatusCode) > 0 {
		statusCode = httpStatusCode[0]
	}
	c.JSON(statusCode, DataResponse(data))
}

// rainbowErrorCode 有值时，message 为 err message，如果 err 为 rainbow error, 则 status code 与 code 都来自 err, 否则来自rainbowErrorCode
// 否则 message 为 err message，status code 与 code 为 ERR_INTERNAL_SERVER_COMMON
func RenderRespError(c *gin.Context, err error, rainbowErrorCode ...appService_errors.RainbowAppServiceError) {
	c.Error(err)
	c.Set("error_stack", fmt.Sprintf("%+v", errors.WithStack(err)))

	if re, ok := err.(appService_errors.RainbowAppServiceError); ok {
		re.RenderJSON(c)
		return
	}

	_rainbowErrorCode := appService_errors.ERR_INTERNAL_SERVER_COMMON

	if len(rainbowErrorCode) > 0 {
		_rainbowErrorCode = rainbowErrorCode[0]
	}

	c.JSON(_rainbowErrorCode.HttpStatusCode(), ErrorResponse(int(_rainbowErrorCode), err))
}
