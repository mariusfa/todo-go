package routes

import (
	"todo/rest/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(controllers *controllers.Controllers) *gin.Engine {
	router := gin.Default()
	controllers.Ping.RegisterRoutes(router)
	router.GET("/todo", controllers.TodoController.Get)
	router.POST("/todo", controllers.TodoController.Post)
	controllers.User.RegisterRoutes(router)
	return router
}
