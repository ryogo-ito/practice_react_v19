package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// TODO 移動する
type Book struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

var books = []Book{
	{ID: 1, Name: "React入門", Status: "在庫あり"},
	{ID: 2, Name: "TypeScript入門", Status: "貸出中"},
	{ID: 3, Name: "Next.js入門", Status: "返却済"},
}

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			// TODO 調整する
			"http://localhost:5173",
		},
	}))

	r.GET("/books", func(c *gin.Context) {
		keyword := c.Query("keyword")
		c.ContentType()

		fmt.Println(keyword)

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
	})

	r.POST("/books", func(c *gin.Context) {
		var newBook Book

		if err := c.BindJSON(&newBook); err != nil {
			return
		}

		newBook.ID = len(books) + 1
		newBook.Status = "在庫あり"

		books = append(books, newBook)
		c.IndentedJSON(http.StatusCreated, newBook)
	})

	err := r.Run()
	if err != nil {
		fmt.Println("server run error")
	}
}
