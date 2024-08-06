package input

import (
	"github.com/ThalesMonteir0/care-central/application/domain"
	"github.com/ThalesMonteir0/care-central/configuration/rest_err"
)

type MovementServiceInterface interface {
	FindAllMovements(domain.MovementDomain) ([]domain.MovementDomain, *rest_err.RestErr)
	CreateMovement(domain.MovementDomain) *rest_err.RestErr
	DeleteMovement(domain.MovementDomain) *rest_err.RestErr
}
