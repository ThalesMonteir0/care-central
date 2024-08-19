package service

import (
	"github.com/ThalesMonteir0/care-central/internal/application/domain"
	"github.com/ThalesMonteir0/care-central/pkg/rest_err"
)

func (us *userService) DeleteUser(domain domain.UserDomain) *rest_err.RestErr {
	if err := us.repository.DeleteUser(domain); err != nil {
		return err
	}

	return nil
}
