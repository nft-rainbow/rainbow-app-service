package middlewares

import (
	"errors"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	discordbot_errors "github.com/nft-rainbow/discordbot-service/discordbot-errors"
	"github.com/nft-rainbow/discordbot-service/utils/ginutils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"runtime/debug"
	"time"
)

var (
	OpenJwtAuthMiddleware *jwt.GinJWTMiddleware
)

const (
	KYCTypeKey   = "KYCType"
	AppUserIdKey = "AppUserId"
)

var OpenJwtIdentityKey = "id"

type appLogin struct {
	AppId     string `form:"app_id" json:"app_id" binding:"required"`
	AppSecret string `form:"app_secret" json:"app_secret" binding:"required"`
}

type App struct {
	Id        uint
	AppId     string
	KycType   uint
	AppUserId uint
}

func InitRainbowJwtMiddleware() {
	// Set jwt timeout to one month if environment is development mode for easy testing
	timeout := time.Hour * 24 * 30
	if viper.GetString("environment") == "production" {
		timeout = time.Hour
	}
	var err error
	OpenJwtAuthMiddleware, err = jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "Rainbow-openapi",
		Key:         []byte(viper.GetString("jwtKeys.openapi")),
		Timeout:     timeout,
		MaxRefresh:  time.Hour * 5,
		IdentityKey: OpenJwtIdentityKey,
		LoginResponse: func(c *gin.Context, code int, message string, time time.Time) {
			ginutils.RenderRespOK(c, gin.H{
				"token":  message,
				"expire": time,
			}, code)
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*App); ok {
				return jwt.MapClaims{
					JwtIdentityKey: v.Id,
					KYCTypeKey:     v.KycType,
					AppUserIdKey:   v.AppUserId,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			id := claims[OpenJwtIdentityKey]
			return uint(id.(float64))
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals appLogin
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			appId := loginVals.AppId

			return &App{
				AppId:     appId,
			}, nil
		},
		RefreshResponse: func(c *gin.Context, code int, message string, time time.Time) {
			ginutils.RenderRespOK(c, gin.H{
				"token":  message,
				"expire": time,
			}, code)
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			ginutils.RenderRespError(c, errors.New(message), discordbot_errors.DiscordBotError(discordbot_errors.GetDiscordBotOthersErrCode(code)))
		},
		TokenLookup:   "header: Authorization",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	if err != nil {
		logrus.WithError(err).WithField("stack", string(debug.Stack())).Fatal("init OpenJwt middleware error")
		return
	}

	logrus.Info("init open jwt middleware done")
}