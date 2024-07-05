package main

import (
	"github.com/ThalesMonteir0/care-central/adapter/input/controller"
	"github.com/ThalesMonteir0/care-central/adapter/output/repository"
	"github.com/ThalesMonteir0/care-central/application/service"
	"github.com/ThalesMonteir0/care-central/configuration/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	userController := initDependencies()

	app := gin.Default()

	controller.InitRoutes(app, userController)

	if err := app.Run(":5000"); err != nil {
		log.Fatal("Error loading application")
	}
}

func initDependencies() controller.UserControllerInterface {
	db := database.NewPostgresql()
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	return userController
}
