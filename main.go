package main

import (
	"github.com/ThalesMonteir0/care-central/adapter/input/controller"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := gin.Default()

	controller.InitRoutes(app)

	if err := app.Run(":5000"); err != nil {
		log.Fatal("Error loading application")
	}
}
