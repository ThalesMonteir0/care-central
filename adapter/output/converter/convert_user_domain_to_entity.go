package converter

import (
	"github.com/ThalesMonteir0/care-central/adapter/output/model/entity"
	"github.com/ThalesMonteir0/care-central/application/domain"
)

func ConvertUserDomainToEntity(domain domain.UserDomain) entity.User {
	return entity.User{
		Name:     domain.Name,
		Email:    domain.Email,
		Password: domain.Password,
		Role:     domain.Role,
		Active:   domain.Active,
		ClinicID: domain.ClinicID,
	}
}
