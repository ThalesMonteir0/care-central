package service

import (
	"github.com/ThalesMonteir0/care-central/application/domain"
	"github.com/ThalesMonteir0/care-central/configuration/rest_err"
)

func (m *movementService) CreateMovement(domain domain.MovementDomain) *rest_err.RestErr {
	if domain.Value <= 0 {
		return rest_err.NewBadRequestError("value cannot be negative or null")
	}

	if err := m.repository.Create(domain); err != nil {
		return err
	}

	return nil
}
