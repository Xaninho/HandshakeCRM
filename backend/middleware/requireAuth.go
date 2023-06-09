package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/handshakeCRM/initializers"
	"github.com/handshakeCRM/models"
)

func RequireAuth(c *gin.Context) {
	// Get the token from the Authorization header
	tokenString := c.GetHeader("Authorization")

	// Check if the token is present
	if tokenString == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized1",
		})
		return
	}

	// Decode/Validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	// log these
	fmt.Println("tokenString: ", tokenString)
	fmt.Println("token: ", token.Valid)
	fmt.Println("err: ", err)
	if err != nil || !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized2",
		})
		return
	}

	// Access the claims from the token
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// Check the exp
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized3",
			})
			return
		}

		// Find the user with token sub
		var user models.User
		initializers.DB.First(&user, claims["userId"])

		if user.ID == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized4",
			})
			return
		}

		// Attach the user to the request context
		c.Set("user", user)
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized5",
		})
	}
}
