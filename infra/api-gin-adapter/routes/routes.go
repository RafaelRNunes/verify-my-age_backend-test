package routes

import (
	"github.com/RafaelRNunes/verify-my-age_backend-test/infra/api-gin-adapter/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func HandleRequests() *gin.Engine {
	routes := gin.Default()
	routes.Use(cors.Default())

	// Users routes
	routes.GET("/users", controllers.FindUsers)
	routes.GET("/users/:id", controllers.FindUserById)
	routes.POST("/users", controllers.CreateUser)
	routes.PATCH("/users/:id", controllers.UpdateUser)
	routes.DELETE("/users/:id", controllers.DeleteUser)

	return routes
}
