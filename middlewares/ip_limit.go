package middlewares

import (
	"log"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	appService_errors "github.com/nft-rainbow/rainbow-app-service/appService-errors"
	"github.com/nft-rainbow/rainbow-app-service/utils"
	"github.com/nft-rainbow/rainbow-app-service/utils/ginutils"
	"github.com/spf13/viper"
)

var (
	ipCounter      sync.Map
	ipCntResetOnce sync.Once
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
		if v.(int) >= viper.GetInt("newYearEvent.ipLimitEveryday") {
			ginutils.RenderRespError(c, appService_errors.ERR_TOO_MANY_REQUEST_COMMON)
			c.Abort()
			return
		}
	}
}

// reset at 0 o'clock everyday
func loopResetIpCounter() {
	log.Print("ip limit everyday", viper.GetInt("newYearEvent.ipLimitEveryday"))
	for {
		resetIpCounter()
		tommorow := utils.TomorrowBegin()
		dur := time.Until(tommorow)
		log.Print("reset ip counter after ", dur)
		<-time.After(dur)
	}
}

func resetIpCounter() {
	ipCounter.Range(func(k any, v any) bool {
		ipCounter.Store(k, 0)
		return true
	})
}
