package setup

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	bizSetup "todo/biz/setup"
	"todo/biz/todo"
)

func TestTodoGetRoute(t *testing.T) {
	repositories := bizSetup.SetupRepositoriesFake()
	controllers := SetupControllers(repositories)
	router := SetupRoutes(controllers)

	todoFakeRepo := repositories.TodoRepository.(*todo.TodoRepositoryFake)
	todoFakeRepo.Todos = []todo.Todo{{Task: "Task 1"}}

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/todo", nil)
	router.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("Response code is %v", response.Code)
	}
	expected := `[{"task":"Task 1"}]`
	if response.Body.String() != expected {
		t.Errorf("Response body is %v", response.Body.String())
	}
}

func TestTodoPostRoute(t *testing.T) {
	repositories := bizSetup.SetupRepositoriesFake()
	controllers := SetupControllers(repositories)
	router := SetupRoutes(controllers)

	requestBody := strings.NewReader(`{"task": "Task 4"}`)
	response := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/todo", requestBody)
	request.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(response, request)

	if response.Code != http.StatusCreated {
		t.Errorf("Response code is %v", response.Code)
	}
	fakeRepo := repositories.TodoRepository.(*todo.TodoRepositoryFake)
	if len(fakeRepo.Todos) != 1 {
		t.Errorf("Todo list length is %v", len(fakeRepo.Todos))
	}
}
