package user

import (
	"net/http"
	"todo/biz/user"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	userService *user.UserService
}

func NewController(userService *user.UserService) *Controller {
	return &Controller{userService: userService}
}

func (ct *Controller) Get(c *gin.Context) {
	users, err := ct.userService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, FromDomainList(users))
}

//https://jsonplaceholder.typicode.com/users
// Use this to create a http client adapter and use http mock for adapter test
