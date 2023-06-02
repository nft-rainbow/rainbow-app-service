package services

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestUploadFileToOss(t *testing.T) {
	bucket, err := getOSSBucket(viper.GetString("oss.bucketName"))
	assert.NoError(t, err)

	activityCode := "99b5b77c"
	_path := path.Join(viper.GetString("posterDir.activity"), activityCode+".png")
	err = bucket.PutObjectFromFile(_path, "/Users/dayong/tmp/abc.png")
	assert.NoError(t, err)
}

func TestGenerateActivityPoster(t *testing.T) {
	buf, err := drawPoster("/Users/dayong/myspace/mywork/rainbow-app-service/assets/images/activityPoster.png",
		"/Users/dayong/myspace/mywork/rainbow-app-service/assets/fonts/PingFang.ttf",
		"99b5b77c",
		"https://nftrainbow.oss-cn-hangzhou.aliyuncs.com/events/YUANLONGYATU/POAP/yuanlong-poster.jpeg",
		"UOVAMETA NO.022数字纪念徽章",
		"UOVAMETA NO.022数字纪念徽章",
		1685030400,
		-1,
	)
	assert.NoError(t, err)
	err = ioutil.WriteFile("/Users/dayong/tmp/abc.png", buf.Bytes(), 0777)
	assert.NoError(t, err)
}

func init() {
	initConfig()
}

func initConfig() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	// viper.AddConfigPath(".")      // optionally look for config in the working directory
	viper.AddConfigPath("..")   // optionally look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Fatalln(fmt.Errorf("fatal error config file: %w", err))
	}
}
