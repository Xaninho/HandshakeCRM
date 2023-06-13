package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/handshakeCRM/initializers"
	"github.com/handshakeCRM/models"
)

type CurrencyRequest struct {
	Code         string
	Name         string
	ExchangeRate float64
}

func CurrencyCreate(c *gin.Context) {

	var body CurrencyRequest

	c.Bind(&body)

	currency := models.Currency{
		Code:         body.Code,
		Name:         body.Name,
		ExchangeRate: body.ExchangeRate,
	}

	result := initializers.DB.Create(&currency)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"currency": currency,
	})

}

func CurrencyIndex(c *gin.Context) {

	var currencies []models.Currency
	initializers.DB.Find(&currencies)

	c.JSON(200, gin.H{
		"currencies": currencies,
	})
}

func CurrencyShow(c *gin.Context) {

	id := c.Param("id")

	var currency []models.Currency
	initializers.DB.First(&currency, id)

	c.JSON(200, gin.H{
		"currency": currency,
	})
}

func CurrencyUpdate(c *gin.Context) {

	id := c.Param("id")

	var body CurrencyRequest

	c.Bind(&body)

	var currency models.Currency
	initializers.DB.First(&currency, id)

	initializers.DB.Model(&currency).Updates(models.Currency{
		Code:         body.Code,
		Name:         body.Name,
		ExchangeRate: body.ExchangeRate,
	})

	c.JSON(200, gin.H{
		"currency": currency,
	})

}

func CurrencyDelete(c *gin.Context) {

	id := c.Param("id")

	initializers.DB.Delete(&models.Currency{}, id)

	c.Status(200)
}
