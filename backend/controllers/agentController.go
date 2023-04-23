package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/handshakeCRM/models"
	"github.com/handshakeCRM/initializers"
)

type AgentRequest struct {
	Name        	string `json:"name"`
	Email       	string `json:"email"`
	Position  		models.EnumType `json:"position"`
	PhoneNumber 	int `json:"phoneNumber"`
	PhotoURL    	string `json:"photoURL"`
}

func AgentCreate (c *gin.Context) {

	// Get data off req body
	var body AgentRequest
	c.Bind(&body)

	// Create a post
	agent := models.Agent{	
		Name: body.Name,
		Email: body.Email,
		Position: body.Position,
		PhoneNumber: body.PhoneNumber,
		PhotoURL: body.PhotoURL,
	}
	
	result := initializers.DB.Create(&agent)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return it
	c.JSON(200, gin.H{
		"agent": agent,
	})

}

func AgentIndex (c *gin.Context) {

	//Get the posts
	var agents []models.Agent
	initializers.DB.Find(&agents)

	//Respond with them
	c.JSON(200, gin.H{
		"agents": agents,
	})
}

func AgentShow (c *gin.Context) {

	//get id off url
	id := c.Param("id")

	//Get the posts
	var agent []models.Agent
	initializers.DB.First(&agent, id)

	//Respond with them
	c.JSON(200, gin.H{
		"agent": agent,
	})
}

func AgentUpdate (c *gin.Context) {

	// Get the id off the url
	id := c.Param("id")

	// Get the data off req body
	var body AgentRequest
	c.Bind(&body)

	// Find the post were updating
	var agent models.Agent
	initializers.DB.First(&agent, id)

	// Update it
	initializers.DB.Model(&agent).Updates(models.Agent{
		Name: body.Name,
		Email: body.Email,
		Position: body.Position,
		PhoneNumber: body.PhoneNumber,
		PhotoURL: body.PhotoURL,
	})

	// Respond with it
	c.JSON(200, gin.H{
		"agent": agent,
	})

}

func AgentDelete (c * gin.Context) {

	// Get the id off the url
	id := c.Param("id")	
	
	// Delete the company
	initializers.DB.Delete(&models.Agent{}, id)

	// Respond
	c.Status(200)
}