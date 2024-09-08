package controllers

import (
	"fmt"
	"net/http"

	"code/database"
	f "code/functions"
	"code/repository"
	. "code/structs"

	"github.com/gin-gonic/gin"
)

type message struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func RegisterUser(c *gin.Context) {
	var user User

	err := c.BindJSON(&user)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(user)

	Encrypted, err := f.PasswordGenerator(user.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, message{
			Code:    http.StatusBadRequest,
			Message: "Error occurs (11): " + err.Error(),
		})
		return
	}

	user.Password = Encrypted

	err = repository.InsertUser(database.DbConnection, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, message{
			Code:    http.StatusBadRequest,
			Message: "Error occurs: (12)" + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, message{
			Code:    http.StatusOK,
			Message: "User Created : " + user.Username,
		})
	}
}

func LoginUser(c *gin.Context) {
	var user User

	err := c.BindJSON(&user)

	if err != nil {
		fmt.Println(err)
	}

	EncryptedPass, err := f.PasswordGenerator(user.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, message{
			Code:    http.StatusBadRequest,
			Message: "Error occurs (15): " + err.Error(),
		})
		return
	}

	err = repository.GetUser(database.DbConnection, &user, EncryptedPass)
	if err != nil {
		c.JSON(http.StatusBadRequest, message{
			Code:    http.StatusBadRequest,
			Message: "Error occurs (16): " + err.Error(),
		})
		return
	}

	res, errors := f.EncodeJWT(map[string]interface{}{
		"data":    user,
		"isLogin": true,
	})

	if errors != nil {
		c.JSON(http.StatusBadRequest, message{
			Code:    http.StatusBadRequest,
			Message: "Error occurs (17): " + err.Error(),
		})
		return

	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"code":  http.StatusOK,
		"token": res,
	})
}
