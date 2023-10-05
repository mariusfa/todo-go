package controller

import (
	"net/http"
	"todo/biz/todo"
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context, todoRepository todo.TodoRepositoryContract) {
	todos, err := todoRepository.GetAll()
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, todos)
}

func Post(c *gin.Context, todoRepository todo.TodoRepositoryContract) {
	var todo todo.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	todoRepository.Insert(todo.Task)
	c.Status(http.StatusCreated)
}
