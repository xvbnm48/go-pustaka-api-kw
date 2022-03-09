package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)
	router.GET("/hello", helloHandler)
	router.GET("books/:id/:title", booksHandler)
	router.GET("/query", queryHandler)

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
