package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"
	bizSetup "todo/biz/setup"
	"todo/rest/controllers"
)

func TestGetPing(t *testing.T) {
	repositories := bizSetup.SetupRepositoriesFake()
	controllers := controllers.NewControllers(repositories)
	router := SetupRoutes(controllers)

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
