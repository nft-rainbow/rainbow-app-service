package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	"github.com/nft-rainbow/discordbot-service/models"
	"github.com/nft-rainbow/discordbot-service/routers"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
)
var s *discordgo.Session

func initConfig() {
	viper.SetConfigName("config")             // name of config file (without extension)
	viper.SetConfigType("yaml")               // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")                  // optionally look for config in the working directory
	err := viper.ReadInConfig()               // Find and read the config file
	if err != nil {                           // Handle errors reading the config file
		log.Fatalln(fmt.Errorf("fatal error config file: %w", err))
	}
}

func init() {
	initConfig()
	var err error
	s, err = discordgo.New("Bot " + viper.GetString("botToken"))
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}
}

func initGin() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Logger())

	return engine
}

func main() {
	models.ConnectDB()

	app := initGin()
	routers.SetupRoutes(app)

	port := viper.GetString("port")
	if port == "" {
		logrus.Panic("port must be specified")
	}

	address := fmt.Sprintf("0.0.0.0:%s", port)
	logrus.Info("Discord-Bot-Service Start Listening and serving HTTP on ", address)
	err := app.Run(address)
	if err != nil {
		log.Panic(err)
	}
}











