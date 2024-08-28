package repository

import (
	"fmt"
	"github.com/ThalesMonteir0/care-central/internal/adapter/input/model/request"
	"github.com/ThalesMonteir0/care-central/internal/adapter/output/converter"
	"github.com/ThalesMonteir0/care-central/internal/adapter/output/model/entity"
	"github.com/ThalesMonteir0/care-central/internal/application/domain"
	"github.com/ThalesMonteir0/care-central/pkg/rest_err"
)

func (m *movementRepository) FindAll(domain domain.MovementDomain, filters request.MovementsFilters) ([]domain.MovementDomain, *rest_err.RestErr) {
	clinicID := domain.ClinicID
	var movements []entity.Movement
	whereString := getFiltersFindAllMovements(filters, clinicID)

	if tx := m.db.Find(&movements, whereString); tx.Error != nil {
		return nil, rest_err.NewBadRequestError("uanble get movements by clinic_id")
	}

	return converter.ConvertMovementsEntitiesToDomains(movements), nil
}

func getFiltersFindAllMovements(filters request.MovementsFilters, clinicID int) string {
	where := fmt.Sprintf("clinic_id = %d", clinicID)

	if filters.Type != "" {
		where += fmt.Sprintf(" AND movement_type_id = %s", filters.Type)
	}

	if filters.DtInitial != "" || filters.DtFinal != "" {
		if filters.DtInitial != "" && filters.DtFinal != "" {
			where += fmt.Sprintf(" AND created_at between '%s 00:00:00' AND '%s 23:59:59'", filters.DtInitial, filters.DtFinal)
		} else if filters.DtInitial != "" {
			where += fmt.Sprintf(" AND created_at between '%s 00:00:00' AND '%s 23:59:59'", filters.DtInitial, filters.DtInitial)
		} else {
			where += fmt.Sprintf(" AND created_at between '%s 00:00:00' AND '%s 23:59:59'", filters.DtFinal, filters.DtFinal)
		}
	}

	return where
}
