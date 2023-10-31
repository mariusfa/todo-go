package todo

type TodoRepositoryFake struct {
	Todos []Todo
}

func (todoRepository *TodoRepositoryFake) Insert(task string) error {
	todo := Todo{Task: task}
	todoRepository.Todos = append(todoRepository.Todos, todo)
	return nil
}

func (todoRepository *TodoRepositoryFake) GetAll() ([]Todo, error) {
	return todoRepository.Todos, nil
}

func NewTodoRepositoryFake() *TodoRepositoryFake {
	return &TodoRepositoryFake{}
}
