package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func todoGetHandler(c *gin.Context, todoRepository TodoRepositoryContract) {
	todos, err := todoRepository.GetAll()
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, todos)
}

func todoPostHandler(c *gin.Context, todoRepository TodoRepositoryContract) {
	var todo Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	todoRepository.Insert(todo.Task)
	c.Status(http.StatusCreated)
}
