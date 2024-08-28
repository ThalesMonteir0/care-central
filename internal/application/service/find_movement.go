package service

import (
	"github.com/ThalesMonteir0/care-central/internal/adapter/input/model/request"
	"github.com/ThalesMonteir0/care-central/internal/application/domain"
	"github.com/ThalesMonteir0/care-central/pkg/rest_err"
)

func (m *movementService) FindAllMovements(domain domain.MovementDomain, filters request.MovementsFilters) ([]domain.MovementDomain, *rest_err.RestErr) {
	movements, err := m.repository.FindAll(domain, filters)
	if err != nil {
		return nil, err
	}

	return movements, nil
}
