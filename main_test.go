package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPingRoute(t *testing.T) {
	mockDb := MockTodoDbAdapter{}
	router := setupRouter(mockDb)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("Response code is %v", response.Code)
	}
	expected := `{"message":"pong"}`
	if response.Body.String() != expected {
		t.Errorf("Response body is %v", response.Body.String())
	}
}

type MockTodoDbAdapter struct {
}

var todoList = []Todo{
	Todo{Task: "Task 1"},
	Todo{Task: "Task 2"},
	Todo{Task: "Task 3"},
}

func (m MockTodoDbAdapter) GetAll() ([]Todo, error) {
	return todoList, nil
}

func (m MockTodoDbAdapter) Insert(task string) error {
	todoList = append(todoList, Todo{Task: task})
	return nil
}

func TestTodoGetRoute(t *testing.T) {
	mockDb := MockTodoDbAdapter{}
	router := setupRouter(mockDb)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/todo", nil)
	router.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("Response code is %v", response.Code)
	}
	expected := `[{"task":"Task 1"},{"task":"Task 2"},{"task":"Task 3"}]`
	if response.Body.String() != expected {
		t.Errorf("Response body is %v", response.Body.String())
	}
}

func TestTodoPostRoute(t *testing.T) {
	mockDb := MockTodoDbAdapter{}
	router := setupRouter(mockDb)

	requestBody := strings.NewReader(`{"task": "Task 4"}`)
	response := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/todo", requestBody)
	request.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(response, request)

	if response.Code != http.StatusCreated {
		t.Errorf("Response code is %v", response.Code)
	}
	if len(todoList) != 4 {
		t.Errorf("Todo list length is %v", len(todoList))
	}

}
