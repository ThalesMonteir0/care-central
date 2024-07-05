package repository

import (
	"github.com/ThalesMonteir0/care-central/adapter/output/converter"
	"github.com/ThalesMonteir0/care-central/application/domain"
	"github.com/ThalesMonteir0/care-central/configuration/rest_err"
	"time"
)

func (ur *userRepository) CreateUser(userDomain domain.UserDomain) *rest_err.RestErr {
	userEntity := converter.ConvertUserDomainToEntity(userDomain)
	userEntity.CreatedAt = time.Now()

	if err := ur.db.Create(&userEntity).Error; err != nil {
		return rest_err.NewInternalServerError("unable create user")
	}

	return nil
}
