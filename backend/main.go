package main

import (
	"github.com/gin-gonic/gin"
	"github.com/handshakeCRM/initializers"
)

func init() {
	initializers.LoadEnvVariables();
	initializers.ConnectToDb()
}

func main() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong2",
		})
	})

	r.Run()
}