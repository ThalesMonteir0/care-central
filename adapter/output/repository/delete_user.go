package repository

import (
	"github.com/ThalesMonteir0/care-central/application/domain"
	"github.com/ThalesMonteir0/care-central/configuration/rest_err"
)

func (ur *userRepository) DeleteUser(userDomain domain.UserDomain) *rest_err.RestErr {
	return nil
}
