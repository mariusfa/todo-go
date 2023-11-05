package todo

import "todo/biz/todo"

type TodoDTO struct {
	Id   int    `json:"id"`
	Task string `json:"task"`
}

func NewTodoDTO(id int, task string) TodoDTO {
	return TodoDTO{Id: id, Task: task}
}

func FromDomain(todo todo.Todo) TodoDTO {
	return NewTodoDTO(todo.Id, todo.Task)
}

func FromDomainList(todos []todo.Todo) []TodoDTO {
	var todoDTOs []TodoDTO
	for _, todo := range todos {
		todoDTOs = append(todoDTOs, FromDomain(todo))
	}
	return todoDTOs
}
