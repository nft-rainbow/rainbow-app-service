package appService_errors

import (
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

type RainbowAppServiceError int

type RainbowAppServiceErrorInfo struct {
	Message        string
	HttpStatusCode int
}

type RainbowAppServiceErrorDetailInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var RainbowAppServiceErrorInfos = make(map[RainbowAppServiceError]RainbowAppServiceErrorInfo)

func (r RainbowAppServiceError) HttpStatusCode() int {
	return RainbowAppServiceErrorInfos[r].HttpStatusCode
}

func (r RainbowAppServiceError) Error() string {
	return RainbowAppServiceErrorInfos[r].Message
}

func (r RainbowAppServiceError) RenderJSON(c *gin.Context) {
	httpStatusCode := RainbowAppServiceErrorInfos[r].HttpStatusCode
	c.JSON(httpStatusCode, r.ErrorResponse())
}

func (r RainbowAppServiceError) AbortWithRenderJSON(c *gin.Context) {
	debug.PrintStack()
	c.Abort()
	r.RenderJSON(c)
}

func (r RainbowAppServiceError) ErrorResponse() *RainbowAppServiceErrorDetailInfo {
	return &RainbowAppServiceErrorDetailInfo{
		Code:    int(r),
		Message: r.Error(),
	}
}
