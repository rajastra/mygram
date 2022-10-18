package main

import (
	"mygram/controllers"
	"mygram/initializers"

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
	r.POST("/register", controllers.Register)
	r.Run()
}