package repository

import (
	"github.com/ThalesMonteir0/care-central/adapter/output/converter"
	"github.com/ThalesMonteir0/care-central/adapter/output/model/entity"
	"github.com/ThalesMonteir0/care-central/application/domain"
	"github.com/ThalesMonteir0/care-central/configuration/rest_err"
)

func (s *sessionRepository) FindSessionByID(domain domain.SessionDomain) (*domain.SessionDomain, *rest_err.RestErr) {
	//TODO implement me
	panic("implement me")
}

func (s *sessionRepository) FindSessions(domain domain.SessionDomain) ([]domain.SessionDomain, *rest_err.RestErr) {
	clinicID := domain.ClinicID
	var sessionEntity []entity.Session

	if result := s.db.Find(&sessionEntity, "clinic_id = ?", clinicID); result.Error != nil {
		return nil, rest_err.NewBadRequestError("unable get sessions by clinic_id")
	}

	return converter.ConverterSessionsEntitiesToDomains(sessionEntity), nil
}
