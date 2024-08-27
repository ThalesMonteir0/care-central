package service

import (
	"github.com/ThalesMonteir0/care-central/internal/adapter/input/model/request"
	"github.com/ThalesMonteir0/care-central/internal/application/domain"
	"github.com/ThalesMonteir0/care-central/pkg/rest_err"
)

func (s *sessionService) GetSessions(domain domain.SessionDomain, filters request.SessionFilters) ([]domain.SessionDomain, *rest_err.RestErr) {

	sessions, err := s.repository.FindSessions(domain, filters)
	if err != nil {
		return nil, err
	}

	return sessions, nil
}
