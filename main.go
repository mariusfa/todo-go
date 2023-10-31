package main

import (
	restSetup "todo/rest/setup"
	bizSetup "todo/biz/setup"
)


func main() {
	db := SetupDB()
	defer db.Close()

	MigrateDB()

	repositories := bizSetup.SetupRepositories(db)
	controllers := restSetup.SetupControllers(repositories)
	router := restSetup.SetupRoutes(controllers)
	router.Run()
}
