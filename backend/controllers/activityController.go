package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/handshakeCRM/initializers"
	"github.com/handshakeCRM/models"
)

type ActivityRequest struct {
	ContactID uint
	UserID    uint
	Date      string
	Title     string
	Location  string
	Duration  string
	Notes     string
	OutcomeID int
}

func ActivityCreate(c *gin.Context) {
	var body ActivityRequest
	c.Bind(&body)

	activity := models.Activity{
		ContactID: body.ContactID,
		UserID:    body.UserID,
		Date:      body.Date,
		Title:     body.Title,
		Location:  body.Location,
		Duration:  body.Duration,
		Notes:     body.Notes,
		OutcomeId: body.OutcomeID,
	}

	result := initializers.DB.Create(&activity)
	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"activity": activity,
	})
}

func ActivityIndex(c *gin.Context) {
	var activities []models.Activity
	initializers.DB.Preload("Contact").Preload("User").Preload("Outcome").Find(&activities)

	c.JSON(200, gin.H{
		"activities": activities,
	})
}

func ActivityShow(c *gin.Context) {
	id := c.Param("id")

	var activity models.Activity
	result := initializers.DB.Preload("Contact").Preload("User").Preload("Outcome").First(&activity, id)
	if result.Error != nil {
		c.Status(404)
		return
	}

	c.JSON(200, gin.H{
		"activity": activity,
	})
}

func ActivityUpdate(c *gin.Context) {
	id := c.Param("id")

	var body ActivityRequest
	c.Bind(&body)

	var activity models.Activity
	result := initializers.DB.First(&activity, id)
	if result.Error != nil {
		c.Status(404)
		return
	}

	initializers.DB.Model(&activity).Updates(models.Activity{
		ContactID: body.ContactID,
		UserID:    body.UserID,
		Date:      body.Date,
		Title:     body.Title,
		Location:  body.Location,
		Duration:  body.Duration,
		Notes:     body.Notes,
		OutcomeId: body.OutcomeID,
	})

	c.JSON(200, gin.H{
		"activity": activity,
	})
}

func ActivityDelete(c *gin.Context) {
	id := c.Param("id")

	result := initializers.DB.Delete(&models.Activity{}, id)
	if result.RowsAffected == 0 {
		c.Status(404)
		return
	}

	c.Status(200)
}
