package main

import (
	"todo/app"
	"todo/config"
	"todo/database"
)

func main() {
	database.Migrate(config.GetMigrationDbConfig(), "./migrations")

	db, err := database.SetupDb(config.GetAppDbConfig())
	if err != nil {
		panic(err)
	}

	router := app.AppSetup(db)
	router.Run()
}
