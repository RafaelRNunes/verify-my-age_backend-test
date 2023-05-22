package database

import (
	"github.com/RafaelRNunes/verify-my-age_backend-test/infra/database/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var (
	DB  *gorm.DB
	err error
)

func Connection() {
	connectionString := "root:123mudar@tcp(localhost:3306)/api-users?charset=utf8mb4&parseTime=true"
	DB, err = gorm.Open(mysql.Open(connectionString))

	if err != nil {
		log.Panic("Can't connect to database", err.Error())
	}

	DB.AutoMigrate(&models.User{}, &models.Address{})
}
