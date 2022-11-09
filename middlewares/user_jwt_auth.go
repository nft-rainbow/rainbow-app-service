package middlewares

import (
	"errors"
	"github.com/nft-rainbow/rainbow-app-service/utils/ginutils"
	"runtime/debug"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	appService_errors "github.com/nft-rainbow/rainbow-app-service/appService-errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	JwtAuthMiddleware *jwt.GinJWTMiddleware
)

type User struct {
	Id    uint
	Email string
	Name  string
}

var JwtIdentityKey = "id"

func InitDashboardJwtMiddleware() {
	var err error
	JwtAuthMiddleware, err = jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "rainbow-api-jwt",
		Key:         []byte(viper.GetString("jwtKeys.dashboard")),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour * 24,
		IdentityKey: JwtIdentityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					JwtIdentityKey: v.Id,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			id := claims[JwtIdentityKey]
			return uint(id.(float64))
		},
		LoginResponse: func(c *gin.Context, code int, message string, time time.Time) {
			c.JSON(code, gin.H{
				"token":  message,
				"expire": time,
			})
		},
		RefreshResponse: func(c *gin.Context, code int, message string, time time.Time) {
			c.JSON(code, gin.H{
				"token":  message,
				"expire": time,
			})
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			ginutils.RenderRespError(c, errors.New(message), appService_errors.RainbowAppServiceError(appService_errors.GetAppServiceOthersErrCode(code)))
		},
		TokenLookup:   "header: Authorization", // cookie: jwt, query: token
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	if err != nil {
		logrus.WithError(err).WithField("stack", string(debug.Stack())).Fatal("init DashboardJWT middleware error")
		return
	}

	logrus.Info("init dashboard jwt middleware done")

}
