package routes

import (
	"todo/rest/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(controllers controllers.Controllers) *gin.Engine {
	router := gin.Default()
	router.GET("/ping", controllers.PingController.Get)
	router.GET("/todo", controllers.TodoController.Get)
	router.POST("/todo", controllers.TodoController.Post)
	return router
}
