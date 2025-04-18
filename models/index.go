package models

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

type BaseModel struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (b BaseModel) GetID() uint { return b.ID }

type ItemsWithCount[T any] struct {
	Count int `json:"count"`
	Items []T `json:"items"`
}

func Init() {
	connectDB()
}

func connectDB() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	var err error
	dbConfig := viper.GetStringMapString("mysql")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConfig["user"], dbConfig["password"], dbConfig["host"], dbConfig["port"], dbConfig["db"])
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Migrate the schema
	if err = db.AutoMigrate(
		&Activity{},

		&BotServer{},
		&SocialUserConfig{},
		&CustomMintCount{},
		&CustomMintResult{},
		&POAPResult{},
		&H5Config{},
		&NFTConfig{},

		&Contract{},
		&Statistic{},
		&MetadataAttribute{},
		&PushInfo{},
		&BotServer{},
		&WalletUser{},
		&PhoneWhiteList{},
		&TokenReserve{},

		&BatchMintTask{},
	); err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return db
}
