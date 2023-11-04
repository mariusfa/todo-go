package ping

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type PingController struct{}

func (tc *PingController) Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func NewPingController() *PingController {
	return &PingController{}
}
