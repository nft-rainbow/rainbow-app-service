package middlewares

import (
	"bytes"
	"fmt"

	"github.com/gin-gonic/gin"
	appService_errors "github.com/nft-rainbow/rainbow-app-service/appService-errors"
	"github.com/nft-rainbow/rainbow-app-service/utils/ginutils"

	"github.com/sirupsen/logrus"
)

func Recovery() gin.HandlerFunc {
	var buf bytes.Buffer
	return gin.CustomRecoveryWithWriter(&buf, gin.RecoveryFunc(func(c *gin.Context, err interface{}) {
		defer func() {
			fmt.Println(buf.String())
			logrus.WithField("recovered", buf.String()).WithField("error", err).Error("panic and recovery")
			buf.Reset()
		}()
		ginutils.RenderRespError(c, appService_errors.ERR_INTERNAL_SERVER_COMMON)
		c.Abort()
	}))
}
