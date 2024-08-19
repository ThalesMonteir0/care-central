package input

import (
	"github.com/ThalesMonteir0/care-central/internal/application/domain"
	"github.com/ThalesMonteir0/care-central/pkg/rest_err"
)

type MovementServiceInterface interface {
	FindAllMovements(domain.MovementDomain) ([]domain.MovementDomain, *rest_err.RestErr)
	CreateMovement(domain.MovementDomain) *rest_err.RestErr
	DeleteMovement(domain.MovementDomain) *rest_err.RestErr
}
