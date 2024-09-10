package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic("could not load .env")
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hellow World!"})
	})

	port := ":" + os.Getenv("PORT")
	r.Run(port)
}
