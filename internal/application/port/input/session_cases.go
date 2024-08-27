package input

import (
	"github.com/ThalesMonteir0/care-central/internal/adapter/input/model/request"
	"github.com/ThalesMonteir0/care-central/internal/application/domain"
	"github.com/ThalesMonteir0/care-central/pkg/rest_err"
)

type SessionServiceInterface interface {
	GetSessions(domain.SessionDomain, request.SessionFilters) ([]domain.SessionDomain, *rest_err.RestErr)
	CreateSession(domain.SessionDomain) *rest_err.RestErr
	UpdateSession(domain.SessionDomain) *rest_err.RestErr
	DeleteSession(domain.SessionDomain) *rest_err.RestErr
	CreateFormSession(domain.SessionDomain) *rest_err.RestErr
}
