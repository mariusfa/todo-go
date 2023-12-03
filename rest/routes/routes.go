package routes

import (
	"todo/rest/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(controllers *controllers.Controllers) *gin.Engine {
	router := gin.Default()
	controllers.Ping.RegisterRoutes(router)
	controllers.Todo.RegisterRoutes(router)
	controllers.User.RegisterRoutes(router)
	return router
}
