package service

import (
	"github.com/ThalesMonteir0/care-central/internal/application/port/input"
	"github.com/ThalesMonteir0/care-central/internal/application/port/output"
	"go.uber.org/zap"
)

type patientService struct {
	repository output.PatientRepositoryInterface
	logger     *zap.Logger
}

func NewPatientService(repository output.PatientRepositoryInterface, logger *zap.Logger) input.PatientServiceInterface {
	return &patientService{
		repository: repository,
		logger:     logger,
	}
}
