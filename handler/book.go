package handler

import (
	"fmt"
	"net/http"
	"strconv"

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
		bookResponse := convertBookToResponse(b)
		booksResponse = append(booksResponse, bookResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": booksResponse,
	})
}

func (h *bookHandler) Getbook(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	b, err := h.bookService.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"errors": "Book not found",
		})
		return
	}

	book := convertBookToResponse(b)
	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

func (h *bookHandler) CreateBook(c *gin.Context) {
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
		"data": convertBookToResponse(book),
		// "sub_title": bookinput.SubTitle,
	})

}

func (h *bookHandler) UpdateBook(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	var book book.BookRequest
	err := c.ShouldBindJSON(&book)
	if err != nil {
		errorMessagges := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("error on field %s , condition %s", e.Field(), e.ActualTag())
			errorMessagges = append(errorMessagges, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessagges,
		})

		return
	}

	bookUpdate, err := h.bookService.Update(id, book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errors": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": convertBookToResponse(bookUpdate),
	})

}

func (h *bookHandler) DeleteBook(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	deleteBook, err := h.bookService.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": deleteBook,
	})

}

func convertBookToResponse(b book.Book) book.BookResponse {
	return book.BookResponse{
		ID:          b.ID,
		Title:       b.Title,
		Description: b.Description,
		Price:       b.Price,
		Rating:      b.Rating,
		Discount:    b.Discount,
	}
}
