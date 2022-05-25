package main

import (
	"fmt"
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

	// // CREATE DATA
	// books := books.Book{}
	// books.Title = "Pemrograman Java Script"
	// books.Price = 300000
	// books.Description = "Pemrograman Java Script For Beginners to Intermediate"
	// books.Rating = 10
	// books.Discount = 10

	// // SAVE DATA
	// err = db.Create(&books).Error
	// if err != nil {
	//  println("Error Create Field books")
	// }

	//  // READ DATA
	// // var books books.Book
	// var findBooks []books.Book
	// err = db.Debug().Where("Title = ?", "Pemrograman Go").Find(&findBooks).Error
	// if err != nil {
	//  println("Error Finding Field books")
	// }
	// for _, book := range findBooks {
	// 	fmt.Println("Title:", book.Title)
	// 	fmt.Println("books object: ", book)
	// }

	//  // UPDATE DATA
	// var books books.Book
	// err = db.Debug().Where("id = ?", 1).First(&books).Error
	// if err != nil {
	//  println("Error Finding Field books")
	// }
	// 	fmt.Println("Error Find Data")

	// // SAVE UPDATE
	// books.Title = "Pemrograman Go-Lang"
	// err = db.Debug().Save(&books).Error
	// if err != nil {
	//  println("Error Finding Field books")
	// }
	// 	fmt.Println("Error Update Data")

	// DELETE DATA
	var books books.Book
	err = db.Debug().Where("id = ?", 3).First(&books).Error
	if err != nil {
		fmt.Println("Error Find Data")
	}
		

	//  PROSES DELETE
	err = db.Delete(&books).Error
	if err != nil {
	 	fmt.Println("Error Delete Data")
	}
		

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

