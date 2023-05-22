package config

import (
	"fmt"
	"log"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

var (
	_config Config
)

func Init() {
	initViper()
	if err := viper.Unmarshal(&_config, func(d *mapstructure.DecoderConfig) {
		d.ErrorUnset = true
	}); err != nil {
		panic(err)
	}
}

func initViper() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	viper.AddConfigPath("..")     // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		log.Fatalln(fmt.Errorf("fatal error config file: %w", err))
	}
}

type Config struct {
	Port  int `yaml:"port"`
	Mysql struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Db       string `yaml:"db"`
	} `yaml:"mysql"`
	DiscordBotToken string `yaml:"discordBotToken"`
	DodoBot         struct {
		ClientID  int    `yaml:"clientId"`
		TokenID   string `yaml:"tokenId"`
		InviteURL string `yaml:"inviteUrl"`
	} `yaml:"dodoBot"`
	JwtKeys struct {
		Openapi   string `yaml:"openapi"`
		Dashboard string `yaml:"dashboard"`
	} `yaml:"jwtKeys"`
	CustomMint struct {
		MintRespPrefix string `yaml:"mintRespPrefix"`
	} `yaml:"customMint"`
	Proxy               string `yaml:"proxy"`
	Advertise           string `yaml:"advertise"`
	RainbowDashboardAPI string `yaml:"rainbowDashboardApi"`
	RainbowOpenAPI      string `yaml:"rainbowOpenApi"`
	Env                 string `yaml:"env"`
	Rpcs                struct {
		Mainnet string `yaml:"mainnet"`
		Testnet string `yaml:"testnet"`
	} `yaml:"rpcs"`
	StartTime int `yaml:"startTime"`
	ImagesDir struct {
		Minted    string `yaml:"minted"`
		NonMinted string `yaml:"nonMinted"`
	} `yaml:"imagesDir"`
	PosterDir struct {
		Activity string `yaml:"activity"`
		Result   string `yaml:"result"`
	} `yaml:"posterDir"`
	Oss struct {
		Endpoint        string `yaml:"endpoint"`
		AccessKeyID     string `yaml:"accessKeyId"`
		AccessKeySecret string `yaml:"accessKeySecret"`
		BucketName      string `yaml:"bucketName"`
	} `yaml:"oss"`
	URL struct {
		Activity string `yaml:"activity"`
	} `yaml:"url"`
	IPLimitEveryday int `yaml:"ipLimitEveryday"`
	Anyweb          struct {
		Appid  string `yaml:"appid"`
		Secret string `yaml:"secret"`
	} `yaml:"anyweb"`
	Log struct {
		Level  string `yaml:"level"`
		Folder string `yaml:"folder"`
	} `yaml:"log"`

	Gasless struct {
		UserID        uint `yaml:"userId"`
		AppID         uint `yaml:"appId"`
		ContractRawID struct {
			Mainnet uint `yaml:"mainnet"`
			Testnet uint `yaml:"testnet"`
		} `yaml:"contractRawId"`
	} `yaml:"gasless"`
}

func GetConfig() *Config {
	return &_config
}
