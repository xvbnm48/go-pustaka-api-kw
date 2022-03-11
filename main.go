package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/xvbnm48/go-pustaka-api-kw/book"
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
	book := &book.Book{}
	book.Title = "Bloom into you"
	book.Price = 150
	book.Description = "bloom into you is a manga from japan about a 2 girls "
	book.Rating = 9

	err = db.Create(&book).Error
	if err != nil {
		log.Fatal("db create error")
	}

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("books/:id/:title", handler.BooksHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.PostBooksHandler)

	router.Run(":8080")
}
