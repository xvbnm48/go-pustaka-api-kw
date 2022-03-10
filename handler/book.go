package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/xvbnm48/go-pustaka-api-kw/book"
)

func RootHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"Name":    "sakura endo",
		"Age":     20,
		"Address": "Tokyo",
	})
}

func HelloHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"Name":    "sakura endo",
		"Age":     20,
		"Address": "Tokyo",
	})
}

func BooksHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")

	c.JSON(200, gin.H{
		"id":    id,
		"title": title,
	})
}

func QueryHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")

	c.JSON(200, gin.H{
		"title": title,
		"price": price,
	})

}

func PostBooksHandler(c *gin.Context) {
	var bookinput book.BookInput
	err := c.ShouldBindJSON(&bookinput)

	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("error on field %s , condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
			// c.JSON(400, errorMessage)
			// fmt.Println(err)
			// return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})

		return

	}

	c.JSON(200, gin.H{
		"title": bookinput.Title,
		"price": bookinput.Price,
		// "sub_title": bookinput.SubTitle,
	})

}
