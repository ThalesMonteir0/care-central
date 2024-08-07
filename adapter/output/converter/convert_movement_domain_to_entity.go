package converter

import (
	"github.com/ThalesMonteir0/care-central/adapter/output/model/entity"
	"github.com/ThalesMonteir0/care-central/application/domain"
)

func ConvertMovementDomainToEntity(domain domain.MovementDomain) entity.Movement {
	movementEntity := entity.Movement{
		ClinicID:       domain.ClinicID,
		MovementTypeID: domain.TypeMovementID,
		Value:          domain.Value,
		Description:    domain.Description,
	}

	movementEntity.ID = uint(domain.ID)
	movementEntity.CreatedAt = domain.CreatedAt

	return movementEntity
}
