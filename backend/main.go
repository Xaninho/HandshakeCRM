package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/handshakeCRM/controllers"
	"github.com/handshakeCRM/initializers"
	"github.com/handshakeCRM/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {

	r := gin.Default()

	// Set up CORS middleware
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:4500"} // Set your desired origins here
	r.Use(cors.New(config))

	// Users
	r.POST("/user", controllers.UserCreate)
	r.GET("/user", controllers.UserIndex)
	r.GET("/user/:id", controllers.UserShow)
	r.PUT("/user/:id", controllers.UserUpdate)
	r.DELETE("/user/:id", controllers.UserDelete)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	// Companies
	r.POST("/company", controllers.CompanyCreate)
	r.GET("/company", controllers.CompanyIndex)
	r.GET("/company/:id", controllers.CompanyShow)
	r.PUT("/company/:id", controllers.CompanyUpdate)
	r.DELETE("/company/:id", controllers.CompanyDelete)

	// Contacts
	r.POST("/contact", controllers.ContactCreate)
	r.GET("/contact", controllers.ContactIndex)
	r.GET("/contact/:id", controllers.ContactShow)
	r.PUT("/contact/:id", controllers.ContactUpdate)
	r.DELETE("/contact/:id", controllers.ContactDelete)

	// Currencies
	r.POST("/currency", controllers.CurrencyCreate)
	r.GET("/currency", controllers.CurrencyIndex)
	r.GET("/currency/:id", controllers.CurrencyShow)
	r.PUT("/currency/:id", controllers.CurrencyUpdate)
	r.DELETE("/currency/:id", controllers.CurrencyDelete)
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)

	// Countries
	r.POST("/country", controllers.CountryCreate)
	r.GET("/country", controllers.CountryIndex)
	r.GET("/country/:id", controllers.CountryShow)
	r.PUT("/country/:id", controllers.CountryUpdate)
	r.DELETE("/country/:id", controllers.CountryDelete)

	// EnumTypes
	r.POST("/enumType", controllers.EnumTypeCreate)
	r.GET("/enumType", controllers.EnumTypeIndex)
	r.GET("/enumType/:typeOfEnum", controllers.EnumTypeShow)
	r.PUT("/enumType/:id", controllers.EnumTypeUpdate)
	r.DELETE("/enumType/:id", controllers.EnumTypeDelete)

	r.Run()
}
