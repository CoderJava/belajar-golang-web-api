package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func QueryHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")
	c.JSON(http.StatusOK, gin.H{
		"title": title,
		"price": price,
	})
}
