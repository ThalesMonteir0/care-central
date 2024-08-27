package output

import (
	"github.com/ThalesMonteir0/care-central/internal/adapter/input/model/request"
	"github.com/ThalesMonteir0/care-central/internal/application/domain"
	"github.com/ThalesMonteir0/care-central/pkg/rest_err"
)

type SessionRepositoryInterface interface {
	FindSessionByID(domain.SessionDomain) (*domain.SessionDomain, *rest_err.RestErr)
	FindSessions(domain.SessionDomain, request.SessionFilters) ([]domain.SessionDomain, *rest_err.RestErr)
	CreateSession(domain.SessionDomain) *rest_err.RestErr
	UpdateSession(domain.SessionDomain) *rest_err.RestErr
	DeleteSession(domain.SessionDomain) *rest_err.RestErr
}
