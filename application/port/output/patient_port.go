package output

import (
	"github.com/ThalesMonteir0/care-central/application/domain"
	"github.com/ThalesMonteir0/care-central/configuration/rest_err"
)

type PatientRepositoryInterface interface {
	CreatePatient(domain.PatientDomain) *rest_err.RestErr
	FindPatients(domain.PatientDomain) ([]domain.PatientDomain, *rest_err.RestErr)
	DeletePatient(domain.PatientDomain) *rest_err.RestErr
	UpdatePatient(domain.PatientDomain) *rest_err.RestErr
	FindPatientByID(domain.PatientDomain) (*domain.PatientDomain, *rest_err.RestErr)
}
