package app

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPing(t *testing.T) {
	router, _ := AppTestSetup()

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
