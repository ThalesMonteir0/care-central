package service

import (
	"github.com/ThalesMonteir0/care-central/application/domain"
	"github.com/ThalesMonteir0/care-central/configuration/rest_err"
)

func (us *userService) FindUserByEmail(domain domain.UserDomain) (*domain.UserDomain, *rest_err.RestErr) {
	return nil, nil
}
