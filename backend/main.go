package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
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
		c.IndentedJSON(http.StatusOK, books)
	})

	err := r.Run()
	if err != nil {
		fmt.Println("server run error")
	}
}
