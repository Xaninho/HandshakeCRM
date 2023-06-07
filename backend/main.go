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
	r.GET("/validate", controllers.Validate)
	r.POST("/login", controllers.Login)

	r.POST("/user", middleware.RequireAuth, controllers.UserCreate)
	r.GET("/user", middleware.RequireAuth, controllers.UserIndex)
	r.GET("/user/:id", middleware.RequireAuth, controllers.UserShow)
	r.PUT("/user/:id", middleware.RequireAuth, controllers.UserUpdate)
	r.DELETE("/user/:id", middleware.RequireAuth, controllers.UserDelete)

	// Companies
	r.POST("/company", middleware.RequireAuth, controllers.CompanyCreate)
	r.GET("/company", middleware.RequireAuth, controllers.CompanyIndex)
	r.GET("/company/:id", middleware.RequireAuth, controllers.CompanyShow)
	r.PUT("/company/:id", middleware.RequireAuth, controllers.CompanyUpdate)
	r.DELETE("/company/:id", middleware.RequireAuth, controllers.CompanyDelete)

	// Contacts
	r.POST("/contact", middleware.RequireAuth, controllers.ContactCreate)
	r.GET("/contact", middleware.RequireAuth, controllers.ContactIndex)
	r.GET("/contact/:id", middleware.RequireAuth, controllers.ContactShow)
	r.PUT("/contact/:id", middleware.RequireAuth, controllers.ContactUpdate)
	r.DELETE("/contact/:id", middleware.RequireAuth, controllers.ContactDelete)

	// Currencies
	r.POST("/currency", middleware.RequireAuth, controllers.CurrencyCreate)
	r.GET("/currency", middleware.RequireAuth, controllers.CurrencyIndex)
	r.GET("/currency/:id", middleware.RequireAuth, controllers.CurrencyShow)
	r.PUT("/currency/:id", middleware.RequireAuth, controllers.CurrencyUpdate)
	r.DELETE("/currency/:id", middleware.RequireAuth, controllers.CurrencyDelete)
	r.POST("/signup", middleware.RequireAuth, controllers.Signup)

	// Countries
	r.POST("/country", middleware.RequireAuth, controllers.CountryCreate)
	r.GET("/country", middleware.RequireAuth, controllers.CountryIndex)
	r.GET("/country/:id", middleware.RequireAuth, controllers.CountryShow)
	r.PUT("/country/:id", middleware.RequireAuth, controllers.CountryUpdate)
	r.DELETE("/country/:id", middleware.RequireAuth, controllers.CountryDelete)

	// EnumTypes
	r.POST("/enumType", middleware.RequireAuth, controllers.EnumTypeCreate)
	r.GET("/enumType", middleware.RequireAuth, controllers.EnumTypeIndex)
	r.GET("/enumType/:typeOfEnum", middleware.RequireAuth, controllers.EnumTypeShow)
	r.PUT("/enumType/:id", middleware.RequireAuth, controllers.EnumTypeUpdate)
	r.DELETE("/enumType/:id", middleware.RequireAuth, controllers.EnumTypeDelete)

	r.Run()
}
