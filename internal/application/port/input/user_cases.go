package input

import (
	"github.com/ThalesMonteir0/care-central/internal/application/domain"
	"github.com/ThalesMonteir0/care-central/pkg/rest_err"
)

type UserServiceInterface interface {
	CreateUser(domain.UserDomain) *rest_err.RestErr
	UpdateUser(domain.UserDomain) *rest_err.RestErr
	DeleteUser(domain.UserDomain) *rest_err.RestErr
	FindUserByID(domain.UserDomain) (*domain.UserDomain, *rest_err.RestErr)
	UserLogin(userDomain domain.UserDomain) (*domain.UserDomain, *domain.ClinicDomain, string, *rest_err.RestErr)
}
