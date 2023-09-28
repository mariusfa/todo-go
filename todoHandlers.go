package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func todoGetHandler(c *gin.Context, db TodoDbAdapter) {
	todos, err := db.GetAll()
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, todos)
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