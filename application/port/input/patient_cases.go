package input

import (
	"github.com/ThalesMonteir0/care-central/application/domain"
	"github.com/ThalesMonteir0/care-central/configuration/rest_err"
)

type PatientServiceInterface interface {
	CreatePatient(domain.PatientDomain) *rest_err.RestErr
	UpdatePatient(domain.PatientDomain) *rest_err.RestErr
	DeletePatient(domain.PatientDomain) *rest_err.RestErr
	FindPatient(domain.PatientDomain) ([]domain.PatientDomain, *rest_err.RestErr)
}
