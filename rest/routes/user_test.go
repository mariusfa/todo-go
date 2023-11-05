package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"todo/biz/adapters"
	"todo/biz/repositories"
	"todo/biz/services"
	"todo/rest/controllers"
)

func TestGetUsers(t *testing.T) {
	// Setup
	repositories := repositories.NewRepositoriesFake()
	adapters := adapters.NewAdapterFakes()
	services := services.NewServices(repositories, adapters)
	controllers := controllers.NewControllers(services)
	router := SetupRoutes(controllers)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/user", nil)
	router.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("Response code is %v", response.Code)
	}
	expected := `[{"id":1,"name":"User 1"},{"id":2,"name":"User 2"},{"id":3,"name":"User 3"}]`
	if response.Body.String() != expected {
		t.Errorf("Response body is %v", response.Body.String())
	}
}
