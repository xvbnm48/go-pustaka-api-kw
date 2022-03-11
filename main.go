package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	book "github.com/xvbnm48/go-pustaka-api-kw/book"
	"github.com/xvbnm48/go-pustaka-api-kw/handler"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	dsn := os.Getenv("DATABASE_URI")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("db connection eror")
	}

	db.AutoMigrate(&book.Book{})
	// book := &book.Book{}
	// book.Title = "Bloom into you"
	// book.Price = 150
	// book.Description = "bloom into you is a manga from japan about a 2 girls "
	// book.Rating = 9

	// err = db.Create(&book).Error
	// if err != nil {
	// 	log.Fatal("db create error")
	// }

	// var book book.Book
	var book book.Book
	// err = db.Debug().First(&book, 2).Error
	// err = db.Debug().Find(&books).Error
	// err = db.Debug().Where("title =?", "Bloom into you").Find(&books).Error
	err = db.Debug().Where("id =?", 1).First(&book).Error
	if err != nil {
		fmt.Println("error find book record ")
	}

	book.Title = "Golang books"
	err = db.Save(&book).Error
	if err != nil {
		fmt.Println("error find book record ")
	}

	// fmt.Println("title:", book.Title)
	// fmt.Println("book object %v", book)

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("books/:id/:title", handler.BooksHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.PostBooksHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
	// router.Run(":8080")
}
