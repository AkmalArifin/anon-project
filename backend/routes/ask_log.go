package routes

import (
	"net/http"
	"strconv"

	"example.com/anon-project/models"
	"github.com/gin-gonic/gin"
)

func getAskLogByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch data"})
		return
	}

	ask, err := models.GetAskByID(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch data"})
		return
	}

	userID := c.GetInt64("user_id")

	if ask.UserID != userID {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorize"})
		return
	}

	c.JSON(http.StatusOK, ask)
}

func getAskLogByUser(c *gin.Context) {
	userID := c.GetInt64("user_id")

	asks, err := models.GetAskByUserID(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch data", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, asks)
}

func createAskLog(c *gin.Context) {
	var question models.Question
	err := c.ShouldBindJSON(&question)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}

	user, err := models.GetUserByUsername(question.Username.String)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not find user"})
		return
	}

	var ask models.Ask
	ask.UserID = user.ID
	ask.Question = question.Question
	err = ask.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not store data", "err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ask)
}
