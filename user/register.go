package user

import (
	"net/http"
	"net/mail"
	"ping-api/database"
	"ping-api/models"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	user := models.User{}

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error binding JSON",
		})
		return
	}

	if user.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Name is required",
		})
		return
	}

	if user.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Email is required",
		})
		return
	}

	if user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Password is required",
		})
		return
	}

	_, err := mail.ParseAddress(user.Email)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid email address",
		})
		return
	}

	if len(user.Password) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Password must be at least 6 characters long",
		})
		return
	}

	if err := database.DB.Where(&models.User{Email: user.Email}).First(&user); err.Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Email already in use",
		})
		return
	}

	errDB := database.DB.Create(&user).Error

	if errDB != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": errDB.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": user,
	})

}
