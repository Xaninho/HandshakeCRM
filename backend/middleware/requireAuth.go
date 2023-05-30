package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/handshakeCRM/go/pkg/mod/github.com/gin-gonic/gin@v1.8.2"
	"github.com/handshakeCRM/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible"
)

func RequireAuth(c *gin.Context) {
	// Get the cookie off req
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized)
	}

	// Decode/Validate It

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.GetEnv("SECRET_KEY")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		// Check the exp
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatusJSON(http.StatusUnauthorized)
		}

		// Find the user with token sub
		var user models.Player

		initializers.DB.First(&user, claims["userId"])

		if user.ID == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized)
		}

		// Attach to req
		c.Set("user", user)

		// Continue
		c.Next()

	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized)
	}

}
