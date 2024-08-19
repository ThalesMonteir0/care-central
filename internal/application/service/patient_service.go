package service

import (
	"github.com/ThalesMonteir0/care-central/internal/application/port/input"
	"github.com/ThalesMonteir0/care-central/internal/application/port/output"
)

type patientService struct {
	repository output.PatientRepositoryInterface
}

func NewPatientService(repository output.PatientRepositoryInterface) input.PatientServiceInterface {
	return &patientService{
		repository: repository,
	}
}
