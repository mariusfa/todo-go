package todo

import "database/sql"

type TodoRepositoryContract interface {
	Insert(string) error
	GetAll() ([]Todo, error)
}

type TodoRepository struct {
	db *sql.DB
}

func (todoRepository *TodoRepository) Insert(task string) error {
	_, err := todoRepository.db.Exec("INSERT INTO todos (task) VALUES ($1)", task)
	return err
}

func (todoRepository *TodoRepository) GetAll() ([]Todo, error) {
	rows, err := todoRepository.db.Query("SELECT task FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo

	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.Task)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func NewTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{db: db}
}