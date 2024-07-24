package service

import (
	"github.com/ThalesMonteir0/care-central/application/domain"
	"github.com/ThalesMonteir0/care-central/configuration/rest_err"
)

func (s *sessionService) CreateFormSession(domain domain.SessionDomain) *rest_err.RestErr {
	session, err := s.repository.FindSessionByID(domain)
	if session == nil {
		if err != nil {
			return err
		}
		return rest_err.NewBadRequestError("session not found")
	}

	session.SessionReport = domain.SessionReport
	session.Obs = domain.Obs

	if err = s.repository.UpdateSession(*session); err != nil {
		return err
	}

	return nil
}
