package clients

import (
	"fmt"
	"log"
	"testing"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("..")     // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		log.Fatalln(fmt.Errorf("fatal error config file: %w", err))
	}
}

func TestGetAnywebAccessToken(t *testing.T) {
	code := "05e98a4f-074c-4b52-93b2-3c77e3402b93"
	anywebAccessToken, err := GetAnywebAccessToken(code)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(anywebAccessToken)

	/* refreshToken, err := GetRefreshToken(anywebAccessToken.RefreshToken)
	if err != nil {
		t.Error(err)
	}
	t.Log(refreshToken) */

	userInfo, err := GetUserInfo(anywebAccessToken.AccessToken, anywebAccessToken.UnionId, []string{"baseInfo"})
	if err != nil {
		t.Error(err)
	}
	t.Log(userInfo)
}
