package converter

import (
	"github.com/ThalesMonteir0/care-central/internal/adapter/input/model/response"
	"github.com/ThalesMonteir0/care-central/internal/application/domain"
)

func ConvertUserAndClinicDomainToResponse(userDomain domain.UserDomain, clinicDomain domain.ClinicDomain, token string) response.UserLoginResponse {
	return response.UserLoginResponse{
		Name:          userDomain.Name,
		Email:         userDomain.Email,
		Role:          userDomain.Role,
		ClinicID:      userDomain.ClinicID,
		ClinicFantasy: clinicDomain.Fantasy,
		Token:         token,
	}
}
