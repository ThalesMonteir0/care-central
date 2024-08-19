package converter

import (
	"github.com/ThalesMonteir0/care-central/internal/adapter/output/model/entity"
	"github.com/ThalesMonteir0/care-central/internal/application/domain"
)

func ConvertUserDomainToEntity(domain domain.UserDomain) entity.User {
	userEntity := entity.User{
		Name:     domain.Name,
		Email:    domain.Email,
		Password: domain.Password,
		Role:     domain.Role,
		Active:   domain.Active,
		ClinicID: domain.ClinicID,
	}

	userEntity.ID = uint(domain.ID)

	return userEntity
}
