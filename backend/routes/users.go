package routes

import (
	"net/http"

	"example.com/anon-project/models"
	"github.com/gin-gonic/gin"
)

func getUsers(c *gin.Context) {
	users, err := models.GetAllUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Colud not fetch data."})
		return
	}

	c.JSON(http.StatusOK, users)
}
