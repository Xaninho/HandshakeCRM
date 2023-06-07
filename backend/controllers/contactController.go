package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/handshakeCRM/initializers"
	"github.com/handshakeCRM/models"
	"gorm.io/gorm"
)

type ContactRequest struct {
	ID            uint   `json:"id"`
	CompanyId     uint   `json:"companyId"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	PhoneNumber   string `json:"phoneNumber"`
	Aniversary    string `json:"aniversary"`
	DispositionId int    `json:"dispositionId"`
}

func ContactCreate(c *gin.Context) {

	var body ContactRequest
	if err := c.Bind(&body); err != nil {
		fmt.Println("Error parsing request body:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	contact := models.Contact{
		CompanyId:     body.CompanyId,
		Name:          body.Name,
		Email:         body.Email,
		PhoneNumber:   body.PhoneNumber,
		Aniversary:    body.Aniversary,
		DispositionId: body.DispositionId,
	}

	result := initializers.DB.Create(&contact)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create contact",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"contact": contact,
	})

}

func ContactIndex(c *gin.Context) {

	var contacts []models.Contact

	initializers.DB.Preload("Company").Preload("Company.Currency").Preload("Company.Country").Preload("Disposition", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, type, description")
	}).Preload("Disposition").Find(&contacts)

	c.JSON(200, gin.H{
		"contacts": contacts,
	})
}

func ContactShow(c *gin.Context) {

	id := c.Param("id")

	var contact []models.Contact
	initializers.DB.First(&contact, id)

	c.JSON(200, gin.H{
		"contact": contact,
	})
}

func ContactUpdate(c *gin.Context) {

	id := c.Param("id")

	var body ContactRequest
	c.Bind(&body)

	var contact models.Contact
	initializers.DB.First(&contact, id)

	initializers.DB.Model(&contact).Updates(models.Contact{
		CompanyId:     body.CompanyId,
		Name:          body.Name,
		Email:         body.Email,
		PhoneNumber:   body.PhoneNumber,
		Aniversary:    body.Aniversary,
		DispositionId: body.DispositionId,
	})

	c.JSON(200, gin.H{
		"contact": contact,
	})
}

func ContactDelete(c *gin.Context) {

	// Get the id off the url
	id := c.Param("id")

	// Delete the company
	initializers.DB.Delete(&models.Contact{}, id)

	// Respond
	c.Status(http.StatusNoContent)
}
