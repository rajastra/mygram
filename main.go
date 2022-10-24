package main

import (
	"mygram/controllers"
	"mygram/initializers"
	"mygram/middleware"

	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncToDb()
}
// compiledaemon --command="./mygram"


func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	// user
	r.POST("/users/register", controllers.Register)
	r.POST("/users/login", controllers.Login)
	r.PUT("/users", middleware.RequireAuth, controllers.UpdateUser)
	r.DELETE("/users", middleware.RequireAuth, controllers.DeleteUser)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	// photo
	r.POST("/photos", middleware.RequireAuth, controllers.CreatePhotos)
	r.GET("/photos", middleware.RequireAuth , controllers.GetPhotos)
	r.Run()
}