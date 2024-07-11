package service

import (
	"github.com/ThalesMonteir0/care-central/application/domain"
	"github.com/ThalesMonteir0/care-central/configuration/rest_err"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

func (us *userService) UserLogin(userDomain domain.UserDomain) (*domain.UserDomain, *domain.ClinicDomain, string, *rest_err.RestErr) {
	userDomain.Email = strings.ToUpper(userDomain.Email)
	user, clinic, _ := us.repository.FindUserByEmail(userDomain)
	if user == nil {
		return nil, nil, "", rest_err.NewBadRequestError("user not exists")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userDomain.Password)); err != nil {
		return nil, nil, "", rest_err.NewBadRequestError("password incorrect")
	}

	token, err := user.GenerateToken()
	if err != nil {
		return nil, nil, "", err
	}

	return user, clinic, token, nil
}
