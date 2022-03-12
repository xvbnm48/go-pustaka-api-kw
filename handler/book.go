package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/xvbnm48/go-pustaka-api-kw/book"
)

type bookHandler struct {
	bookService book.Service
}

func NewbookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService: bookService}
}

func (h *bookHandler) GetBooks(c *gin.Context) {
	books, err := h.bookService.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errors": err,
		})
		return
	}

	var booksResponse []book.BookResponse

	for _, b := range books {
		bookResponse := book.BookResponse{
			ID:          b.ID,
			Title:       b.Title,
			Description: b.Description,
			Price:       b.Price,
			Rating:      b.Rating,
			Discount:    b.Discount,
		}
		booksResponse = append(booksResponse, bookResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": booksResponse,
	})
}

func (h *bookHandler) PostBooksHandler(c *gin.Context) {
	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest)

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

	book, err := h.bookService.Create(bookRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errors": err,
		})
	}

	c.JSON(200, gin.H{
		"data": book,
		// "sub_title": bookinput.SubTitle,
	})

}
