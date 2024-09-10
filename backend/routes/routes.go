package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	r.GET("/ping", helloWorld)

	// User
	r.POST("/login", login)
	r.POST("/signup", signup)

	// User
	// r.GET("/users", getUsers)
	// r.GET("/users/:id", getUser)
	// r.POST("/users", createUser)
	// r.PUT("/users/:id", updateUser)
	// r.DELETE("/users/:id", deleteUser)

}

func helloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hellow World!"})
}
