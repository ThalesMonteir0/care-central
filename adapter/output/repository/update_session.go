package repository

import (
	"github.com/ThalesMonteir0/care-central/adapter/output/converter"
	"github.com/ThalesMonteir0/care-central/application/domain"
	"github.com/ThalesMonteir0/care-central/configuration/rest_err"
)

func (s *sessionRepository) UpdateSession(domain domain.SessionDomain) *rest_err.RestErr {
	sessionEntity := converter.ConverterSessionDomainToEntity(domain)

	if result := s.db.Save(&sessionEntity); result.Error != nil {
		return rest_err.NewInternalServerError("unable update session")
	}

	return nil
}
