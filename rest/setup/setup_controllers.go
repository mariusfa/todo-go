package setup

import (
	"todo/biz/setup"
	pc "todo/rest/ping"
	tc "todo/rest/todo"
)

type Controllers struct {
	PingController pc.PingController
	TodoController tc.TodoController
}

func SetupControllers(repositories setup.Repositories) Controllers {
	return Controllers{
		PingController: pc.PingController{},
		TodoController: tc.TodoController{
			TodoRepository: repositories.TodoRepository,
		},
	}
}
