package controllers

import (
	"mygram/initializers"
	"mygram/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePhotos(c *gin.Context){
	user, _ := c.Get("user")
	var Body struct {
		Title string
		Caption string
		Photo_url string
	}
	if c.BindJSON(&Body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to read body",
		})
		return
	}

	user_id := user.(models.User).ID
	photo := models.Photo{Title : Body.Title , Caption: Body.Caption, Photo_url: Body.Photo_url, UserRefer: user_id}
	result := initializers.DB.Create(&photo)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to create photo",
		})
		return
	}
	
	// http status 201
		c.JSON(http.StatusCreated, gin.H{
		"message": photo,

	})

}