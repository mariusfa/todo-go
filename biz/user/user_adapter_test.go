package user

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestUserAdapterGet(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://jsonplaceholder.typicode.com/users/1",
		httpmock.NewJsonResponderOrPanic(http.StatusOK, httpmock.File("./users.json")))


	// fetch json using User adapter

	// assert

}

