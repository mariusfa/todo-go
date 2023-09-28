package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
)

var todoList = []string{
	"Task 1",
	"Task 2",
	"Task 3",
}

func pingHandler(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "pong",
    })
}

func todoGetHandler(c *gin.Context, db DbInstance) {
	rows, err := db.Query("SELECT task FROM todos")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var todos []Todo

	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.Task)
		if err != nil {
			panic(err)
		}
		todos = append(todos, todo)
	}
	c.JSON(http.StatusOK, todos)
}

type Todo struct {
	Task string `json:"task"`
}

func todoPostHandler(c *gin.Context, db DbInstance) {
	var todo Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	db.Exec("INSERT INTO todos (task) VALUES ($1)", todo.Task)
	c.Status(http.StatusCreated)
}

func setupRouter(db DbInstance) *gin.Engine {
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

type DbInstance interface {
	Exec(string, ...interface{}) (sql.Result, error)
	Query(string, ...interface{}) (*sql.Rows, error)
	Close() error
}

type TodoDbAdapter interface {
	Insert(string) error
	GetAll() ([]Todo, error)
}




func main() {
	var db DbInstance
	db = setupDB()
	db.Exec("CREATE TABLE IF NOT EXISTS todos (id SERIAL PRIMARY KEY, task VARCHAR NOT NULL)")

	defer db.Close()
	router := setupRouter(db)
	router.Run()
}
