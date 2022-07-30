package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"content": "Hello World v1",
		"subtitle": "Belajar Golang",
	})
}