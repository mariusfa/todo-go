package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"todo/biz/adapters"
	"todo/biz/repositories"
	"todo/biz/services"
	"todo/rest/controllers"
	"todo/rest/routes"
)

func TestGetPing(t *testing.T) {
	repositories := repositories.NewRepositoriesFake()
	adapters := adapters.NewAdapterFakes()
	services := services.NewServices(repositories, adapters)
	controllers := controllers.NewControllers(services)
	router := routes.SetupRoutes(controllers)

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
