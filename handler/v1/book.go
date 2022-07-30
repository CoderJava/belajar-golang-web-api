package v1

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"pustaka-api/book"
)

func BooksHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")
	c.JSON(http.StatusOK, gin.H{
		"id":    id,
		"title": title,
	})
}

func PostBooksHandler(c *gin.Context) {
	var bookInput book.BookInput

	err := c.ShouldBindJSON(&bookInput)
	if err != nil {
		errorMessages := []map[string]string{}
		for _, e := range err.(validator.ValidationErrors) {
			field := strings.ToLower(e.Field())
			error := e.ActualTag()
			errorMessage := map[string]string{
				"field": field,
				"error": error,
			}
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title":    bookInput.Title,
		"price":    bookInput.Price,
		"subtitle": bookInput.SubTitle,
	})
}
