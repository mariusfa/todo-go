package todo

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"
	"todo/biz/testcontainers"
	"todo/config"
	"todo/database"
)

var todoDbConfig config.DbConfig

func TestMain(m *testing.M) {
	container, err := testcontainers.CreatePostgresContainer()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := container.Terminate(context.Background())
		if err != nil {
			log.Fatal(err)
		}
	}()

	migrationConfig, err := testcontainers.GetMigrationTestConfig(container)
	if err != nil {
		log.Fatal(err)
	}
	todoDbConfig = testcontainers.GetAppTestConfig(migrationConfig)

	err = database.Migrate(migrationConfig, "../../migrations")
	if err != nil {
		log.Fatal(err)
	}

	code := m.Run()
	os.Exit(code)
}

func TestInsert(t *testing.T) {
	// given
	db, err := database.SetupDb(todoDbConfig)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	todoRepository := NewTodoRepository(db)

	// when
	todoToInsert := NewTodo(0, "Test")
	err = todoRepository.Insert(todoToInsert)
	if err != nil {
		t.Fatal(err)
	}

	// then
	var todo Todo
	if err := db.QueryRow("SELECT task FROM todoschema.todos").Scan(&todo.Task); err != nil {
		t.Errorf("Failed to get task: %v", err)
	}
	if todo.Task != "Test" {
		t.Errorf("Expected task to be 'Test', got '%s'", todo.Task)
	}

	// cleanup
	if err := clearTodoTable(db); err != nil {
		t.Fatal(err)
	}
}

func TestGetAll(t *testing.T) {
	// given
	db, err := database.SetupDb(todoDbConfig)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	todoRepository := NewTodoRepository(db)

	todoToInsert := NewTodo(0, "Test")
	err = todoRepository.Insert(todoToInsert)
	if err != nil {
		t.Fatal(err)
	}

	// when
	todos, err := todoRepository.GetAll()
	if err != nil {
		t.Fatal(err)
	}

	// then
	if len(todos) != 1 {
		t.Errorf("Expected 1 todo, got %d", len(todos))
	}
	if todos[0].Task != "Test" {
		t.Errorf("Expected task to be 'Test', got '%s'", todos[0].Task)
	}

	// cleanup
	if err := clearTodoTable(db); err != nil {
		t.Fatal(err)
	}
}

func clearTodoTable(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM todoschema.todos")
	return err
}
