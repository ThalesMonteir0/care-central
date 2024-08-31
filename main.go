package main

import (
	"github.com/ThalesMonteir0/care-central/internal/adapter/input/controller"
	"github.com/ThalesMonteir0/care-central/internal/adapter/output/http_client"
	"github.com/ThalesMonteir0/care-central/internal/adapter/output/repository"
	"github.com/ThalesMonteir0/care-central/internal/application/service"
	"github.com/ThalesMonteir0/care-central/pkg/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := database.NewPostgresql()

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	userController, patientController, sessionController, movementController, createPixController := initDependencies(db, logger)

	app := gin.Default()

	controller.InitRoutes(
		app,
		userController,
		patientController,
		sessionController,
		movementController,
		createPixController,
	)

	if err := app.Run(":5000"); err != nil {
		log.Fatal("Error loading application")
	}
}

func initDependencies(db *gorm.DB, logger *zap.Logger) (
	controller.UserControllerInterface,
	controller.PatientControllerInterface,
	controller.SessionControllerInterface,
	controller.MovementControllerInterface,
	controller.CreatePixController,
) {

	userRepository := repository.NewUserRepository(db, logger)
	userService := service.NewUserService(userRepository, logger)
	userController := controller.NewUserController(userService, logger)

	patientRepository := repository.NewPatientRepository(db, logger)
	patientService := service.NewPatientService(patientRepository, logger)
	patientController := controller.NewPatientController(patientService, logger)

	sessionRepository := repository.NewSessionsRepository(db, logger)
	sessionService := service.NewSessionService(sessionRepository, logger)
	sessionController := controller.NewSessionController(sessionService, logger)

	movementRepository := repository.NewMovementRepository(db, logger)
	movementService := service.NewMovementService(movementRepository, logger)
	movementController := controller.NewMovementController(movementService, logger)

	pixGeradosRepository := repository.NewPixGeradosRepository(db, logger)
	pixGeradosService := service.NewPixGeradosService(pixGeradosRepository, logger)

	httpClientPix := http_client.NewHttpClientPix(logger)
	createPixService := service.NewCreatePixService(httpClientPix, pixGeradosService, logger)
	createPixController := controller.NewCreatePixController(createPixService, logger)

	return userController, patientController, sessionController, movementController, createPixController
}
