package api

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/SaidovZohid/auth-signin-signup-middleware/api/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func (sr *RouteOptions) ReQuireAuth(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["expired_at"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		user, err := sr.Storage.User().Get(claims["email"].(string))
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ResponseError{
				Error: err.Error(),
			})
		}

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Set("user", user)

		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
