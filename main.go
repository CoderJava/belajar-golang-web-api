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

	// bookRepository := book.NewRepository(db)

	// Find all
	/* books, err := bookRepository.FindAll()
	if err != nil {
		fmt.Println("Error finding all books")
		return
	}
	for _, book := range books {
		fmt.Println("Title:", book.Title)
	} */

	// Find by ID
	/* book, err := bookRepository.FindById(2)
	if err != nil {
		fmt.Println("Error find by ID")
		return
	}
	fmt.Println("Title:", book.Title) */

	// Create
	/* book := book.Book{
		Title:       "Belajar Golang",
		Description: "Buku ini sangat direkomendasikan untuk belajar Golang bgi pemula",
		Price:       120000,
		Discount:    20,
		Rating:      4,
	}
	newBook, err := bookRepository.Create(book)
	if err != nil {
		fmt.Println("Error creating data")
		return
	}
	fmt.Printf("Buku berhasil disimpan %v", newBook) */

	// Delete by ID
	/* book, err := bookRepository.FindById(4)
	if err != nil {
		fmt.Println("Error find by ID")
		return
	}
	newBook, err := bookRepository.DeleteById(book, 4)
	if err != nil {
		fmt.Println("Error deleting data")
		return
	}
	fmt.Printf("Buku dengan title %s berhasil dihapus", newBook.Title) */

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
