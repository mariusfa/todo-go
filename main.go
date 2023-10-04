package main

import (
	"github.com/gin-gonic/gin"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func setupRouter(todoRepository TodoRepositoryContract) *gin.Engine {
	router := gin.Default()
	router.GET("/ping", pingHandler)
	router.GET("/todo", func(c *gin.Context) {
		todoGetHandler(c, todoRepository)
	})
	router.POST("/todo", func(c *gin.Context) {
		todoPostHandler(c, todoRepository)
	})
	return router
}

func main() {
	db := setupDB()
	// db.Exec("CREATE TABLE IF NOT EXISTS todos (id SERIAL PRIMARY KEY, task VARCHAR NOT NULL)")
	defer db.Close()

	migrateDB()

	todoRepository := &TodoRepository{db: db}
	router := setupRouter(todoRepository)
	router.Run()
}
