package todo

type TodoRepositoryFake struct {
	Todos []Todo
}

func (todoRepository *TodoRepositoryFake) Insert(todo Todo) error {
	todoRepository.Todos = append(todoRepository.Todos, todo)
	return nil
}

func (todoRepository *TodoRepositoryFake) GetAll() ([]Todo, error) {
	return todoRepository.Todos, nil
}

func NewTodoRepositoryFake() *TodoRepositoryFake {
	return &TodoRepositoryFake{}
}
