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

	var body EnumTypeRequest

	c.Bind(&body)

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

	c.JSON(200, gin.H{
		"enumType": enumType,
	})

}

func EnumTypeIndex(c *gin.Context) {

	typeOfEnum := c.Query("typeOfEnum")

	var enumTypes []models.EnumType
	initializers.DB.Where("type = ?", typeOfEnum).Find(&enumTypes)

	c.JSON(200, gin.H{
		"enumTypes": enumTypes,
	})
}

func EnumTypeShow(c *gin.Context) {

	id := c.Param("id")

	var enumType []models.EnumType
	initializers.DB.First(&enumType, id)

	c.JSON(200, gin.H{
		"enumType": enumType,
	})
}

func EnumTypeUpdate(c *gin.Context) {

	id := c.Param("id")

	var body EnumTypeRequest
	c.Bind(&body)

	var enumType models.EnumType
	initializers.DB.First(&enumType, id)

	initializers.DB.Model(&enumType).Updates(models.EnumType{
		Type:        body.Type,
		Description: body.Description,
		Active:      body.Active,
	})

	c.JSON(200, gin.H{
		"enuymType": enumType,
	})

}

func EnumTypeDelete(c *gin.Context) {

	id := c.Param("id")

	initializers.DB.Delete(&models.EnumType{}, id)

	c.Status(200)
}
