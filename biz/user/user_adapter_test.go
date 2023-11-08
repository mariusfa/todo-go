package user

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestUserAdapterGet(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://jsonplaceholder.typicode.com/users",
		httpmock.NewJsonResponderOrPanic(http.StatusOK, httpmock.File("./users.json")))

	userAdapter := NewUserAdapter()
	users, err := userAdapter.GetAll()
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if len(users) != 2 {
		t.Errorf("Expected 2 users, got %v", len(users))
	}

	if users[0].ID != 1 {
		t.Errorf("Expected ID 1, got %v", users[0].ID)
	}

	if users[0].Name != "John Doe" {
		t.Errorf("Expected Name John Doe, got %v", users[0].Name)
	}
}
