package controllers

import (
	"mygram/initializers"
	"mygram/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var Body struct {
		Username string
		Email string
		Password string
		Age int
	}
	if c.BindJSON(&Body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to read body",
		})
		return
	}
	hash,err := bcrypt.GenerateFromPassword([]byte(Body.Password), 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to hash password",
		})
		return
	} 
	user := models.User{Username:Body.Username, Email: Body.Email, Password: string(hash), Age: Body.Age}
	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to create user",
		})
		return
	}
	// http status 201
	c.JSON(http.StatusCreated, gin.H{
		"age" : user.Age,
		"email" : user.Email,
		"id" : user.ID,
		"username" : user.Username,
	})
}