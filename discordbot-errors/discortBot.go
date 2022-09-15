package discordbot_errors

import (
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

type DiscordBotError int

type DiscordBotErrorInfo struct {
	Message        string
	HttpStatusCode int
}

type DiscordBotErrorDetailInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var discordBotErrorInfos = make(map[DiscordBotError]DiscordBotErrorInfo)

func (r DiscordBotError) HttpStatusCode() int {
	return discordBotErrorInfos[r].HttpStatusCode
}

func (r DiscordBotError) Error() string {
	return discordBotErrorInfos[r].Message
}

func (r DiscordBotError) RenderJSON(c *gin.Context) {
	httpStatusCode := discordBotErrorInfos[r].HttpStatusCode
	c.JSON(httpStatusCode, r.ErrorResponse())
}

func (r DiscordBotError) AbortWithRenderJSON(c *gin.Context) {
	debug.PrintStack()
	c.Abort()
	r.RenderJSON(c)
}

func (r DiscordBotError) ErrorResponse() *DiscordBotErrorDetailInfo {
	return &DiscordBotErrorDetailInfo{
		Code:    int(r),
		Message: r.Error(),
	}
}
