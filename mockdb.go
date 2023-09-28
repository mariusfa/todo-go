package main

type MockTodoDbAdapter struct {
	todoList []Todo
}

func NewMockTodoDbAdapter() *MockTodoDbAdapter {
	return &MockTodoDbAdapter{
		todoList: []Todo{
			Todo{Task: "Task 1"},
			Todo{Task: "Task 2"},
			Todo{Task: "Task 3"},
		},
	}
}

func (m *MockTodoDbAdapter) GetAll() ([]Todo, error) {
	return m.todoList, nil
}

func (m *MockTodoDbAdapter) Insert(task string) error {
	m.todoList = append(m.todoList, Todo{Task: task})
	return nil
}
