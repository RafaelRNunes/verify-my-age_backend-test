package main

import (
	"github.com/RafaelRNunes/verify-my-age_backend-test/infra/api-gin-adapter/routes"
	"github.com/RafaelRNunes/verify-my-age_backend-test/infra/database"
	"github.com/spf13/viper"
)

func main() {
	// Viper
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	// Gorm
	database.Connection()

	// Gin web framework
	routes.HandleRequests().Run()
}
