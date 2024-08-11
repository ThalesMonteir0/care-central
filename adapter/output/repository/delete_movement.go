package repository

import (
	"github.com/ThalesMonteir0/care-central/adapter/output/converter"
	"github.com/ThalesMonteir0/care-central/application/domain"
	"github.com/ThalesMonteir0/care-central/configuration/rest_err"
)

func (m *movementRepository) Delete(domain domain.MovementDomain) *rest_err.RestErr {
	movementEntity := converter.ConvertMovementDomainToEntity(domain)

	if tx := m.db.Delete(&movementEntity); tx.Error != nil {
		return rest_err.NewBadRequestError("unable delete movement")
	}

	return nil
}
