package user

import (
	"net/http"
	"ping-api/database"
	"ping-api/models"

	"github.com/gin-gonic/gin"
)

func Me(c *gin.Context) {
	id, exists := c.Get("id")

	if exists == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "User id not found",
		})
		return
	}

	user := models.User{}

	userID, ok := id.(string)

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid user id type",
		})
		return
	}

	err := database.DB.Where(models.User{ID: userID}).First(&user).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid or expired token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": gin.H{
			"id":        user.ID,
			"name":      user.Name,
			"email":     user.Email,
			"createdAt": user.CreatedAt,
			"updatedAt": user.UpdatedAt,
		},
	})
}
