package bookHandler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

// TODO 移動する
type Book struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

// TODO DBに移動する
var books = []Book{
	{ID: 1, Name: "React入門", Status: "在庫あり"},
	{ID: 2, Name: "TypeScript入門", Status: "貸出中"},
	{ID: 3, Name: "Next.js入門", Status: "返却済"},
}

func GetAll(c *gin.Context) {
	keyword := c.Query("keyword")

	if keyword == "" {
		c.IndentedJSON(http.StatusOK, books)
		return
	}

	var filteredBooks = make([]Book, 0)
	for _, book := range books {
		if strings.Contains(book.Name, keyword) {
			filteredBooks = append(filteredBooks, book)
		}
	}

	c.IndentedJSON(http.StatusOK, filteredBooks)
}

func Create(c *gin.Context) {
	var newBook Book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	newBook.ID = len(books) + 1
	newBook.Status = "在庫あり"

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}

	var newBook Book
	if err2 := c.BindJSON(&newBook); err2 != nil {
		return
	}

	for _, book := range books {
		if book.ID == id {
			newBook.Name = book.Name
			newBook.ID = book.ID
		}
	}

	for _, book := range books {
		if book.ID == newBook.ID {
			book = newBook
		}
	}

	c.IndentedJSON(http.StatusCreated, newBook)
}
