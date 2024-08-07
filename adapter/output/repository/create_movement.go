package repository

import (
	"github.com/ThalesMonteir0/care-central/adapter/output/converter"
	"github.com/ThalesMonteir0/care-central/application/domain"
	"github.com/ThalesMonteir0/care-central/configuration/rest_err"
)

func (m *movementRepository) Create(domain domain.MovementDomain) *rest_err.RestErr {
	movementEntity := converter.ConvertMovementDomainToEntity(domain)

	if tx := m.db.Create(&movementEntity); tx.Error != nil {
		return rest_err.NewInternalServerError("unable create movement")
	}

	return nil
}
