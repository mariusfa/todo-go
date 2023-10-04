package main

import (
	"todo/biz/todo"
	pingController "todo/rest/ping/controller"
	todoController "todo/rest/todo/controller"

	"github.com/gin-gonic/gin"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func setupRouter(todoRepository todo.TodoRepositoryContract) *gin.Engine {
	router := gin.Default()
	router.GET("/ping", pingController.GetPong)
	router.GET("/todo", func(c *gin.Context) {
		todoController.Get(c, todoRepository)
	})
	router.POST("/todo", func(c *gin.Context) {
		todoController.Post(c, todoRepository)
	})
	return router
}

func main() {
	db := SetupDB()
	defer db.Close()

	MigrateDB()

	todoRepository := todo.NewTodoRepository(db)
	router := setupRouter(todoRepository)
	router.Run()
}
