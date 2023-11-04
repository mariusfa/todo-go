package repositories

import (
	"database/sql"
	"todo/biz/todo"
)

type Repositories struct {
	TodoRepository todo.TodoRepositoryContract
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		TodoRepository: todo.NewTodoRepository(db),
	}
}

func NewRepositoriesFake() *Repositories {
	return &Repositories{
		TodoRepository: todo.NewTodoRepositoryFake(),
	}
}
