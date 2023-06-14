package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/handshakeCRM/initializers"
	"github.com/handshakeCRM/models"
)

type CountryRequest struct {
	Name        string
	CurrencyId  int
	ContinentId int
}

func CountryCreate(c *gin.Context) {

	var body CountryRequest

	c.Bind(&body)

	country := models.Country{
		Name:       body.Name,
		CurrencyId: body.CurrencyId,
	}

	result := initializers.DB.Create(&country)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"country": country,
	})

}

func CountryIndex(c *gin.Context) {

	var countries []models.Country
	initializers.DB.Find(&countries)

	c.JSON(200, gin.H{
		"countries": countries,
	})
}

func CountryShow(c *gin.Context) {

	id := c.Param("id")

	var country []models.Country
	initializers.DB.First(&country, id)

	c.JSON(200, gin.H{
		"country": country,
	})
}

func CountryUpdate(c *gin.Context) {

	id := c.Param("id")

	var body CountryRequest
	c.Bind(&body)

	var country models.Country
	initializers.DB.First(&country, id)

	initializers.DB.Model(&country).Updates(models.Country{
		Name:       body.Name,
		CurrencyId: body.CurrencyId,
	})

	c.JSON(200, gin.H{
		"country": country,
	})

}

func CountryDelete(c *gin.Context) {

	id := c.Param("id")

	initializers.DB.Delete(&models.Country{}, id)

	c.Status(200)
}
