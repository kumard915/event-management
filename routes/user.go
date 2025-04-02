package routes

import (
	"fmt"
	"net/http"

	"example.com/restapi/models"
	"example.com/restapi/utils"
	"github.com/gin-gonic/gin"
)

func signup(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
		return
	}

	err = user.Save()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not save user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user created"})
}

func login(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)

	fmt.Println("Error:", err)
	fmt.Println("User:", user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
		return
	}

	err = user.ValidateCredentials()
	fmt.Println("Login Error:", err)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}
	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not authenticate user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Login successful!", "token": token})
}
