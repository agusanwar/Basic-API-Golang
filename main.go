package main

import (
	"log"
	"pustaka-api/books"
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
	db.AutoMigrate(&books.Book{})


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

