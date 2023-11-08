package app

import (
	"database/sql"
	"todo/biz/adapters"
	"todo/biz/repositories"
	"todo/biz/services"
	"todo/rest/controllers"
	"todo/rest/routes"

	"github.com/gin-gonic/gin"
)

func AppSetup(db *sql.DB) *gin.Engine {
	repositories := repositories.NewRepositories(db)
	adapters := adapters.NewAdapters()
	services := services.NewServices(repositories, adapters)
	controllers := controllers.NewControllers(services)
	return routes.SetupRoutes(controllers)
}

func AppTestSetup() (*gin.Engine, *repositories.Repositories) {
	repositories := repositories.NewRepositoriesFake()
	adapters := adapters.NewAdapterFakes()
	services := services.NewServices(repositories, adapters)
	controllers := controllers.NewControllers(services)
	return routes.SetupRoutes(controllers), repositories
}
