package repository

import (
	"github.com/ThalesMonteir0/care-central/adapter/output/converter"
	"github.com/ThalesMonteir0/care-central/adapter/output/model/entity"
	"github.com/ThalesMonteir0/care-central/application/domain"
	"github.com/ThalesMonteir0/care-central/configuration/rest_err"
)

func (m *movementRepository) FindAll(domain domain.MovementDomain) ([]domain.MovementDomain, *rest_err.RestErr) {
	clinicID := domain.ClinicID
	var movements []entity.Movement

	if tx := m.db.Find(&movements, "clinic_id = ?", clinicID); tx.Error != nil {
		return nil, rest_err.NewBadRequestError("uanble get movements by clinic_id")
	}

	return converter.ConvertMovementsEntitiesToDomains(movements), nil
}
