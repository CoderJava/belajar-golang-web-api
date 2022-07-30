package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name": "Yudi Setiawan",
			"bio":  "A Software Engineer",
		})
	})

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"content":  "Hello World",
			"subtitle": "Belajar Golang",
		})
	})

	router.Run()
}
