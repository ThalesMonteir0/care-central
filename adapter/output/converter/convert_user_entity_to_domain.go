package converter

import (
	"github.com/ThalesMonteir0/care-central/adapter/output/model/entity"
	"github.com/ThalesMonteir0/care-central/application/domain"
)

func ConverterUserEntityToDomain(userEntity entity.User) *domain.UserDomain {
	return &domain.UserDomain{
		ID:       int(userEntity.ID),
		Name:     userEntity.Name,
		Email:    userEntity.Email,
		Password: userEntity.Password,
		Active:   userEntity.Active,
		ClinicID: userEntity.ClinicID,
		Role:     userEntity.Role,
	}
}
