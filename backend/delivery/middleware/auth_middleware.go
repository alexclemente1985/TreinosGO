package middleware

import (
	"bank-app/config/auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := strings.Split(c.GetHeader("Authorization"), " ")

		if len(tokenStr) < 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token invalid"})

			return
		}

		token, err := jwt.Parse(tokenStr[1], func(t *jwt.Token) (interface{}, error) {
			return auth.JwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token invalid or not provided"})

			return
		}

		claims := token.Claims.(jwt.MapClaims)
		c.Set("userID", claims["sub"])

		c.Next()
	}
}
