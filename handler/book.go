package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"pustaka-api/book"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (h *bookHandler) GetBooks(c *gin.Context) {
	books, err := h.bookService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var booksResponse []book.BookResponse
	for _, b := range books {
		bookResponse := convertToBookResponse(b)
		booksResponse = append(booksResponse, bookResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data" : booksResponse,
	})
}

func (h *bookHandler) GetBook(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	data_book, err := h.bookService.FindByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	bookResponse := convertToBookResponse(data_book)

	c.JSON(http.StatusOK, gin.H{
		"data" : bookResponse,
	})
}

func (h *bookHandler) CreateBook(c *gin.Context)  {
	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": errorMessages,
			})
			return
	}

	book, err := h.bookService.Create(bookRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data" : book,
	})
}

func convertToBookResponse(bookObj book.Book) book.BookResponse {
	return book.BookResponse {
		ID: bookObj.ID,
		Title: bookObj.Title,
		Description: bookObj.Description,
		Price: bookObj.Price,
		Rating: bookObj.Rating,
		Discount: bookObj.Discount,
	}
}