package service

import (
	"github.com/ThalesMonteir0/care-central/application/domain"
	"github.com/ThalesMonteir0/care-central/configuration/rest_err"
)

func (us *userService) FindUserByID(domain domain.UserDomain) (*domain.UserDomain, *rest_err.RestErr) {
	user, _, err := us.repository.FindUserByID(domain)
	if err != nil {
		return nil, err
	}

	return user, nil
}
