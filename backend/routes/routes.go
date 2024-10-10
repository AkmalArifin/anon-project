package routes

import (
	"example.com/anon-project/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	auth := r.Group("/")
	auth.Use(middlewares.Authenticate)

	// User
	r.POST("/login", login)
	r.POST("/signup", signup)

	// Log
	auth.GET("/log", getLogByUser)
	auth.GET("/log/:id", getLogByID)
	auth.PUT("/log/star/:id", starLog)
	auth.DELETE("/log/delete/:id", deleteLog)
	r.POST("/log", createLog)

	// User
	r.GET("/users", getUsers)
	r.GET("/users/:id", getUser)
	r.GET("/users/username/:username", getUserUsername)
	r.POST("/users", createUser)
	r.PUT("/users/:id", updateUser)
	r.DELETE("/users/:id", deleteUser)

}
