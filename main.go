package main

import (
	"fmt"
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
	// CRUD

	// Create
	book := book.Book{}
	book.Title = "Atomic Habits"
	book.Price = 120000
	book.Discount = 15
	book.Rating = 4
	book.Description = "Buku self developent tentang membangun kebiasaan baik dan menghilangkan kebiasaan buruk"
	err = db.Create(&book).Error
	if err != nil {
		fmt.Println("==========================")
		fmt.Println("Error creating book record")
		fmt.Println("==========================")
	}

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
