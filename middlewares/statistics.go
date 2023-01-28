package middlewares

import (
	"github.com/dodo-open/dodo-open-go/log"
	"github.com/gin-gonic/gin"
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/utils"
)

func Statistic() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if c.IsAborted() {
			return
		}
		err := models.IncreaseStatistic(c.Request.Method, c.FullPath(), c.ClientIP(), utils.TodayDateStr())
		if err != nil {
			log.Errorf("static error: %v", err)
		}
	}
}
