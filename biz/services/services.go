package services

import (
	"todo/biz/adapters"
	"todo/biz/repositories"
	"todo/biz/todo"
	"todo/biz/user"
)

type Services struct {
	TodoService *todo.TodoService
	UserService *user.UserService
}

func NewServices(repositories *repositories.Repositories, adapters *adapters.Adapters) *Services {
	return &Services{
		TodoService: todo.NewTodoService(repositories.TodoRepository),
		UserService: user.NewUserService(adapters.UserAdapter),
	}
}
