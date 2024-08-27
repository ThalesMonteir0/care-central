package input

import (
	"github.com/ThalesMonteir0/care-central/internal/adapter/input/model/request"
	"github.com/ThalesMonteir0/care-central/internal/application/domain"
	"github.com/ThalesMonteir0/care-central/pkg/rest_err"
)

type PatientServiceInterface interface {
	CreatePatient(domain.PatientDomain) *rest_err.RestErr
	UpdatePatient(domain.PatientDomain) *rest_err.RestErr
	DeletePatient(domain.PatientDomain) *rest_err.RestErr
	FindPatient(domain.PatientDomain, request.PatientFilters) ([]domain.PatientDomain, *rest_err.RestErr)
}
