package app

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"todo/biz/todo"
)

func TestTodoGetRoute(t *testing.T) {
	router, repositories := AppTestSetup()

	todoFakeRepo := repositories.TodoRepository.(*todo.TodoRepositoryFake)
	todoFakeRepo.Todos = []todo.Todo{todo.NewTodo(0, "Task 1")}

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/todo", nil)
	router.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("Response code is %v", response.Code)
	}
	expected := `[{"id":0,"task":"Task 1"}]`
	if response.Body.String() != expected {
		t.Errorf("Response body is %v", response.Body.String())
	}
}

func TestTodoPostRoute(t *testing.T) {
	router, repositories := AppTestSetup()

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
