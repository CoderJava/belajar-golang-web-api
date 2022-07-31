package main

import (
	"github.com/gin-gonic/gin"

	handlerV1 "pustaka-api/handler/v1"
	handlerV2 "pustaka-api/handler/v2"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	v1.GET("/", handlerV1.RootHandler)
	v1.GET("/hello", handlerV1.HelloHandler)
	v1.GET("/books/:id/:title", handlerV1.BooksHandler)
	v1.POST("/books", handlerV1.PostBooksHandler)
	v1.GET("/query", handlerV1.QueryHandler)

	v2 := router.Group("/v2")
	v2.GET("/", handlerV2.RootHandler)

	router.Run()
}
