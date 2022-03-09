package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)
	router.GET("/hello", helloHandler)
	router.GET("books/:id/:title", booksHandler)
	router.GET("/query", queryHandler)
	router.POST("/books", postBooksHandler)

	router.Run(":8080")
}

func rootHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"Name":    "sakura endo",
		"Age":     20,
		"Address": "Tokyo",
	})
}

func helloHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"Name":    "sakura endo",
		"Age":     20,
		"Address": "Tokyo",
	})
}

func booksHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")

	c.JSON(200, gin.H{
		"id":    id,
		"title": title,
	})
}

func queryHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")

	c.JSON(200, gin.H{
		"title": title,
		"price": price,
	})

}

type BookInput struct {
	Title string `json:"title" binding:"required"`
	Price int    `json:"price" binding:"required|number"`
	// SubTitle string `json:"sub_title"`
}

func postBooksHandler(c *gin.Context) {
	var bookinput BookInput
	err := c.ShouldBindJSON(&bookinput)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		fmt.Println(err)
		return
	}

	c.JSON(200, gin.H{
		"title": bookinput.Title,
		"price": bookinput.Price,
		// "sub_title": bookinput.SubTitle,
	})

}
