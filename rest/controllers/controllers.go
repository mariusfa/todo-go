package controllers

import (
	"todo/biz/repositories"
	"todo/rest/ping"
	"todo/rest/todo"
)

type Controllers struct {
	PingController *ping.PingController
	TodoController *todo.TodoController
}

func NewControllers(repos *repositories.Repositories) *Controllers {
	return &Controllers{
		PingController: ping.NewPingController(),
		TodoController: todo.NewTodoController(repos.TodoRepository),
	}
}
