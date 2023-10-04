package main

import (
	"github.com/gin-gonic/gin"
	pingController "todo/rest/ping/controller"
	todoController "todo/rest/todo/controller"
	"todo/biz/todo"

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
	db := setupDB()
	// db.Exec("CREATE TABLE IF NOT EXISTS todos (id SERIAL PRIMARY KEY, task VARCHAR NOT NULL)")
	defer db.Close()

	migrateDB()


	todoRepository := todo.NewTodoRepository(db)
	router := setupRouter(todoRepository)
	router.Run()
}
