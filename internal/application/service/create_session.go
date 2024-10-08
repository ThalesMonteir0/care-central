package service

import (
	"github.com/ThalesMonteir0/care-central/internal/application/domain"
	"github.com/ThalesMonteir0/care-central/pkg/rest_err"
)

func (s *sessionService) CreateSession(domain domain.SessionDomain) *rest_err.RestErr {
	//todo: validar se patient by clinicID existe.

	if err := s.repository.CreateSession(domain); err != nil {
		return err
	}

	return nil
}
