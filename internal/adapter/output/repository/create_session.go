package repository

import (
	"github.com/ThalesMonteir0/care-central/internal/adapter/output/converter"
	"github.com/ThalesMonteir0/care-central/internal/application/domain"
	"github.com/ThalesMonteir0/care-central/pkg/rest_err"
)

func (s *sessionRepository) CreateSession(domain domain.SessionDomain) *rest_err.RestErr {
	sessionEntity := converter.ConverterSessionDomainToEntity(domain)

	if result := s.db.Create(&sessionEntity); result.Error != nil {
		return rest_err.NewInternalServerError("unable create session")
	}

	return nil
}
