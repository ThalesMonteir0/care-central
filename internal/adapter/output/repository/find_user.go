package repository

import (
	"github.com/ThalesMonteir0/care-central/internal/adapter/output/converter"
	"github.com/ThalesMonteir0/care-central/internal/application/domain"
	"github.com/ThalesMonteir0/care-central/pkg/rest_err"
)

func (ur *userRepository) FindUserByEmail(userDomain domain.UserDomain) (*domain.UserDomain, *domain.ClinicDomain, *rest_err.RestErr) {
	userEntity := converter.ConvertUserDomainToEntity(userDomain)

	if err := ur.db.Preload("Clinic").First(&userEntity, "email = ?", userEntity.Email).Error; err != nil {
		return nil, nil, rest_err.NewBadRequestError("unable get user")
	}

	return converter.ConverterUserEntityToDomain(userEntity), converter.ConvertUserEntityToClinicDomain(userEntity), nil
}

func (ur *userRepository) FindUserByID(userDomain domain.UserDomain) (*domain.UserDomain, *domain.ClinicDomain, *rest_err.RestErr) {
	userEntity := converter.ConvertUserDomainToEntity(userDomain)

	if err := ur.db.Preload("Clinic").First(&userEntity, userEntity.ID).Error; err != nil {
		return nil, nil, rest_err.NewBadRequestError("unable get user")
	}

	return converter.ConverterUserEntityToDomain(userEntity), converter.ConvertUserEntityToClinicDomain(userEntity), nil
}
