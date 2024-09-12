package main

import (
	"os"

	"example.com/anon-project/db"
	"example.com/anon-project/middlewares"
	"example.com/anon-project/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic("could not load .env")
	}

	db.InitDB()

	r := gin.Default()

	corsConfig := middlewares.SetCORS()
	r.Use(cors.New(corsConfig))

	routes.RegisterRoutes(r)

	port := ":" + os.Getenv("PORT")
	r.Run(port)
}
