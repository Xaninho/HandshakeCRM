package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/handshakeCRM/initializers"
	"github.com/handshakeCRM/models"
	"gorm.io/gorm"
)

type ClientRequest struct {
	CompanyId       int    `json:"companyId"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	PhotoURL        string `json:"photoURL"`
	FunctionId      int    `json:"functionId"`
	TypeId          int    `json:"typeId"`
	PhoneNumber     int    `json:"phoneNumber"`
	Aniversary      string `json:"aniversary"`
	AssignedAgentId int    `json:"assignedAgentId"`
}

func ClientCreate(c *gin.Context) {

	// Get data off req body
	var body ClientRequest
	c.Bind(&body)

	// Create a post
	client := models.Client{
		CompanyId:       body.CompanyId,
		Name:            body.Name,
		Email:           body.Email,
		PhotoURL:        body.PhotoURL,
		FunctionId:      body.FunctionId,
		TypeId:          body.TypeId,
		PhoneNumber:     body.PhoneNumber,
		Aniversary:      body.Aniversary,
		AssignedAgentId: body.AssignedAgentId,
	}

	result := initializers.DB.Create(&client)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return it
	c.JSON(200, gin.H{
		"client": client,
	})
}

func ClientIndex(c *gin.Context) {

	//Get the posts
	var clients []models.Client

	initializers.DB.Preload("AssignedAgent", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, name, photo_url")
	}).Preload("Company").Preload("Company.Currency").Preload("Company.Country").Preload("Type", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, type, description")
	}).Find(&clients)

	//Respond with them
	c.JSON(200, gin.H{
		"clients": clients,
	})
}

func ClientShow(c *gin.Context) {
	//get id off url
	id := c.Param("id")

	//Get the posts
	var client []models.Client
	initializers.DB.First(&client, id)

	//Respond with them
	c.JSON(200, gin.H{
		"client": client,
	})
}

func ClientUpdate(c *gin.Context) {
	// Get the id off the url
	id := c.Param("id")

	// Get the data off req body
	var body ClientRequest
	c.Bind(&body)

	// Find the post were updating

	var client models.Client
	initializers.DB.First(&client, id)

	// Update it
	initializers.DB.Model(&client).Updates(models.Client{
		CompanyId:       body.CompanyId,
		Name:            body.Name,
		Email:           body.Email,
		PhotoURL:        body.PhotoURL,
		FunctionId:      body.FunctionId,
		TypeId:          body.TypeId,
		PhoneNumber:     body.PhoneNumber,
		Aniversary:      body.Aniversary,
		AssignedAgentId: body.AssignedAgentId,
	})

	// Respond with it
	c.JSON(200, gin.H{
		"client": client,
	})
}

func ClientDelete(c *gin.Context) {

	// Get the id off the url
	id := c.Param("id")

	// Delete the company
	initializers.DB.Delete(&models.Client{}, id)

	// Respond
	c.Status(200)
}
