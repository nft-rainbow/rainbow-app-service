package services

import (
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/spf13/viper"
	"testing"
	"time"
)

func TestAddLogoAndUpload(t *testing.T) {
	initConfig()
	err := AddLogoAndUpload("https://nft-rainbow.oss-cn-hangzhou.aliyuncs.com/rabbit-poap/e1/3.png", "3.png", "12356")
	if err != nil {
		panic(err)
	}

}

func TestGenerateActivityPoster(t *testing.T) {
	GenerateActivityPoster(
		"https://nft-rainbow.oss-cn-hangzhou.aliyuncs.com/rabbit-poap/e1/3.png",
		"尤伦斯双年展纪念徽章",
		"尤伦斯200位常驻艺术家双年展，由200名艺术家共同完成徽章设计 ",
		time.Now().Format("2006-01-02"),
		time.Now().Format("2006-01-02"),
		"https://nft-rainbow.oss-cn-hangzhou.aliyuncs.com/rabbit-poap/e1/3.png",
	)
}

func TestGenerateResultPoster(t *testing.T) {
	GenerateResultPoster(
		models.POAPResult{
			Address: "cfx:aar9up0wsbgtw7f0g5tyc4hbwb2wa5wf7eab4t69tk",
			FileURL: "https://nft-rainbow.oss-cn-hangzhou.aliyuncs.com/rabbit-poap/e1/3.png",
			TokenID: "12312",
		},
		"尤伦斯双年展纪念徽章",
		"https://nft-rainbow.oss-cn-hangzhou.aliyuncs.com/rabbit-poap/e1/3.png",
	)
}

func initConfig() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	viper.AddConfigPath("../")    // optionally look for config in the working directory
	_ = viper.ReadInConfig()      // Find and read the config file
}
