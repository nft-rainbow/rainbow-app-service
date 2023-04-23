package services

import (
	"fmt"
	"log"
	"path"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func _TestUploadFileToOss(t *testing.T) {
	bucket, err := getOSSBucket(viper.GetString("oss.bucketName"))
	assert.NoError(t, err)

	activityCode := "3eb6e93b"
	_path := path.Join(viper.GetString("posterDir.activity"), activityCode+".png")
	err = bucket.PutObjectFromFile(_path, "/Users/dayong/Downloads/0fcccd86.png")
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
