package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"
)

func TestPingRoute(t *testing.T) {
	router := setupRouter()

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
	router := setupRouter()

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/todo", nil)
	router.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("Response code is %v", response.Code)
	}
	expected := `["Task 1","Task 2","Task 3"]`
	if response.Body.String() != expected {
		t.Errorf("Response body is %v", response.Body.String())
	}
}

func TestTodoPostRoute(t *testing.T) {
	router := setupRouter()

	requestBody := strings.NewReader(`{"task": "Task 4"}`)
	response := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/todo", requestBody)
	request.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("Response code is %v", response.Code)
	}
	expected := `["Task 1","Task 2","Task 3","Task 4"]`
	if response.Body.String() != expected {
		t.Errorf("Response body is %v", response.Body.String())
	}

}