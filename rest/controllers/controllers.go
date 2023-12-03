package controllers

import (
	"todo/biz/services"
	"todo/rest/ping"
	"todo/rest/todo"
	"todo/rest/user"
)

type Controllers struct {
	Ping *ping.Controller
	Todo *todo.TodoController
	User *user.Controller
}

func NewControllers(services *services.Services) *Controllers {
	return &Controllers{
		Ping: ping.NewController(),
		Todo: todo.NewTodoController(services.TodoService),
		User: user.NewController(services.UserService),
	}
}
