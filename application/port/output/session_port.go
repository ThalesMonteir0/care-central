package output

import (
	"github.com/ThalesMonteir0/care-central/application/domain"
	"github.com/ThalesMonteir0/care-central/configuration/rest_err"
)

type SessionRepositoryInterface interface {
	FindSessionByID(domain.SessionDomain) (*domain.SessionDomain, *rest_err.RestErr)
	FindSessions(domain.SessionDomain) ([]domain.SessionDomain, *rest_err.RestErr)
	CreateSession(domain.SessionDomain) *rest_err.RestErr
	UpdateSession(domain.SessionDomain) *rest_err.RestErr
	DeleteSession(domain.SessionDomain) *rest_err.RestErr
}
