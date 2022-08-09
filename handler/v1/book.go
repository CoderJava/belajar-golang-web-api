package v1

import (
	"net/http"
	"strconv"
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

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (h *bookHandler) GetAllBooksHandler(c *gin.Context) {
	books, err := h.bookService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	booksResponse := []book.BookResponse{}
	for _, entityBook := range books {
		bookResponse := convertEntityBookToBookResponse(entityBook)
		booksResponse = append(booksResponse, bookResponse)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": booksResponse,
	})
}

func (h *bookHandler) GetBookHandler(c *gin.Context) {
	paramId := c.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	bookData, err := h.bookService.FindById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	bookResponse := convertEntityBookToBookResponse(bookData)
	c.JSON(http.StatusOK, bookResponse)
}

func (h *bookHandler) DeleteBookHandler(c *gin.Context) {
	paramId := c.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	_, err = h.bookService.DeleteById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"title":   "Success",
		"message": "Book successfully deleted",
	})
}

func (h *bookHandler) CreateBookHandler(c *gin.Context) {
	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest)
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

	book, err := h.bookService.Create(bookRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, convertEntityBookToBookResponse(book))
}

func (h *bookHandler) UpdateBookHandler(c *gin.Context) {
	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest)
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

	paramId := c.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	book, err := h.bookService.Update(id, bookRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, convertEntityBookToBookResponse(book))
}

func convertEntityBookToBookResponse(bookEntity book.Book) book.BookResponse {
	return book.BookResponse{
		ID:          bookEntity.ID,
		Title:       bookEntity.Title,
		Price:       bookEntity.Price,
		Description: bookEntity.Description,
		Rating:      bookEntity.Rating,
		Discount:    bookEntity.Discount,
	}
}
