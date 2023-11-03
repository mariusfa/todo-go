package controllers

import (
	"todo/biz/setup"
	"todo/rest/ping"
	"todo/rest/todo"
)

type Controllers struct {
	PingController ping.PingController
	TodoController todo.TodoController
}

func NewControllers(repos setup.Repositories) Controllers {
	return Controllers{
		PingController: ping.NewPingController(),
		TodoController: todo.NewTodoController(repos.TodoRepository),
	}
}
