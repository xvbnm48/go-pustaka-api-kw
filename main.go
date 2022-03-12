package main

import (
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
	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewbookHandler(bookService)

	router := gin.Default()

	v1 := router.Group("/v1")

	// v1.GET("/", bookHandler.RootHandler)
	// v1.GET("/hello", bookHandler.HelloHandler)
	// v1.GET("books/:id/:title", bookHandler.BooksHandler)
	// v1.GET("/query", bookHandler.QueryHandler)
	v1.GET("/books", bookHandler.GetBooks)
	v1.POST("/books", bookHandler.PostBooksHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
	// router.Run(":8080")
}
