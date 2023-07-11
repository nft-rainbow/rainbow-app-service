package main

import (
	"fmt"
	"log"
	"os"

	"encoding/csv"

	"flag"

	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/spf13/viper"
)

func init() {
	initConfig()
	models.Init()
}

func initConfig() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		log.Fatalln(fmt.Errorf("fatal error config file: %w", err))
	}
}

type PhoneInfo struct {
	Phone string
}

// NOTE: phone shoule be the first column
func createPhoneList(data [][]string) []PhoneInfo {
	var shoppingList []PhoneInfo
	for i, line := range data {
		if i > 0 { // omit header line
			var rec PhoneInfo
			for j, field := range line {
				if j == 0 {
					rec.Phone = field
				}
			}
			shoppingList = append(shoppingList, rec)
		}
	}
	return shoppingList
}

// TODO update this two value
const ACTIVITY_CODE = "changeAnDao"
const CSV_FILE = "./scripts/data.csv"

func main() {
	activityId := flag.String("activity_code", ACTIVITY_CODE, "the activity id")
	csvFile := flag.String("csv", CSV_FILE, "the csv file")
	flag.Parse()

	fmt.Println("Start importing.")

	// open file
	f, err := os.Open(*csvFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	// convert records to array of structs
	list := createPhoneList(data)

	for index, phone := range list {
		fmt.Println("Dealing", index, phone.Phone)
		inList := models.IsPhoneInWhiteList(*activityId, phone.Phone)
		if !inList {
			models.GetDB().Create(&models.PhoneWhiteList{
				ActivityCode: *activityId,
				Phone:        phone.Phone,
			})
		}
	}

	fmt.Println("Finished ")
}
