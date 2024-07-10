package repository

import (
	"github.com/ThalesMonteir0/care-central/adapter/output/converter"
	"github.com/ThalesMonteir0/care-central/application/domain"
	"github.com/ThalesMonteir0/care-central/configuration/rest_err"
)

func (ur *userRepository) DeleteUser(userDomain domain.UserDomain) *rest_err.RestErr {
	userEntity := converter.ConvertUserDomainToEntity(userDomain)

	if err := ur.db.Delete(&userEntity).Error; err != nil {
		return rest_err.NewInternalServerError("unable delete user")
	}

	return nil
}
