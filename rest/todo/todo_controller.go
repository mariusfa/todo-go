package todo

import (
	"net/http"
	"todo/biz/todo"
	"github.com/gin-gonic/gin"
)

type TodoController struct {
	TodoRepository todo.TodoRepositoryContract
}

func (tc *TodoController) Get(c *gin.Context) {
	todos, err := tc.TodoRepository.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, todos)
}

func (tc *TodoController) Post(c *gin.Context) {
	var todo todo.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	tc.TodoRepository.Insert(todo.Task)
	c.Status(http.StatusCreated)
}
