package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	"github.com/nft-rainbow/rainbow-app-service/middlewares"
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/routers"
	"github.com/nft-rainbow/rainbow-app-service/services"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
)


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
	middlewares.InitRainbowJwtMiddleware()
}

func initGin() {
	engine := gin.New()
	engine.Use(gin.Logger())
	routers.SetupRoutes(engine)

	port := viper.GetString("port")
	if port == "" {
		logrus.Panic("port must be specified")
	}

	address := fmt.Sprintf("0.0.0.0:%s", port)
	logrus.Info("Rainbow-App-Service Start Listening and serving HTTP on ", address)
	err := engine.Run(address)
	if err != nil {
		log.Panic(err)
	}
}

func initDiscordBot() {
	var err error
	s := services.InitSession()

	err = s.Open()
	if err != nil {
		panic(err)
	}


	log.Println("Adding commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(services.Commands))
	for i, v := range services.Commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, "", v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}

	defer s.Close()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop
}

func main() {
	models.ConnectDB()
	go initGin()

	initDiscordBot()
}









