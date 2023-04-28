package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/handshakeCRM/initializers"
	"github.com/handshakeCRM/models"
)

type EnumTypeRequest struct {
	Type        string `json:"type"`
	Description string `json:"description"`
	Active      bool   `json:"active"`
}

func EnumTypeCreate(c *gin.Context) {

	// Get data off req body
	var body EnumTypeRequest

	c.Bind(&body)

	// Create a post
	enumType := models.EnumType{
		Type:        body.Type,
		Description: body.Description,
		Active:      body.Active,
	}

	result := initializers.DB.Create(&enumType)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return it
	c.JSON(200, gin.H{
		"enumType": enumType,
	})

}

func EnumTypeIndex(c *gin.Context) {

	//Get the typeOfEnum from the HTTP request
	typeOfEnum := c.Query("typeOfEnum")

	//Get the enumTypes that match the typeOfEnum
	var enumTypes []models.EnumType
	initializers.DB.Where("type = ?", typeOfEnum).Find(&enumTypes)

	//Respond with them
	c.JSON(200, gin.H{
		"enumTypes": enumTypes,
	})
}

func EnumTypeShow(c *gin.Context) {

	//get id off url
	id := c.Param("id")

	//Get the posts
	var enumType []models.EnumType
	initializers.DB.First(&enumType, id)

	//Respond with them
	c.JSON(200, gin.H{
		"enumType": enumType,
	})
}

func EnumTypeUpdate(c *gin.Context) {

	// Get the id off the url
	id := c.Param("id")

	// Get the data off req body
	var body EnumTypeRequest

	c.Bind(&body)

	// Find the post were updating
	var enumType models.EnumType
	initializers.DB.First(&enumType, id)

	// Update it
	initializers.DB.Model(&enumType).Updates(models.EnumType{
		Type:        body.Type,
		Description: body.Description,
		Active:      body.Active,
	})

	// Respond with it
	c.JSON(200, gin.H{
		"enuymType": enumType,
	})

}

func EnumTypeDelete(c *gin.Context) {

	// Get the id off the url
	id := c.Param("id")

	// Delete the company
	initializers.DB.Delete(&models.EnumType{}, id)

	// Respond
	c.Status(200)
}
