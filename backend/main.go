package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	bookHandler "github.com/ryogo-ito/practice_react_v19/backend/handlers"
)

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			// TODO 調整する
			"http://localhost:5173",
		},
		AllowMethods: []string{
			"GET", "POST", "PUT", "DELETE",
		},
		AllowHeaders: []string{
			"Content-Type",
		},
	}))

	r.GET("/books", bookHandler.GetBooks)
	r.POST("/books", bookHandler.CreateBook)
	r.PUT("/books/:id", bookHandler.UpdateBook)

	err := r.Run()
	if err != nil {
		fmt.Println("server run error")
	}
}
