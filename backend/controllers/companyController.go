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
	CountryId            int    `json:"countryId"`
	PostalCode           string `json:"postalCode"`
	Address              string `json:"address"`
	Notes                string `json:"notes"`
}

func CompanyCreate(c *gin.Context) {

	var body CompanyRequest
	c.Bind(&body)

	company := models.Company{
		NIF:                  body.NIF,
		Name:                 body.Name,
		CurrencyId:           body.CurrencyId,
		HasElectronicBilling: body.HasElectronicBilling,
		CountryId:            body.CountryId,
		PostalCode:           body.PostalCode,
		Address:              body.Address,
		Notes:                body.Notes,
	}

	result := initializers.DB.Create(&company)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"company": company,
	})

}

func CompanyIndex(c *gin.Context) {

	var companies []models.Company
	initializers.DB.Preload("Currency").Preload("Country").Find(&companies)

	c.JSON(200, gin.H{
		"companies": companies,
	})
}

func CompanyShow(c *gin.Context) {

	id := c.Param("id")

	var company []models.Company
	initializers.DB.First(&company, id)

	c.JSON(200, gin.H{
		"companies": company,
	})
}

func CompanyUpdate(c *gin.Context) {

	id := c.Param("id")

	var body CompanyRequest
	c.Bind(&body)

	var company models.Company
	initializers.DB.First(&company, id)

	initializers.DB.Model(&company).Updates(models.Company{
		NIF:                  body.NIF,
		Name:                 body.Name,
		CurrencyId:           body.CurrencyId,
		HasElectronicBilling: body.HasElectronicBilling,
		CountryId:            body.CountryId,
		PostalCode:           body.PostalCode,
		Address:              body.Address,
		Notes:                body.Notes,
	})

	c.JSON(200, gin.H{
		"company": company,
	})

}

func CompanyDelete(c *gin.Context) {

	id := c.Param("id")

	initializers.DB.Delete(&models.Company{}, id)

	c.Status(200)
}
