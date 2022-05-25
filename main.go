package main

import (
	"log"

	"pustaka-api/book"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// CONNECT DATABASE
	dsn := "root:@tcp(127.0.0.1:3306)/go-pustaka?charset=utf8mb4&parseTime=True&loc=Local"
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB COnnection gagal")
	}

	// AUTO MIGRATE DATABASE
	db.AutoMigrate(&book.Book{})

	// Repository
	bookRepository := book.NewReposiroty(db)

	// // GET FIND ALL
	// books, err := bookRepository.FindAll()

	// for _, book := range books {
	// 	fmt.Println("title: ", book.Title)
	// }

	// Create Book
	book := book.Book{
		Title: "Go Programming JAVA",
		Price: 100000,
		Description:  "This is my first JAVA app",
		Rating: 8,
		Discount: 10,
	}

	bookRepository.Create(book)

	// Router
	router := gin.Default()

	// VERSIONING ROUTE
	// v1 := router.Group("v1")
	// v1.GET("/", rootHandler)

	// API GET
	router.GET("/", handler.RootHandler)
	router.GET("/hello", handler.HelloHandler)
	router.GET("/books/:id", handler.BooksHandler)
	router.GET("/query", handler.QueryHandler)

	//API POST
	router.POST("/books", handler.PostBooksHandler)
	router.Run()
}

