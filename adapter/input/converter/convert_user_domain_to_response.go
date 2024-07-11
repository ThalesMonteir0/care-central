package converter

import (
	"github.com/ThalesMonteir0/care-central/adapter/input/model/response"
	"github.com/ThalesMonteir0/care-central/application/domain"
)

func ConvertUserDomainToResponse(userDomain domain.UserDomain) response.UserResponse {
	return response.UserResponse{
		Name:     userDomain.Name,
		Email:    userDomain.Email,
		Role:     userDomain.Role,
		ClinicID: userDomain.ClinicID,
	}
}
