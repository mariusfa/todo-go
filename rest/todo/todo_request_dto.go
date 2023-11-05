package todo

import "todo/biz/todo"

type TodoRequestDTO struct {
	Task string `json:"task"`
}

func ToDomain(todoRequestDTO TodoRequestDTO) todo.Todo {
	return todo.NewTodo(
		0,
		todoRequestDTO.Task,
	)
}
