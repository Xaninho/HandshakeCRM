package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/handshakeCRM/models"
	"github.com/handshakeCRM/initializers"
)

type CompanyRequest struct {
	NIF                 	int `json:"nif"`
	Name                	string `json:"name"`
	Currency          		models.Currency `json:"currency"`
	HasElectronicBilling	bool `json:"hasElectronicBilling"`
	PhotoURL            	string `json:"photoURL"`
	Country           		models.Country `json:"country"`
	PostalCode          	string `json:"postalCode"`
	Address             	string `json:"address"`
	AssignedAgent     		models.Agent `json:"assignedAgent"`
	State             		models.EnumType `json:"state"`
	Notes               	string `json:"notes"`
}

func CompanyCreate (c *gin.Context) {

	// Get data off req body
	var body CompanyRequest
	c.Bind(&body)

	// Create a post
	company := models.Company{	
		NIF: body.NIF,
		Name: body.Name,
		Currency: body.Currency,
		HasElectronicBilling: body.HasElectronicBilling,
		PhotoURL: body.PhotoURL,
		Country: body.Country,
		PostalCode: body.PostalCode,
		Address: body.Address,
		State: body.State,
		Notes: body.Notes,
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

func CompanyIndex (c *gin.Context) {

	//Get the posts
	var companies []models.Company
	initializers.DB.Find(&companies)

	//Respond with them
	c.JSON(200, gin.H{
		"companies": companies,
	})
}

func CompanyShow (c *gin.Context) {

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

func CompanyUpdate (c *gin.Context) {

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
		NIF: body.NIF,
		Name: body.Name,
		Currency: body.Currency,
		HasElectronicBilling: body.HasElectronicBilling,
		PhotoURL: body.PhotoURL,
		Country: body.Country,
		PostalCode: body.PostalCode,
		Address: body.Address,
		State: body.State,
		Notes: body.Notes,
	})

	// Respond with it
	c.JSON(200, gin.H{
		"company": company,
	})

}

func CompanyDelete (c * gin.Context) {

	// Get the id off the url
	id := c.Param("id")	
	
	// Delete the company
	initializers.DB.Delete(&models.Company{}, id)

	// Respond
	c.Status(200)
}