package main

import (
	"github.com/gin-gonic/gin"
	"github.com/handshakeCRM/initializers"
	"github.com/handshakeCRM/controllers"
)

func init() {
	initializers.LoadEnvVariables();
	initializers.ConnectToDb()
}

func main() {

	r := gin.Default()

	r.POST("/company", controllers.CompanyCreate)
	r.GET("/company", controllers.CompanyIndex)
	r.GET("/company/:id", controllers.CompanyShow)
	r.PUT("/company/:id", controllers.CompanyUpdate)
	r.DELETE("/company/:id", controllers.CompanyDelete)

	r.Run()
}