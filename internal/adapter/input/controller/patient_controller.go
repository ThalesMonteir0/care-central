package controller

import (
	"github.com/ThalesMonteir0/care-central/internal/application/port/input"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type patientController struct {
	service input.PatientServiceInterface
	logger  *zap.Logger
}

type PatientControllerInterface interface {
	getPatient(ctx *gin.Context)
	CreatePatient(ctx *gin.Context)
	DeletePatient(ctx *gin.Context)
	UpdatePatient(ctx *gin.Context)
}

func NewPatientController(service input.PatientServiceInterface, logger *zap.Logger) PatientControllerInterface {
	return &patientController{
		service: service,
		logger:  logger,
	}
}
