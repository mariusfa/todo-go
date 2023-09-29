package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPingRoute(t *testing.T) {
	todoRepository := NewTodoRepositoryFake()
	router := setupRouter(todoRepository)

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

func TestTodoGetRoute(t *testing.T) {
	todoRepository := NewTodoRepositoryFake()
	router := setupRouter(todoRepository)

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
	todoRepository := NewTodoRepositoryFake()
	router := setupRouter(todoRepository)

	requestBody := strings.NewReader(`{"task": "Task 4"}`)
	response := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/todo", requestBody)
	request.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(response, request)

	if response.Code != http.StatusCreated {
		t.Errorf("Response code is %v", response.Code)
	}
	if len(todoRepository.todoList) != 4 {
		t.Errorf("Todo list length is %v", len(todoRepository.todoList))
	}
}
