package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"code/database"
	f "code/functions"
	"code/repository"
	. "code/structs"

	"github.com/gin-gonic/gin"
)

func InsertBooks(c *gin.Context) {
	var books Book

	err := c.BindJSON(&books)

	if err != nil {
		c.JSON(http.StatusBadRequest, message{
			Code:    http.StatusBadRequest,
			Message: "Error occurs (51): " + err.Error(),
		})
		return
	}

	if books.Title == "" {
		c.JSON(http.StatusBadRequest, message{
			Code:    http.StatusBadRequest,
			Message: "Error occurs (53): TITLE CANNOT BE NULL",
		})
		return
	}

	fmt.Println(books)

	fmt.Println(books)

	key := c.GetHeader("Authorization")

	err, user := f.AuthLogin(key)

	if err != nil {
		c.JSON(http.StatusBadRequest, message{
			Code:    http.StatusBadRequest,
			Message: "Error occurs (50): " + err.Error(),
		})
		return
	}

	books.CreatedBy = user
	books.ModifiedBy = user
	books.CreatedAt = time.Now()
	books.ModifiedAt = time.Now()

	err = repository.InsertBooks(database.DbConnection, books)

	if err != nil {
		c.JSON(http.StatusBadRequest, message{
			Code:    http.StatusBadRequest,
			Message: "Error occurs (54): " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Successfully created " + books.Title,
	})
}

func GetAllBooks(c *gin.Context) {
	var book []Book
	var books []string

	key := c.GetHeader("Authorization")

	err, user := f.AuthLogin(key)

	if err != nil {
		c.JSON(http.StatusBadRequest, message{
			Code:    http.StatusBadRequest,
			Message: "Error occurs (14): " + err.Error(),
		})
		return
	}

	fmt.Println(user)

	err = repository.GetAllBooks(database.DbConnection, &book)
	if err != nil {
		c.JSON(http.StatusBadRequest, message{
			Code:    http.StatusBadRequest,
			Message: "Error occurs: (19)" + err.Error(),
		})
		return
	}

	for _, val := range book {
		books = append(books, val.Title)
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"code": http.StatusOK,
		"data": books,
	})

}

func GetBookDet(c *gin.Context) {
	var book Book

	id, _ := strconv.Atoi(c.Param("id"))

	book.ID = id

	key := c.GetHeader("Authorization")

	err, _ := f.AuthLogin(key)

	if err != nil {
		c.JSON(http.StatusBadRequest, message{
			Code:    http.StatusBadRequest,
			Message: "Error occurs (20): " + err.Error(),
		})
		return
	}

	err = repository.GetBooksDet(database.DbConnection, &book)

	if err != nil {
		c.JSON(http.StatusBadRequest, message{
			Code:    http.StatusBadRequest,
			Message: "Error occurs (32): " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"code": http.StatusOK,
		"data": book,
	})
}

func DeleteBook(c *gin.Context) {
	var book Book

	id, _ := strconv.Atoi(c.Param("id"))

	book.ID = id

	key := c.GetHeader("Authorization")

	err, _ := f.AuthLogin(key)

	if err != nil {
		c.JSON(http.StatusBadRequest, message{
			Code:    http.StatusBadRequest,
			Message: "Error occurs (20): " + err.Error(),
		})
		return
	}

	err = repository.DeleteBook(database.DbConnection, &book)
	if err != nil {
		c.JSON(http.StatusBadRequest, message{
			Code:    http.StatusBadRequest,
			Message: "Error occurs (41): " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, message{
		Code:    http.StatusOK,
		Message: fmt.Sprintf("Successfully delete %d", book.ID),
	})
}
