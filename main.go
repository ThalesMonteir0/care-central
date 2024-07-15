package main

import (
	"github.com/ThalesMonteir0/care-central/adapter/input/controller"
	"github.com/ThalesMonteir0/care-central/adapter/output/repository"
	"github.com/ThalesMonteir0/care-central/application/service"
	"github.com/ThalesMonteir0/care-central/configuration/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := database.NewPostgresql()

	userController, patientController := initDependencies(db)

	app := gin.Default()

	controller.InitRoutes(app, userController, patientController)

	if err := app.Run(":5000"); err != nil {
		log.Fatal("Error loading application")
	}
}

func initDependencies(db *gorm.DB) (controller.UserControllerInterface, controller.PatientControllerInterface) {

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	patientRepository := repository.NewPatientRepository(db)
	patientService := service.NewPatientService(patientRepository)
	patientController := controller.NewPatientController(patientService)

	return userController, patientController
}
