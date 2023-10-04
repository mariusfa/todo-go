package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
