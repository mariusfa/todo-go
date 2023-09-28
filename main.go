package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
)

func pingHandler(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "pong",
    })
}

func todoGetHandler(c *gin.Context, db TodoDbAdapter) {
	todos, err := db.GetAll()
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, todos)
}

type Todo struct {
	Task string `json:"task"`
}

func todoPostHandler(c *gin.Context, db TodoDbAdapter) {
	var todo Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	db.Insert(todo.Task)
	c.Status(http.StatusCreated)
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

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "password"
	dbname = "postgres"
)

func setupDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	return db
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
