package output

import (
	"github.com/ThalesMonteir0/care-central/application/domain"
	"github.com/ThalesMonteir0/care-central/configuration/rest_err"
)

type UserRepositoryInterface interface {
	CreateUser(userDomain domain.UserDomain) *rest_err.RestErr
	UpdateUser(userDomain domain.UserDomain) *rest_err.RestErr
	DeleteUser(userDomain domain.UserDomain) *rest_err.RestErr
	FindUserByEmail(userDomain domain.UserDomain) *domain.UserDomain
}
