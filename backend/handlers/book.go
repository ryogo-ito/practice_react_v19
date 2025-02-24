package bookHandler

import (
	"github.com/gin-gonic/gin"
	bookModel "github.com/ryogo-ito/practice_react_v19/backend/models"
	"net/http"
	"strconv"
)

func GetAll(c *gin.Context) {
	keyword := c.Query("keyword")
	c.IndentedJSON(http.StatusOK, bookModel.GetBooks(keyword))
}

func Create(c *gin.Context) {
	var newBook bookModel.Book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	c.IndentedJSON(http.StatusCreated, newBook.Add())
}

func Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}
	newBook := bookModel.Book{
		ID: id,
	}
	if err2 := c.BindJSON(&newBook); err2 != nil {
		return
	}
	c.IndentedJSON(http.StatusCreated, newBook.Update())
}
