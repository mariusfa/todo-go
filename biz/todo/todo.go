package todo

type Todo struct {
	Id   int
	Task string
}

func NewTodo(id int, task string) Todo {
	return Todo{Id: id, Task: task}
}
