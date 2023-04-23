package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/handshakeCRM/models"
	"github.com/handshakeCRM/initializers"
)

type CurrencyRequest struct {
    Code         	string `json:"code"`
	Name         	string `json:"name"`
	ExchangeRate 	float64 `json:"exchangeRate"`
}

func CurrencyCreate (c *gin.Context) {

	// Get data off req body
	var body CurrencyRequest

	c.Bind(&body)

	// Create a post
	currency := models.Currency{	
		Code: body.Code,
		Name: body.Name,
		ExchangeRate: body.ExchangeRate,
	}
	
	result := initializers.DB.Create(&currency)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return it
	c.JSON(200, gin.H{
		"currency": currency,
	})

}

func CurrencyIndex (c *gin.Context) {

	//Get the posts
	var currencies []models.Currency
	initializers.DB.Find(&currencies)

	//Respond with them
	c.JSON(200, gin.H{
		"currencies": currencies,
	})
}

func CurrencyShow (c *gin.Context) {

	//get id off url
	id := c.Param("id")

	//Get the posts
	var currency []models.Currency
	initializers.DB.First(&currency, id)

	//Respond with them
	c.JSON(200, gin.H{
		"currency": currency,
	})
}

func CurrencyUpdate (c *gin.Context) {

	// Get the id off the url
	id := c.Param("id")

	// Get the data off req body
	var body CurrencyRequest

	c.Bind(&body)

	// Find the post were updating
	var currency models.Currency
	initializers.DB.First(&currency, id)

	// Update it
	initializers.DB.Model(&currency).Updates(models.Currency{
		Code: body.Code,
		Name: body.Name,
		ExchangeRate: body.ExchangeRate,
	})

	// Respond with it
	c.JSON(200, gin.H{
		"currency": currency,
	})

}

func CurrencyDelete (c * gin.Context) {

	// Get the id off the url
	id := c.Param("id")	
	
	// Delete the company
	initializers.DB.Delete(&models.Currency{}, id)

	// Respond
	c.Status(200)
}