package converter

import (
	"github.com/ThalesMonteir0/care-central/internal/adapter/output/model/entity"
	"github.com/ThalesMonteir0/care-central/internal/application/domain"
)

func ConvertUserEntityToClinicDomain(userEntity entity.User) *domain.ClinicDomain {
	return &domain.ClinicDomain{
		Fantasy:      userEntity.Clinic.Fantasy,
		CpfCnpj:      userEntity.Clinic.CpfCnpj,
		SocialReason: userEntity.Clinic.SocialReason,
	}
}
