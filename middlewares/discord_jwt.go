package middlewares

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	discordbot_errors "github.com/nft-rainbow/discordbot-service/discordbot-errors"
	"github.com/nft-rainbow/discordbot-service/models"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"runtime/debug"
	"time"
)
var (
	JwtAuthMiddleware *jwt.GinJWTMiddleware
)

var JwtIdentityKey = "user_id"

func InitDiscordJwtMiddleware() {
	var err error
	JwtAuthMiddleware, err = jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "rainbow-app-jwt",
		Key:         []byte(viper.GetString("jwtKeys.discord")),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour * 24,
		IdentityKey: JwtIdentityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.MintReq); ok {
				return jwt.MapClaims{
					JwtIdentityKey: v.UserID,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			id := claims[JwtIdentityKey]
			return id
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var mintConfig models.MintReq
			if err := c.ShouldBind(&mintConfig); err != nil {
				return "", discordbot_errors.ERR_INVALID_REQUEST_COMMON
			}

			resp, err := models.FindBindingActivityConfigByChannelId(mintConfig.ChannelID)
			if err != nil {
				return nil, jwt.ErrFailedAuthentication
			}

			return resp, nil
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
			c.JSON(code, gin.H{
				"code":    discordbot_errors.ERR_AUTHORIZATION_JWT,
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization", // cookie: jwt, query: token
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	if err != nil {
		logrus.WithError(err).WithField("stack", string(debug.Stack())).Fatal("init DashboardJWT middleware error")
		return
	}

	logrus.Info("init rainbow app jwt middleware done")

}

func LoginHandler(c *gin.Context) {
	JwtAuthMiddleware.LoginHandler(c)
}
