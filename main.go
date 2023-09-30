package main

import (
	"github.com/gin-gonic/gin"
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
	db.Exec("CREATE TABLE IF NOT EXISTS todos (id SERIAL PRIMARY KEY, task VARCHAR NOT NULL)")
	defer db.Close()

	todoRepository := &TodoRepository{db: db}
	router := setupRouter(todoRepository)
	router.Run()
}
