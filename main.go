package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Todo struct {
	Task string `json:"task"`
}

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

type TodoRepositoryContract interface {
	Insert(string) error
	GetAll() ([]Todo, error)
}

type TodoRepository struct {
	db *sql.DB
}

func (todoRepository *TodoRepository) Insert(task string) error {
	_, err := todoRepository.db.Exec("INSERT INTO todos (task) VALUES ($1)", task)
	return err
}

func (todoRepository *TodoRepository) GetAll() ([]Todo, error) {
	rows, err := todoRepository.db.Query("SELECT task FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo

	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.Task)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func main() {
	db := setupDB()
	db.Exec("CREATE TABLE IF NOT EXISTS todos (id SERIAL PRIMARY KEY, task VARCHAR NOT NULL)")
	defer db.Close()

	todoRepository := &TodoRepository{db: db}
	router := setupRouter(todoRepository)
	router.Run()
}
