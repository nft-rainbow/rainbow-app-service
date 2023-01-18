package middlewares

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nft-rainbow/rainbow-app-service/utils"
	"github.com/spf13/viper"
)

var (
	ipCounter      sync.Map
	ipCntResetOnce sync.Once
	limitCount     = viper.GetInt("newYearEvent.ipLimitEveryday")
)

type any = interface{}

func IpLimitMiddleware() gin.HandlerFunc {
	ipCntResetOnce.Do(func() {
		go loopResetIpCounter()
	})
	return func(c *gin.Context) {
		ip := c.RemoteIP()
		v, _ := ipCounter.LoadOrStore(ip, 0)
		ipCounter.Store(ip, v.(int)+1)
		if v.(int) >= limitCount {
			c.JSON(http.StatusTooManyRequests, map[string]string{"message": "too many requests"})
			c.Abort()
			return
		}
	}
}

// reset at 0 o'clock everyday
func loopResetIpCounter() {
	for {
		resetIpCounter()
		tommorow := utils.TomorrowBegin()
		dur := time.Until(tommorow)
		log.Print("settle cost after ", dur)
		<-time.After(dur)
	}
}

func resetIpCounter() {
	ipCounter.Range(func(k any, v any) bool {
		ipCounter.Store(k, 0)
		return true
	})
}
