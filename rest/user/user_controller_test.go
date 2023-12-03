package user

import (
	"net/http"
	"net/http/httptest"
	"testing"
	bizUser "todo/biz/user"

	"github.com/gin-gonic/gin"
)

func setup() *gin.Engine {
	router := gin.New()
	usFake := &bizUser.UserServiceFake{}
	controller := NewController(usFake)
	controller.RegisterRoutes(router)
	return router
}

func TestGet(t *testing.T) {
	router := setup()

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
