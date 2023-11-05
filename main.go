package main

import (
	"todo/biz/repositories"
	"todo/biz/services"
	"todo/config"
	"todo/database"
	"todo/rest/controllers"
	"todo/rest/routes"
)

func main() {
	database.Migrate(config.GetMigrationDbConfig(), "./migrations")

	db, err := database.SetupDb(config.GetAppDbConfig())
	if err != nil {
		panic(err)
	}

	repositories := repositories.NewRepositories(db)
	services := services.NewServices(repositories)
	controllers := controllers.NewControllers(services)
	router := routes.SetupRoutes(controllers)
	router.Run()
}
