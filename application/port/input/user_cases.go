package input

import (
	"github.com/ThalesMonteir0/care-central/application/domain"
	"github.com/ThalesMonteir0/care-central/configuration/rest_err"
)

type UserServiceInterface interface {
	CreateUser(domain.UserDomain) *rest_err.RestErr
	UpdateUser(domain.UserDomain) *rest_err.RestErr
	DeleteUser(domain.UserDomain) *rest_err.RestErr
	FindUserByEmail(domain.UserDomain) (*domain.UserDomain, *rest_err.RestErr)
	FindUserByID(domain.UserDomain) (*domain.UserDomain, *rest_err.RestErr)
}
