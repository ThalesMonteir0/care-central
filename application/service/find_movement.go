package service

import (
	"github.com/ThalesMonteir0/care-central/application/domain"
	"github.com/ThalesMonteir0/care-central/configuration/rest_err"
)

func (m *movementService) FindAllMovements(domain domain.MovementDomain) ([]domain.MovementDomain, *rest_err.RestErr) {
	movements, err := m.repository.FindAll(domain)
	if err != nil {
		return nil, err
	}

	return movements, nil
}
