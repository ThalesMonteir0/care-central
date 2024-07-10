package service

import (
	"github.com/ThalesMonteir0/care-central/application/domain"
	"github.com/ThalesMonteir0/care-central/configuration/rest_err"
	"strings"
)

func (us *userService) CreateUser(domain domain.UserDomain) *rest_err.RestErr {
	if user, _ := us.repository.FindUserByEmail(domain); user != nil {
		return rest_err.NewBadRequestError("user exists")
	}

	domain.Email = strings.ToUpper(domain.Email)
	if err := domain.EncryptedPassword(); err != nil {
		return err
	}

	if err := us.repository.CreateUser(domain); err != nil {
		return err
	}

	return nil
}
