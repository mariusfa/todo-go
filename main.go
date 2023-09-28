package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
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

func todoGetHandler(c *gin.Context) {
	c.JSON(http.StatusOK, todoList)
}

type Todo struct {
	Task string `json:"task"`
}

func todoPostHandler(c *gin.Context) {
	var todo Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	todoList = append(todoList, todo.Task)
	c.JSON(http.StatusOK, todoList)
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", pingHandler)
	router.GET("/todo", todoGetHandler)
	router.POST("/todo", todoPostHandler)
	return router
}

func main() {
	router := setupRouter()
	router.Run()
}
