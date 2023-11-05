package todo

type TodoService struct {
	todoRepository TodoRepositoryContract
}

func NewTodoService(todoRepository TodoRepositoryContract) *TodoService {
	return &TodoService{todoRepository: todoRepository}
}

func (todoService *TodoService) Create(todo Todo) error {
	return todoService.todoRepository.Insert(todo)
}

func (todoService *TodoService) GetAll() ([]Todo, error) {
	return todoService.todoRepository.GetAll()
}
