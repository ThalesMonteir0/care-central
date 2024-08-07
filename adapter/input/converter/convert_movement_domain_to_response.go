package converter

import (
	"github.com/ThalesMonteir0/care-central/adapter/input/model/response"
	"github.com/ThalesMonteir0/care-central/application/domain"
)

func ConvertMovementDomainToResponse(domain domain.MovementDomain) response.MovementResponse {
	return response.MovementResponse{
		ID:             domain.ID,
		ClinicID:       domain.ClinicID,
		MovementTypeID: domain.TypeMovementID,
		Description:    domain.Description,
		CreatedAt:      domain.CreatedAt,
	}
}
