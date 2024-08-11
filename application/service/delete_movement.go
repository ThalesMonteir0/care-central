package service

import (
	"github.com/ThalesMonteir0/care-central/application/domain"
	"github.com/ThalesMonteir0/care-central/configuration/rest_err"
)

func (m *movementService) DeleteMovement(domain domain.MovementDomain) *rest_err.RestErr {
	if err := m.repository.Delete(domain); err != nil {
		return err
	}

	return nil
}
