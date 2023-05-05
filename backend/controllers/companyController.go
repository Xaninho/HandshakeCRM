package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/handshakeCRM/initializers"
	"github.com/handshakeCRM/models"
)

type CompanyRequest struct {
	NIF                  int    `json:"nif"`
	Name                 string `json:"name"`
	CurrencyId           int    `json:"currencyId"`
	HasElectronicBilling bool   `json:"hasElectronicBilling"`
	PhotoURL             string `json:"photoURL"`
	CountryId            int    `json:"countryId"`
	PostalCode           string `json:"postalCode"`
	Address              string `json:"address"`
	StateId              int    `json:"stateId"`
	Notes                string `json:"notes"`
}

func CompanyCreate(c *gin.Context) {

	// Get data off req body
	var body CompanyRequest
	c.Bind(&body)

	// Create a post
	company := models.Company{
		NIF:                  body.NIF,
		Name:                 body.Name,
		CurrencyId:           body.CurrencyId,
		HasElectronicBilling: body.HasElectronicBilling,
		PhotoURL:             body.PhotoURL,
		CountryId:            body.CountryId,
		PostalCode:           body.PostalCode,
		Address:              body.Address,
		StateId:              body.StateId,
		Notes:                body.Notes,
	}

	result := initializers.DB.Create(&company)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return it
	c.JSON(200, gin.H{
		"company": company,
	})

}

func CompanyIndex(c *gin.Context) {

	//Get the posts
	var companies []models.Company
	initializers.DB.Preload("Currency").Preload("Country").Find(&companies)

	//Respond with them
	c.JSON(200, gin.H{
		"companies": companies,
	})
}

func CompanyShow(c *gin.Context) {

	//get id off url
	id := c.Param("id")

	//Get the posts
	var company []models.Company
	initializers.DB.First(&company, id)

	//Respond with them
	c.JSON(200, gin.H{
		"companies": company,
	})
}

func CompanyUpdate(c *gin.Context) {

	// Get the id off the url
	id := c.Param("id")

	// Get the data off req body
	var body CompanyRequest
	c.Bind(&body)

	// Find the post were updating
	var company models.Company
	initializers.DB.First(&company, id)

	// Update it

	initializers.DB.Model(&company).Updates(models.Company{
		NIF:                  body.NIF,
		Name:                 body.Name,
		CurrencyId:           body.CurrencyId,
		HasElectronicBilling: body.HasElectronicBilling,
		PhotoURL:             body.PhotoURL,
		CountryId:            body.CountryId,
		PostalCode:           body.PostalCode,
		Address:              body.Address,
		StateId:              body.StateId,
		Notes:                body.Notes,
	})

	// Respond with it
	c.JSON(200, gin.H{
		"company": company,
	})

}

func CompanyDelete(c *gin.Context) {

	// Get the id off the url
	id := c.Param("id")

	// Delete the company
	initializers.DB.Delete(&models.Company{}, id)

	// Respond
	c.Status(200)
}
