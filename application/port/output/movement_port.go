package output

import (
	"github.com/ThalesMonteir0/care-central/application/domain"
	"github.com/ThalesMonteir0/care-central/configuration/rest_err"
)

type MovementRepositoryInterface interface {
	FindAll(domain.MovementDomain) ([]domain.MovementDomain, *rest_err.RestErr)
	Create(domain.MovementDomain) *rest_err.RestErr
	Delete(domain.MovementDomain) *rest_err.RestErr
}
