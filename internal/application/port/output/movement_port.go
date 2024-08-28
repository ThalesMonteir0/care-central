package output

import (
	"github.com/ThalesMonteir0/care-central/internal/adapter/input/model/request"
	"github.com/ThalesMonteir0/care-central/internal/application/domain"
	"github.com/ThalesMonteir0/care-central/pkg/rest_err"
)

type MovementRepositoryInterface interface {
	FindAll(domain.MovementDomain, request.MovementsFilters) ([]domain.MovementDomain, *rest_err.RestErr)
	Create(domain.MovementDomain) *rest_err.RestErr
	Delete(domain.MovementDomain) *rest_err.RestErr
}
