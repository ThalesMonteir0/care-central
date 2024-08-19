package output

import (
	"github.com/ThalesMonteir0/care-central/internal/application/domain"
	"github.com/ThalesMonteir0/care-central/pkg/rest_err"
)

type UserRepositoryInterface interface {
	CreateUser(userDomain domain.UserDomain) *rest_err.RestErr
	UpdateUser(userDomain domain.UserDomain) *rest_err.RestErr
	DeleteUser(userDomain domain.UserDomain) *rest_err.RestErr
	FindUserByEmail(userDomain domain.UserDomain) (*domain.UserDomain, *domain.ClinicDomain, *rest_err.RestErr)
	FindUserByID(userDomain domain.UserDomain) (*domain.UserDomain, *domain.ClinicDomain, *rest_err.RestErr)
}
