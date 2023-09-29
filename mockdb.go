package main

type TodoRepositoryFake struct {
	todoList []Todo
}

func NewTodoRepositoryFake() *TodoRepositoryFake {
	return &TodoRepositoryFake{
		todoList: []Todo{
			{Task: "Task 1"},
			{Task: "Task 2"},
			{Task: "Task 3"},
		},
	}
}

func (todoRepository *TodoRepositoryFake) GetAll() ([]Todo, error) {
	return todoRepository.todoList, nil
}

func (todoRepository *TodoRepositoryFake) Insert(task string) error {
	todoRepository.todoList = append(todoRepository.todoList, Todo{Task: task})
	return nil
}
