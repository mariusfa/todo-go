package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct{}

func (ct *Controller) RegisterRoutes(router *gin.Engine) {
	router.GET("/ping", ct.Get)
}

func (ct *Controller) Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func NewController() *Controller {
	return &Controller{}
}
