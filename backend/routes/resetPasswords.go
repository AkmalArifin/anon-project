package routes

import (
	"net/http"
	"time"

	"example.com/anon-project/models"
	"example.com/anon-project/utils"
	"github.com/gin-gonic/gin"
)

func createResetPassword(c *gin.Context) {
	var request models.ResetPasswordRequest
	err := c.ShouldBindJSON(&request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not bind JSON"})
		return
	}

	user, err := models.GetUserByEmail(request.Email.ValueOrZero())

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not find user"})
		return
	}

	var resetPassword models.ResetPassword
	resetPassword.UserID.SetValid(user.ID)
	err = resetPassword.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not store data"})
		return
	}

	// send email reset password

	c.JSON(http.StatusOK, gin.H{"message": "Data has been saved", "data": resetPassword})
}

func verifyResetPassword(c *gin.Context) {
	var resetPassword models.ResetPassword
	err := c.ShouldBindJSON(&resetPassword)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Unauthorized"})
		return
	}

	isValid, err := utils.VerifyResetPasswordToken(resetPassword.Token.ValueOrZero())

	if err != nil || !isValid {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Unauthorized"})
		return
	}

	expectedResetPassword, err := models.GetResetPasswordByID(resetPassword.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Unauthorized"})
		return
	}

	isValid = expectedResetPassword.Token.ValueOrZero() == resetPassword.Token.ValueOrZero()

	if !isValid {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Unauthorized"})
		return
	}

	isValid = time.Now().Before(expectedResetPassword.ExpiresAt.Time)

	if !isValid {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Unauthorized"})
		return
	}

	user, err := models.GetUserByID(expectedResetPassword.UserID.ValueOrZero())

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Unauthorized"})
		return
	}

	token, err := utils.GenerateJWTToken(user.ID, user.Username.ValueOrZero())

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Unauthorized"})
		return
	}

	c.JSON(http.StatusOK, token)
}
