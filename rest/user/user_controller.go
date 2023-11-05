package user

import (
	"net/http"
	"todo/biz/user"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *user.UserService
}

func NewUserController(userService *user.UserService) *UserController {
	return &UserController{userService: userService}
}

func (u *UserController) Get(c *gin.Context) {
	users, err := u.userService.GetAll()
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
