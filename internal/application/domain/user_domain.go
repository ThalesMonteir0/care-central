package domain

import (
	"github.com/ThalesMonteir0/care-central/pkg/rest_err"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

const (
	JWT_SECRET = "JWT_SECRET"
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

func (ud *UserDomain) GenerateToken() (string, *rest_err.RestErr) {
	secret := os.Getenv(JWT_SECRET)

	claims := jwt.MapClaims{
		"id":        ud.ID,
		"name":      ud.Name,
		"email":     ud.Email,
		"clinic_id": ud.ClinicID,
		"role":      ud.Role,
		"exp":       time.Now().Add(time.Hour * 3).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", rest_err.NewInternalServerError("error trying to generate token jwt")
	}

	return tokenString, nil
}
