package setup

import (
	"database/sql"
	"todo/biz/todo"
)

type Repositories struct {
	TodoRepository todo.TodoRepositoryContract
}

func SetupRepositories(db *sql.DB) Repositories {
	return Repositories{
		TodoRepository: todo.NewTodoRepository(db),
	}
}

func SetupRepositoriesFake() Repositories {
	return Repositories{
		TodoRepository: todo.NewTodoRepositoryFake(),
	}
}