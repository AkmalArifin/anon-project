package routes

import (
	"net/http"
	"strconv"

	"example.com/anon-project/models"
	"github.com/gin-gonic/gin"
)

func getLogByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch data"})
		return
	}

	log, err := models.GetLogByID(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch data"})
		return
	}

	userID := c.GetInt64("user_id")

	if log.UserID != userID {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorize"})
		return
	}

	c.JSON(http.StatusOK, log)
}

func getLogByUser(c *gin.Context) {
	userID := c.GetInt64("user_id")

	logs, err := models.GetLogByUserID(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch data", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, logs)
}

func createLog(c *gin.Context) {
	var message models.Message
	err := c.ShouldBindJSON(&message)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}

	user, err := models.GetUserByUsername(message.Username.String)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not find user"})
		return
	}

	var log models.Log
	log.UserID = user.ID
	log.Message = message.Message
	err = log.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not store data", "err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, log)
}
