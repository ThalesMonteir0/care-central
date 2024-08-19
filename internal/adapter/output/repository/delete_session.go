package repository

import (
	"github.com/ThalesMonteir0/care-central/internal/adapter/output/converter"
	"github.com/ThalesMonteir0/care-central/internal/application/domain"
	"github.com/ThalesMonteir0/care-central/pkg/rest_err"
)

func (s *sessionRepository) DeleteSession(domain domain.SessionDomain) *rest_err.RestErr {
	sessionEntity := converter.ConverterSessionDomainToEntity(domain)

	if result := s.db.Where("clinic_id = ?", sessionEntity.ClinicID).
		Delete(&sessionEntity); result.Error != nil {
		return rest_err.NewBadRequestError("unable delete session")
	}

	return nil
}
