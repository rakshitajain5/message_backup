package main

import (

	"gopkg.in/gin-gonic/gin.v1"
	"message_backup/controllers"
)

func main() {


	r := gin.New()

	// Global middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())


	r.GET("/",func(c *gin.Context){
		Index(c)
	})

	r.GET("/hello", func(c *gin.Context){
		Hello(c)
	})

	r.POST("/jcm/messages/backup", func(c *gin.Context){controllers.MsgBackup(c)})
	r.Run(":8080")

}

func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "server running",
	})
}

func Hello( c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello",
	})
}

