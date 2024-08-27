package repository

import (
	"fmt"
	"github.com/ThalesMonteir0/care-central/internal/adapter/input/model/request"
	"github.com/ThalesMonteir0/care-central/internal/adapter/output/converter"
	"github.com/ThalesMonteir0/care-central/internal/adapter/output/model/entity"
	"github.com/ThalesMonteir0/care-central/internal/application/domain"
	"github.com/ThalesMonteir0/care-central/pkg/rest_err"
)

func (s *sessionRepository) FindSessionByID(domain domain.SessionDomain) (*domain.SessionDomain, *rest_err.RestErr) {
	sessionID := domain.ID
	var sessionEntity entity.Session

	if result := s.db.First(&sessionEntity, sessionID); result.Error != nil {
		return nil, rest_err.NewBadRequestError("unable get session by session_id")
	}

	return converter.ConverterSessionEntityToDomain(sessionEntity), nil
}

func (s *sessionRepository) FindSessions(domain domain.SessionDomain, filters request.SessionFilters) ([]domain.SessionDomain, *rest_err.RestErr) {
	clinicID := domain.ClinicID
	var sessionEntity []entity.Session
	whereString := getFindSessionsFilters(filters, clinicID)

	if result := s.db.Find(&sessionEntity, whereString); result.Error != nil {
		return nil, rest_err.NewBadRequestError("unable get sessions by clinic_id")
	}

	return converter.ConverterSessionsEntitiesToDomains(sessionEntity), nil
}

func getFindSessionsFilters(filters request.SessionFilters, clinicID int) string {
	where := fmt.Sprintf("clinic_id = %d", clinicID)

	if filters.PatientID != "" {
		where += fmt.Sprintf(" AND patient_id = %s", filters.PatientID)
	}

	if filters.Paid != "" {
		where += fmt.Sprintf(" AND paid = %s", filters.Paid)
	}

	if filters.DtSessionInitial != "" || filters.DtSessionFinal != "" {
		if filters.DtSessionInitial != "" && filters.DtSessionFinal != "" {
			where += fmt.Sprintf(" AND dt_session between '%s 00:00:00' and '%s 23:59:59'", filters.DtSessionInitial, filters.DtSessionFinal)
		} else if filters.DtSessionInitial != "" {
			where += fmt.Sprintf(" AND dt_session between '%s 00:00:00' and '%s 23:59:59'", filters.DtSessionInitial, filters.DtSessionInitial)
		} else {
			where += fmt.Sprintf(" AND dt_session between '%s 00:00:00' and '%s 23:59:59'", filters.DtSessionFinal, filters.DtSessionFinal)
		}
	}

	return where
}
