package user

import (
	"net/http"
	"net/mail"
	"ping-api/auth"
	"ping-api/database"
	"ping-api/inputs"
	"ping-api/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	user := models.User{}
	userInputs := inputs.Login{}

	c.BindJSON(&userInputs)

	if userInputs.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Email is required",
		})
		return
	}

	if userInputs.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Password is required",
		})
		return
	}

	_, err := mail.ParseAddress(userInputs.Email)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid email address",
		})
		return
	}

	errData := database.DB.Where(&inputs.Login{Email: userInputs.Email}).First(&user)

	if errData.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid email or password",
		})
		return
	}

	passStatus := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInputs.Password))

	if passStatus != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid email or password",
		})
		return
	}

	token, errToken := auth.JwtGen(user.ID)

	if errToken != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid generate token please try again later",
			"error":   errToken.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome " + user.Name,
		"token":   token,
	})
}
