package controllers

import (
	"todo/biz/services"
	"todo/rest/ping"
	"todo/rest/todo"
)

type Controllers struct {
	PingController *ping.PingController
	TodoController *todo.TodoController
}

func NewControllers(services *services.Services) *Controllers {
	return &Controllers{
		PingController: ping.NewPingController(),
		TodoController: todo.NewTodoController(services.TodoService),
	}
}
