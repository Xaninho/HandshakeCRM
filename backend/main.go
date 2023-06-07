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

	// Companies
	r.POST("/company", controllers.CompanyCreate)
	r.GET("/company", controllers.CompanyIndex)
	r.GET("/company/:id", controllers.CompanyShow)
	r.PUT("/company/:id", controllers.CompanyUpdate)
	r.DELETE("/company/:id", controllers.CompanyDelete)

	// Clients
	r.POST("/client", controllers.ClientCreate)
	r.GET("/client", controllers.ClientIndex)
	r.GET("/client/:id", controllers.ClientShow)
	r.PUT("/client/:id", controllers.ClientUpdate)
	r.DELETE("/client/:id", controllers.ClientDelete)

	// Agents
	r.POST("/agent", controllers.AgentCreate)
	r.GET("/agent", controllers.AgentIndex)
	r.GET("/agent/:id", controllers.AgentShow)
	r.PUT("/agent/:id", controllers.AgentUpdate)
	r.DELETE("/agent/:id", controllers.AgentDelete)

	// Currencies
	r.POST("/currency", controllers.CurrencyCreate)
	r.GET("/currency", controllers.CurrencyIndex)
	r.GET("/currency/:id", controllers.CurrencyShow)
	r.PUT("/currency/:id", controllers.CurrencyUpdate)
	r.DELETE("/currency/:id", controllers.CurrencyDelete)
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

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
