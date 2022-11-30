package models

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
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

func ConnectDB() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	var err error
	dbConfig := viper.GetStringMapString("mysql")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConfig["user"], dbConfig["password"], dbConfig["host"], dbConfig["port"], dbConfig["db"])
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Migrate the schema
	db.AutoMigrate(&DiscordCustomProjectConfig{})
	db.AutoMigrate(&DoDoCustomProjectConfig{})
	db.AutoMigrate(&DiscordCustomActivityConfig{})
	db.AutoMigrate(&DoDoCustomActivityConfig{})
	db.AutoMigrate(&BindCFXWithDiscord{})
	db.AutoMigrate(&BindCFXWithDoDo{})
	db.AutoMigrate(&CustomMintCount{})
	db.AutoMigrate(&CustomMintResult{})
	db.AutoMigrate(&POAPResult{})
	db.AutoMigrate(&POAPActivityConfig{})
}

func GetDB() *gorm.DB {
	return db
}
