package repository

import (
	"github.com/ThalesMonteir0/care-central/application/domain"
	"github.com/ThalesMonteir0/care-central/configuration/rest_err"
)

func (s *sessionRepository) FindSessionByID(domain domain.SessionDomain) (*domain.SessionDomain, *rest_err.RestErr) {
	//TODO implement me
	panic("implement me")
}

func (s *sessionRepository) FindSessions(domain domain.SessionDomain) ([]domain.SessionDomain, *rest_err.RestErr) {
	//TODO implement me
	panic("implement me")
}
