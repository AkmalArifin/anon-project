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

	// Ask
	auth.GET("/ask-log", getAskLogByUser)
	auth.GET("/ask-log/:id", getAskLogByID)
	r.POST("/ask-log", createAskLog)

	// User
	r.GET("/users", getUsers)
	r.GET("/users/:id", getUser)
	r.POST("/users", createUser)
	r.PUT("/users/:id", updateUser)
	r.DELETE("/users/:id", deleteUser)

}
