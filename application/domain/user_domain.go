package domain

import (
	"github.com/ThalesMonteir0/care-central/configuration/rest_err"
	"golang.org/x/crypto/bcrypt"
)

type UserDomain struct {
	ID       int
	Name     string
	Email    string
	Password string
	Active   bool
	ClinicID int
	Role     string
}

func (ud *UserDomain) EncryptedPassword() *rest_err.RestErr {
	passwordEncrypted, err := bcrypt.GenerateFromPassword([]byte(ud.Password), bcrypt.DefaultCost)
	if err != nil {
		restErr := rest_err.NewInternalServerError("unable encrypted password")
		return restErr
	}

	ud.Password = string(passwordEncrypted)

	return nil
}
