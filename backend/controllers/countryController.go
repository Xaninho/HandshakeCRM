package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/handshakeCRM/initializers"
	"github.com/handshakeCRM/models"
)

type CountryRequest struct {
	Name        string `json:"name"`
	CurrencyId  int    `json:"currencyId"`
	ContinentId int    `json:"continentId"`
}

func CountryCreate(c *gin.Context) {

	// Get data off req body
	var body CountryRequest

	c.Bind(&body)

	// Create a post
	country := models.Country{
		Name:        body.Name,
		CurrencyId:  body.CurrencyId,
		ContinentId: body.ContinentId,
	}

	result := initializers.DB.Create(&country)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return it
	c.JSON(200, gin.H{
		"country": country,
	})

}

func CountryIndex(c *gin.Context) {
	//Get the posts
	var countries []models.Country
	initializers.DB.Find(&countries)

	//Respond with them
	c.JSON(200, gin.H{
		"countries": countries,
	})
}

func CountryShow(c *gin.Context) {
	//get id off url
	id := c.Param("id")

	//Get the posts
	var country []models.Country
	initializers.DB.First(&country, id)

	//Respond with them
	c.JSON(200, gin.H{
		"country": country,
	})
}

func CountryUpdate(c *gin.Context) {
	// Get the id off the url
	id := c.Param("id")

	// Get the data off req body
	var body CountryRequest
	c.Bind(&body)

	// Find the post were updating
	var country models.Country
	initializers.DB.First(&country, id)

	// Update it
	initializers.DB.Model(&country).Updates(models.Country{
		Name:        body.Name,
		CurrencyId:  body.CurrencyId,
		ContinentId: body.ContinentId,
	})

	// Respond with it
	c.JSON(200, gin.H{
		"country": country,
	})

}

func CountryDelete(c *gin.Context) {

	// Get the id off the url
	id := c.Param("id")

	// Delete the company
	initializers.DB.Delete(&models.Country{}, id)

	// Respond
	c.Status(200)
}
