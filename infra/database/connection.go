package database

import (
	"fmt"
	"github.com/RafaelRNunes/verify-my-age_backend-test/infra/database/models"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var (
	DB  *gorm.DB
	err error
)

func Connection() {
	DB_USER := viper.GetString("DB_USER")
	DB_PASS := viper.GetString("DB_PASS")
	DB_HOST := viper.GetString("DB_HOST")
	DB_NAME := viper.GetString("DB_NAME")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true",
		DB_USER, DB_PASS, DB_HOST, DB_NAME)

	DB, err = gorm.Open(mysql.Open(connectionString))

	if err != nil {
		log.Panic("Can't connect to database", err.Error())
	}

	DB.AutoMigrate(&models.User{}, &models.Address{})
}
