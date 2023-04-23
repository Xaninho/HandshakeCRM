package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/handshakeCRM/models"
	"github.com/handshakeCRM/initializers"
)

type ClientRequest struct {
	Company  		models.Company `json:"company"`
	Name       		string `json:"name"`
	Email      		string `json:"email"`
	PhotoURL   		string `json:"photoURL"`
	Function 		models.EnumType `json:"function"`
	Type     		models.EnumType `json:"type"`
	PhoneNumber 	int `json:"phoneNumber"`
	Aniversary 		string `json:"aniversary"`
}

func ClientCreate (c *gin.Context) {

	// Get data off req body
	var body ClientRequest
	c.Bind(&body)

	// Create a post
	client := models.Client{	
		Company: body.Company,
		Name: body.Name,
		Email: body.Email,
		PhotoURL: body.PhotoURL,
		Function: body.Function,
		Type: body.Type,
		PhoneNumber: body.PhoneNumber,
		Aniversary: body.Aniversary,
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

func ClientIndex (c *gin.Context) {
	//Get the posts
	var clients []models.Client
	initializers.DB.Find(&clients)

	//Respond with them
	c.JSON(200, gin.H{
		"clients": clients,
	})
}

func ClientShow (c *gin.Context) {
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

func ClientUpdate (c *gin.Context) {
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
		Company: body.Company,
		Name: body.Name,
		Email: body.Email,
		PhotoURL: body.PhotoURL,
		Function: body.Function,
		Type: body.Type,
		PhoneNumber: body.PhoneNumber,
		Aniversary: body.Aniversary,
	})

	// Respond with it
	c.JSON(200, gin.H{
		"client": client,
	})

}

func ClientDelete (c * gin.Context) {

	// Get the id off the url
	id := c.Param("id")	
	
	// Delete the company
	initializers.DB.Delete(&models.Client{}, id)

	// Respond
	c.Status(200)
}