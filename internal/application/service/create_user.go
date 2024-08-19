package service

import (
	"github.com/ThalesMonteir0/care-central/internal/application/domain"
	"github.com/ThalesMonteir0/care-central/pkg/rest_err"
	"strings"
)

func (us *userService) CreateUser(domain domain.UserDomain) *rest_err.RestErr {
	if user, _, _ := us.repository.FindUserByEmail(domain); user != nil {
		return rest_err.NewBadRequestError("user already exists")
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
