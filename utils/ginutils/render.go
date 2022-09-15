package ginutils

import (
	discordbot_errors "github.com/nft-rainbow/discordbot-service/discordbot-errors"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
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
func RenderRespError(c *gin.Context, err error, rainbowErrorCode ...discordbot_errors.DiscordBotError) {
	c.Error(err)
	c.Set("error_stack", string(debug.Stack()))

	if re, ok := err.(discordbot_errors.DiscordBotError); ok {
		re.RenderJSON(c)
		return
	}

	_rainbowErrorCode := discordbot_errors.ERR_INTERNAL_SERVER_COMMON

	if len(rainbowErrorCode) > 0 {
		_rainbowErrorCode = rainbowErrorCode[0]
	}

	c.JSON(_rainbowErrorCode.HttpStatusCode(), ErrorResponse(int(_rainbowErrorCode), err))
}
