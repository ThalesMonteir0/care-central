package controller

import (
	"github.com/ThalesMonteir0/care-central/adapter/input/converter"
	"github.com/ThalesMonteir0/care-central/adapter/input/model/response"
	"github.com/ThalesMonteir0/care-central/application/domain"
	"github.com/ThalesMonteir0/care-central/configuration/rest_err"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (m *movementController) FindAll(ctx *gin.Context) {
	clinicID, errConv := strconv.Atoi(ctx.Param("clinic_id"))
	if errConv != nil {
		errRest := rest_err.NewInternalServerError("unable convert clinic_id")
		ctx.JSON(errRest.Code, errRest.Message)
		return
	}

	movementDomain := domain.MovementDomain{
		ClinicID: clinicID,
	}

	movements, err := m.service.FindAllMovements(movementDomain)
	if err != nil {
		ctx.JSON(err.Code, err.Message)
		return
	}

	movementsResponse := convertMovementsDomainToResponses(movements)

	ctx.JSON(http.StatusOK, movementsResponse)
}

func convertMovementsDomainToResponses(movementsDomain []domain.MovementDomain) (movementsEntities []response.MovementResponse) {
	for _, movement := range movementsDomain {
		movementEntity := converter.ConvertMovementDomainToResponse(movement)
		movementsEntities = append(movementsEntities, movementEntity)
	}

	return
}
