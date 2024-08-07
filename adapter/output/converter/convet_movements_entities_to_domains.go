package converter

import (
	"github.com/ThalesMonteir0/care-central/adapter/output/model/entity"
	"github.com/ThalesMonteir0/care-central/application/domain"
)

func ConvertMovementsEntitiesToDomains(movementsEntities []entity.Movement) []domain.MovementDomain {
	var movementsDomain []domain.MovementDomain

	for _, movementEntity := range movementsEntities {
		movementDomain := domain.MovementDomain{
			ID:             int(movementEntity.ID),
			ClinicID:       movementEntity.ClinicID,
			TypeMovementID: movementEntity.MovementTypeID,
			Value:          movementEntity.Value,
			Description:    movementEntity.Description,
			CreatedAt:      movementEntity.CreatedAt,
		}

		movementsDomain = append(movementsDomain, movementDomain)
	}

	return movementsDomain
}
