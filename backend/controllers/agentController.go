package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/handshakeCRM/initializers"
	"github.com/handshakeCRM/models"
	"golang.org/x/crypto/bcrypt"
)

type AgentRequest struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PositionId  int    `json:"positionId"`
	PhoneNumber int    `json:"phoneNumber"`
	PhotoURL    string `json:"photoURL"`
}

func Signup(c *gin.Context) {

	// Get the email/pass off req body
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	// Create the user
	user := models.Agent{Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create user",
		})
		return
	}
}

func Login(c *gin.Context) {
	// Get the email/pass off req body
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Look up requested user
	var user models.Agent
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid Email or Password",
		})
		return
	}

	// Compare sent in pass with saved user pass
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid Email or Password",
		})
		return
	}

	// Create a JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID,
		"exp":    time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign in ang get the complete endoced token as a string using the secret key
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create token",
		})
		return
	}

	// send it back
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 60*60*24*30, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}

func AgentCreate(c *gin.Context) {

	// Get data off req body
	var body AgentRequest
	c.Bind(&body)

	// Create a post
	agent := models.Agent{
		Name:        body.Name,
		Email:       body.Email,
		PositionId:  body.PositionId,
		PhoneNumber: body.PhoneNumber,
		PhotoURL:    body.PhotoURL,
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

func AgentIndex(c *gin.Context) {

	//Get the posts
	var agents []models.Agent

	initializers.DB.Preload("Position").Find(&agents)

	//Respond with them
	c.JSON(200, gin.H{
		"agents": agents,
	})
}

func AgentShow(c *gin.Context) {

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

func AgentUpdate(c *gin.Context) {

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
		Name:        body.Name,
		Email:       body.Email,
		PositionId:  body.PositionId,
		PhoneNumber: body.PhoneNumber,
		PhotoURL:    body.PhotoURL,
	})

	// Respond with it
	c.JSON(200, gin.H{
		"agent": agent,
	})

}

func AgentDelete(c *gin.Context) {

	// Get the id off the url
	id := c.Param("id")

	// Delete the company
	initializers.DB.Delete(&models.Agent{}, id)

	// Respond
	c.Status(200)
}
