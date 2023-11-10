package main

import (
	"todo/app"
	"todo/config"
	"todo/database"
)

func main() {
	if err := database.Migrate(config.GetMigrationDbConfig(), "./migrations"); err != nil {
		panic(err)
	}

	db, err := database.SetupDb(config.GetAppDbConfig())
	if err != nil {
		panic(err)
	}

	router := app.AppSetup(db)
	router.Run()
}
