package main

import (
	"github.com/RafaelRNunes/verify-my-age_backend-test/infra/api-gin-adapter/routes"
	"github.com/RafaelRNunes/verify-my-age_backend-test/infra/database"
)

func main() {
	database.Connection()
	routes.HandleRequests().Run()
}
