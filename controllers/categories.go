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

func GetAllCategories(c *gin.Context) {
	var categories []Category
	var cats []string

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

	err = repository.GetAllCategories(database.DbConnection, &categories)
	if err != nil {
		c.JSON(http.StatusBadRequest, message{
			Code:    http.StatusBadRequest,
			Message: "Error occurs: (19)" + err.Error(),
		})
		return
	}

	for _, val := range categories {
		cats = append(cats, val.Name)
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"code": http.StatusOK,
		"data": cats,
	})

}

func InsertCategory(c *gin.Context) {
	var categories Category

	err := c.BindJSON(&categories)

	if err != nil {
		c.JSON(http.StatusBadRequest, message{
			Code:    http.StatusBadRequest,
			Message: "Error occurs (25): " + err.Error(),
		})
		return
	}

	if categories.Name == "" {
		c.JSON(http.StatusBadRequest, message{
			Code:    http.StatusBadRequest,
			Message: "Error occurs (23): NAME CANNOT BE NULL",
		})
		return
	}

	fmt.Println(categories)

	key := c.GetHeader("Authorization")

	err, user := f.AuthLogin(key)

	if err != nil {
		c.JSON(http.StatusBadRequest, message{
			Code:    http.StatusBadRequest,
			Message: "Error occurs (20): " + err.Error(),
		})
		return
	}

	categories.CreatedBy = user
	categories.ModifiedBy = user
	categories.CreatedAt = time.Now()
	categories.ModifiedAt = time.Now()

	fmt.Println(categories)

	err = repository.InsertCategories(database.DbConnection, categories)

	if err != nil {
		c.JSON(http.StatusBadRequest, message{
			Code:    http.StatusBadRequest,
			Message: "Error occurs (21): " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Category created " + categories.Name,
	})

}

func GetCategoryDetails(c *gin.Context) {
	var category Category

	id, _ := strconv.Atoi(c.Param("id"))

	category.ID = id

	key := c.GetHeader("Authorization")

	err, _ := f.AuthLogin(key)

	if err != nil {
		c.JSON(http.StatusBadRequest, message{
			Code:    http.StatusBadRequest,
			Message: "Error occurs (20): " + err.Error(),
		})
		return
	}

	err = repository.GetCategoryDet(database.DbConnection, &category)

	if err != nil {
		c.JSON(http.StatusBadRequest, message{
			Code:    http.StatusBadRequest,
			Message: "Error occurs (32): " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"code": http.StatusOK,
		"data": category,
	})
}

func DeleteCat(c *gin.Context) {
	var category Category

	id, _ := strconv.Atoi(c.Param("id"))

	category.ID = id

	key := c.GetHeader("Authorization")

	err, _ := f.AuthLogin(key)

	if err != nil {
		c.JSON(http.StatusBadRequest, message{
			Code:    http.StatusBadRequest,
			Message: "Error occurs (20): " + err.Error(),
		})
		return
	}

	err = repository.DeleteCategories(database.DbConnection, &category)
	if err != nil {
		c.JSON(http.StatusBadRequest, message{
			Code:    http.StatusBadRequest,
			Message: "Error occurs (41): " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, message{
		Code:    http.StatusOK,
		Message: fmt.Sprintf("Successfully delete %d", category.ID),
	})
}

func GetCategoryBooks(c *gin.Context) {
	var categories Category
	var books []string

	id, _ := strconv.Atoi(c.Param("id"))

	categories.ID = id

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

	err = repository.GetCatBook(database.DbConnection, &categories, &books)
	if err != nil {
		c.JSON(http.StatusBadRequest, message{
			Code:    http.StatusBadRequest,
			Message: "Error occurs: (19)" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"code": http.StatusOK,
		"data": books,
	})
}
