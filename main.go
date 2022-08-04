package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"pustaka-api/book"
	handlerV1 "pustaka-api/handler/v1"
	handlerV2 "pustaka-api/handler/v2"
)

func main() {
	dsn := "root:admin@tcp(127.0.0.1:3306)/pustaka-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Db connection error")
	}
	db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	// fileRepository := book.NewFileRepository()
	// bookService := book.NewService(fileRepository)
	bookHandler := handlerV1.NewBookHandler(bookService)

	router := gin.Default()

	v1 := router.Group("/v1")
	v1.GET("/", handlerV1.RootHandler)
	v1.GET("/hello", handlerV1.HelloHandler)
	v1.GET("/query", handlerV1.QueryHandler)

	v1.GET("/books/:id/:title", handlerV1.BooksHandler)
	v1.GET("/books", bookHandler.GetAllBooksHandler)
	v1.GET("/books/:id", bookHandler.GetBookHandler)
	v1.POST("/books", bookHandler.PostBooksHandler)
	v1.DELETE("/books/:id", bookHandler.DeleteBookHandler)

	v2 := router.Group("/v2")
	v2.GET("/", handlerV2.RootHandler)

	router.Run()

	// main
	// handler
	// service
	// repository
	// db
	// mariadb
}
