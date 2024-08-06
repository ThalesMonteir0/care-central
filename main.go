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

	userController, patientController, sessionController, movementController := initDependencies(db)

	app := gin.Default()

	controller.InitRoutes(
		app,
		userController,
		patientController,
		sessionController,
		movementController,
	)

	if err := app.Run(":5000"); err != nil {
		log.Fatal("Error loading application")
	}
}

func initDependencies(db *gorm.DB) (
	controller.UserControllerInterface,
	controller.PatientControllerInterface,
	controller.SessionControllerInterface,
	controller.MovementControllerInterface) {

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	patientRepository := repository.NewPatientRepository(db)
	patientService := service.NewPatientService(patientRepository)
	patientController := controller.NewPatientController(patientService)

	sessionRepository := repository.NewSessionsRepository(db)
	sessionService := service.NewSessionService(sessionRepository)
	sessionController := controller.NewSessionController(sessionService)

	movementRepository := repository.NewMovementRepository(db)
	movementService := service.NewMovementService(movementRepository)
	movementController := controller.NewMovementController(movementService)

	return userController, patientController, sessionController, movementController
}
