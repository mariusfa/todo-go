package controllers

import (
	"todo/biz/services"
	"todo/rest/ping"
	"todo/rest/todo"
	"todo/rest/user"
)

type Controllers struct {
	PingController *ping.PingController
	TodoController *todo.TodoController
	UserController *user.UserController
}

func NewControllers(services *services.Services) *Controllers {
	return &Controllers{
		PingController: ping.NewPingController(),
		TodoController: todo.NewTodoController(services.TodoService),
		UserController: user.NewUserController(services.UserService),
	}
}
