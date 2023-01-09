package main

import (
	"crypto/tls"
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
	"net/http"
	"net/url"
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
	middlewares.InitDashboardJwtMiddleware()
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
	proxy, _ := url.Parse(viper.GetString("proxy"))

	tr := &http.Transport{
		Proxy:           http.ProxyURL(proxy),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	s.Client.Transport =tr
	s.Dialer.Proxy = http.ProxyURL(proxy)

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

func initDoDoBot(){
	ws := services.InitInstance()
	fmt.Println("Start to connect")

	err := ws.Connect()
	if err != nil {
		panic(err)
	}
	fmt.Println("Start to listen")
	err = ws.Listen()
	if err != nil {
		panic(err)
	}
}

// @title       Rainbow-APP-Service
// @version     1.0
// @description The responses of the open api in swagger focus on the data field rather than the code and the message fields

// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html

// @host     https://console.nftrainbow.cn/apps
// @BasePath /v1
// @schemes  http https
func main() {
	models.ConnectDB()
	go initDoDoBot()
	go initGin()
	go services.UpdateEveryday()
	initDiscordBot()
}









