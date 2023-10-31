package main

import (
	bizSetup "todo/biz/setup"
	"todo/config"
	"todo/database"
	restSetup "todo/rest/setup"
)

func main() {
	database.Migrate(config.GetMigrationDbConfig(), "./migrations")

	db, err := database.SetupDb(config.GetAppDbConfig())
	if err != nil {
		panic(err)
	}

	repositories := bizSetup.SetupRepositories(db)
	controllers := restSetup.SetupControllers(repositories)
	router := restSetup.SetupRoutes(controllers)
	router.Run()
}
