package service

import (
	"github.com/ThalesMonteir0/care-central/application/domain"
	"github.com/ThalesMonteir0/care-central/configuration/rest_err"
)

func (s *sessionService) DeleteSession(domain domain.SessionDomain) *rest_err.RestErr {
	if err := s.repository.DeleteSession(domain); err != nil {
		return err
	}

	return nil
}
