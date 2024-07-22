package service

import (
	"github.com/ThalesMonteir0/care-central/application/domain"
	"github.com/ThalesMonteir0/care-central/configuration/rest_err"
)

func (s *sessionService) GetSessions(domain domain.SessionDomain) ([]domain.SessionDomain, *rest_err.RestErr) {

	sessions, err := s.repository.FindSessions(domain)
	if err != nil {
		return nil, err
	}

	return sessions, nil
}
