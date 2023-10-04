package main

import (
	"todo/biz/todo"
)

type TodoRepositoryFake struct {
	todoList []todo.Todo
}

func NewTodoRepositoryFake() *TodoRepositoryFake {
	return &TodoRepositoryFake{
		todoList: []todo.Todo{
			{Task: "Task 1"},
			{Task: "Task 2"},
			{Task: "Task 3"},
		},
	}
}

func (todoRepository *TodoRepositoryFake) GetAll() ([]todo.Todo, error) {
	return todoRepository.todoList, nil
}

func (todoRepository *TodoRepositoryFake) Insert(task string) error {
	todoRepository.todoList = append(todoRepository.todoList, todo.Todo{Task: task})
	return nil
}
