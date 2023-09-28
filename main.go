package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Todo struct {
	Task string `json:"task"`
}

func setupRouter(db TodoDbAdapter) *gin.Engine {
	router := gin.Default()
	router.GET("/ping", pingHandler)
	router.GET("/todo", func(c *gin.Context) {
		todoGetHandler(c, db)
	})
	router.POST("/todo", func(c *gin.Context) {
		todoPostHandler(c, db)
	})
	return router
}

type TodoDbAdapter interface {
	Insert(string) error
	GetAll() ([]Todo, error)
}

type TodoDb struct {
	db *sql.DB
}

func (db *TodoDb) Insert(task string) error {
	_, err := db.db.Exec("INSERT INTO todos (task) VALUES ($1)", task)
	return err
}

func (db *TodoDb) GetAll() ([]Todo, error) {
	rows, err := db.db.Query("SELECT task FROM todos")
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

	todoAdapter := &TodoDb{db: db}
	router := setupRouter(todoAdapter)
	router.Run()
}
