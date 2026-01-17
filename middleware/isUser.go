package middleware

import (
	"net/http"
	"ping-api/auth"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func IsUser(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "You are not authorized",
		})
		return
	}

	tokenData, errToken := auth.JwtVer(token)

	if errToken != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid or expired token",
		})
		return
	}

	if tokenData == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid or expired token",
		})
		return
	}

	claims, ok := tokenData.Claims.(jwt.MapClaims)

	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid or expired token",
		})
		return
	}

	userID, okUser := claims["id"]

	if !okUser {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid or expired token",
		})
		return
	}

	c.Set("id", userID)
	c.Next()
}
