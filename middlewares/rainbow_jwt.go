package middlewares

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strings"

	jwt "github.com/appleboy/gin-jwt/v2"

	"net/http"
	"runtime/debug"
	"time"

	"github.com/gin-gonic/gin"
	appService_errors "github.com/nft-rainbow/rainbow-app-service/appService-errors"
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/utils/ginutils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	OpenJwtAuthMiddleware *jwt.GinJWTMiddleware
)

const (
	KYCTypeKey   = "KYCType"
	AppUserIdKey = "AppUserId"
)

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
		IdentityKey: JwtIdentityKey,
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
			id := claims[JwtIdentityKey]
			return uint(id.(float64))
		},
		LoginResponse: func(c *gin.Context, code int, message string, time time.Time) {
			ginutils.RenderRespOK(c, gin.H{
				"token":  message,
				"expire": time,
			}, code)
		},
		RefreshResponse: func(c *gin.Context, code int, message string, time time.Time) {
			ginutils.RenderRespOK(c, gin.H{
				"token":  message,
				"expire": time,
			}, code)
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			ginutils.RenderRespError(c, errors.New(message), appService_errors.RainbowAppServiceError(appService_errors.GetAppServiceOthersErrCode(code)))
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

func GenerateDiscordOpenJWT(channelId string) (string, error) {
	activity, err := models.FindDiscordCustomActivityConfigByChannelId(channelId)
	if err != nil {
		return "", err
	}
	config, err := models.FindDiscordConfigById(int(activity.AppId))
	if err != nil {
		return "", err
	}

	kycType, err := getKycType(config.RainbowUserId)
	if err != nil {
		return "", err
	}

	data := &App{
		Id:        uint(config.AppId),
		KycType:   kycType,
		AppUserId: uint(config.RainbowUserId),
	}

	tokenString, _, err := OpenJwtAuthMiddleware.TokenGenerator(data)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GenDiscordOpenJWTByRainbowUserId(id uint) (string, error) {
	config, err := models.FindDiscordConfigByUserId(int(id))
	if err != nil {
		return "", err
	}

	kycType, err := getKycType(int32(id))
	if err != nil {
		return "", err
	}

	data := &App{
		Id:        uint(config.AppId),
		KycType:   kycType,
		AppUserId: uint(config.RainbowUserId),
	}

	tokenString, _, err := OpenJwtAuthMiddleware.TokenGenerator(data)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GenerateDoDoOpenJWT(channelId string) (string, error) {
	activity, err := models.FindDoDoCustomActivityConfigByChannelId(channelId)
	if err != nil {
		return "", err
	}
	config, err := models.FindDoDoConfigById(int(activity.AppId))
	if err != nil {
		return "", err
	}

	kycType, err := getKycType(config.RainbowUserId)
	if err != nil {
		return "", err
	}

	data := &App{
		Id:        uint(config.AppId),
		KycType:   kycType,
		AppUserId: uint(config.RainbowUserId),
	}

	tokenString, _, err := OpenJwtAuthMiddleware.TokenGenerator(data)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GenDoDoOpenJWTByRainbowUserId(id uint) (string, error) {
	config, err := models.FindDoDoConfigByUserId(int(id))
	if err != nil {
		return "", err
	}

	kycType, err := getKycType(config.RainbowUserId)
	if err != nil {
		return "", err
	}

	data := &App{
		Id:        uint(config.AppId),
		KycType:   kycType,
		AppUserId: uint(config.RainbowUserId),
	}

	tokenString, _, err := OpenJwtAuthMiddleware.TokenGenerator(data)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GeneratePOAPOpenJWT(userId, appId int32) (string, error) {
	kycType, err := getKycType(userId)
	if err != nil {
		return "", err
	}
	data := &App{
		Id:        uint(appId),
		KycType:   kycType,
		AppUserId: uint(userId),
	}

	tokenString, _, err := OpenJwtAuthMiddleware.TokenGenerator(data)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GenPOAPOpenJWTByRainbowUserId(activity models.POAPActivityConfig) (string, error) {
	kycType, err := getKycType(activity.RainbowUserId)
	if err != nil {
		return "", err
	}

	data := &App{
		Id:        uint(activity.AppId),
		KycType:   kycType,
		AppUserId: uint(activity.RainbowUserId),
	}

	tokenString, _, err := OpenJwtAuthMiddleware.TokenGenerator(data)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func PrefixToken(token string) string {
	if strings.HasPrefix(token, "Bearer ") {
		return token
	}
	return "Bearer " + token
}

func GenOpenJWTByRainbowUserId(rainbowUserId, appId int32) (string, error) {
	kycType, err := getKycType(rainbowUserId)
	if err != nil {
		return "", err
	}

	data := &App{
		Id:        uint(appId),
		KycType:   kycType,
		AppUserId: uint(rainbowUserId),
	}

	tokenString, _, err := OpenJwtAuthMiddleware.TokenGenerator(data)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func getKycType(userId int32) (uint, error) {
	user := &User{
		Id: uint(userId),
	}

	tokenString, _, err := JwtAuthMiddleware.TokenGenerator(user)
	if err != nil {
		return 0, err
	}
	kycType, err := queryKycType(tokenString)
	if err != nil {
		return 0, err
	}
	return uint(kycType), nil
}

func queryKycType(tokenString string) (float64, error) {
	req, err := http.NewRequest("GET", viper.GetString("rainbowDashboardApi")+"/dashboard/users/profile", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", PrefixToken(tokenString))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	t := make(map[string]interface{})
	err = json.Unmarshal(content, &t)
	if err != nil {
		return 0, err
	}
	if t["code"] != nil {
		return 0, errors.New(t["message"].(string))
	}

	return t["type"].(float64), nil
}
