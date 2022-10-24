package controllers

import (
	"mygram/initializers"
	"mygram/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
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

func Login(c *gin.Context){
	var Body struct {
		Email string
		Password string
	}

	if c.BindJSON(&Body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to read body",
		})
		return
	}

	var user models.User

	initializers.DB.First(&user, "email = ?", Body.Email)

	if user.ID == 0{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid email or password",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(Body.Password))

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid email or password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to generate token",
		})
		return
	}

	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("token", tokenString, 60 * 60 * 24 * 30, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"jwt": tokenString,
	})

}

func Validate(c *gin.Context){
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}