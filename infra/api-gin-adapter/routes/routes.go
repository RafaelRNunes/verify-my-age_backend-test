package routes

import (
	docs "github.com/RafaelRNunes/verify-my-age_backend-test/docs"
	"github.com/RafaelRNunes/verify-my-age_backend-test/infra/api-gin-adapter/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func HandleRequests() *gin.Engine {
	routes := gin.Default()
	routes.Use(cors.Default())
	docs.SwaggerInfo.BasePath = "/"

	// Swagger
	routes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Users routes
	routes.GET("/users", controllers.FindUsers)
	routes.GET("/users/:id", controllers.FindUserById)
	routes.POST("/users", controllers.CreateUser)
	routes.PATCH("/users/:id", controllers.UpdateUser)
	routes.DELETE("/users/:id", controllers.DeleteUser)

	return routes
}
