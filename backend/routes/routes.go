package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	r.GET("/ping", helloWorld)

	// User
	r.GET("/users", getUsers)

}

func helloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hellow World!"})
}
