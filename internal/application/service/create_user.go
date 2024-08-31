package service

import (
	"encoding/json"
	"github.com/ThalesMonteir0/care-central/internal/application/domain"
	"github.com/ThalesMonteir0/care-central/pkg/rest_err"
	"go.uber.org/zap"
	"strings"
)

func (us *userService) CreateUser(domain domain.UserDomain) *rest_err.RestErr {
	domainBytes, _ := json.Marshal(domain)
	us.logger.Info("inciando fluxo do userService",
		zap.String("userDomain", string(domainBytes)),
	)

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
