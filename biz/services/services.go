package services

import (
	"todo/biz/repositories"
	"todo/biz/todo"
)

type Services struct {
	TodoService *todo.TodoService
}

func NewServices(repositories *repositories.Repositories) *Services {
	return &Services{
		TodoService: todo.NewTodoService(repositories.TodoRepository),
	}
}