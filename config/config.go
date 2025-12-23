package config

import (
	"fmt"
	"log"

	"github.com/mitchellh/mapstructure"
	"github.com/nft-rainbow/rainbow-app-service/models/enums"
	"github.com/spf13/viper"
)

var (
	_config Config
)

func Init() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	viper.AddConfigPath("..")     // optionally look for config in the working directory
	loadViper()
}

func InitByFile(configPath string) {
	viper.SetConfigFile(configPath)
	loadViper()
}

func loadViper() {
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Fatalln(fmt.Errorf("fatal error config file: %w", err))
	}
	fmt.Printf("viper user config file: %v\n", viper.ConfigFileUsed())
	if err := viper.Unmarshal(&_config, func(dc *mapstructure.DecoderConfig) {
		dc.ErrorUnset = true
	}); err != nil {
		panic(err)
	}
}

// func initViper() {
// 	viper.SetConfigName("config") // name of config file (without extension)
// 	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
// 	viper.AddConfigPath(".")      // optionally look for config in the working directory
// 	viper.AddConfigPath("..")     // optionally look for config in the working directory
// 	err := viper.ReadInConfig()   // Find and read the config file
// 	if err != nil {               // Handle errors reading the config file
// 		log.Fatalln(fmt.Errorf("fatal error config file: %w", err))
// 	}
// }

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
	BotEnable       bool   `yaml:"botEnable"`
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
	Storage struct {
		Base              string `yaml:"base"`
		BatchMintRequests string `yaml:"batchMintRequests"`
	} `yaml:"storage"`
	URL struct {
		Activity string `yaml:"activity"`
	} `yaml:"url"`
	IPLimitEveryday int `yaml:"ipLimitEveryday"`
	Wallet          struct {
		Anyweb struct {
			Appid  string `yaml:"appid"`
			Secret string `yaml:"secret"`
		} `yaml:"anyweb"`
		Cellar struct {
			Mainnet Cellar `yaml:"mainnet"`
			Testnet Cellar `yaml:"testnet"`
		}
	} `yaml:"wallet"`
	Log struct {
		Level  string `yaml:"level"`
		Folder string `yaml:"folder"`
	} `yaml:"log"`

	Gasless struct {
		UserID        uint `yaml:"userId"`
		AppID         uint `yaml:"appId"`
		MaxAmount     uint `yaml:"maxAmount"`
		ContractRawID struct {
			Mainnet uint `yaml:"mainnet"`
			Testnet uint `yaml:"testnet"`
		} `yaml:"contractRawId"`
	} `yaml:"gasless"`
}

type Cellar struct {
	Appid string `yaml:"appid"`
	Host  string `yaml:"host"`
}

func GetConfig() *Config {
	return &_config
}

func GetGaslessContractIdByChain(chain enums.Chain) uint {
	contractRawId := _config.Gasless.ContractRawID.Testnet
	if chain == enums.CHAIN_CONFLUX {
		contractRawId = _config.Gasless.ContractRawID.Mainnet
	}
	return contractRawId
}

func GetCellarByChain(chain enums.Chain) (*Cellar, error) {
	switch chain {
	case enums.CHAIN_CONFLUX:
		return &_config.Wallet.Cellar.Mainnet, nil
	case enums.CHAIN_CONFLUX_TEST:
		return &_config.Wallet.Cellar.Testnet, nil
	default:
		return nil, fmt.Errorf("unknown chain: %v", chain)
	}
}
