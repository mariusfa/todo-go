package todo

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo/biz/todo"
)

type TodoController struct {
	todoService *todo.TodoService
}

func (tc *TodoController) Get(c *gin.Context) {
	todos, err := tc.todoService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, FromDomainList(todos))
}

func (tc *TodoController) Post(c *gin.Context) {
	var todoRequestDTO TodoRequestDTO
	if err := c.ShouldBindJSON(&todoRequestDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	tc.todoService.Create(ToDomain(todoRequestDTO))

	c.Status(http.StatusCreated)
}

func NewTodoController(todoService *todo.TodoService) *TodoController {
	return &TodoController{todoService: todoService}
}
