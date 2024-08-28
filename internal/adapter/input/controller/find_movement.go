package controller

import (
	"github.com/ThalesMonteir0/care-central/internal/adapter/input/converter"
	"github.com/ThalesMonteir0/care-central/internal/adapter/input/model/request"
	"github.com/ThalesMonteir0/care-central/internal/adapter/input/model/response"
	"github.com/ThalesMonteir0/care-central/internal/application/domain"
	"github.com/ThalesMonteir0/care-central/pkg/rest_err"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (m *movementController) FindAll(ctx *gin.Context) {
	filters := getMovementsFilters(ctx)
	clinicID, errConv := strconv.Atoi(ctx.Param("clinic_id"))
	if errConv != nil {
		errRest := rest_err.NewInternalServerError("unable convert clinic_id")
		ctx.JSON(errRest.Code, errRest.Message)
		return
	}

	movementDomain := domain.MovementDomain{
		ClinicID: clinicID,
	}

	movements, err := m.service.FindAllMovements(movementDomain, filters)
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

func getMovementsFilters(ctx *gin.Context) request.MovementsFilters {
	dtInitial := ctx.Query("dt_initial")
	dtFinal := ctx.Query("dt_final")
	Type := ctx.Query("type")

	return request.MovementsFilters{
		Type:      Type,
		DtInitial: dtInitial,
		DtFinal:   dtFinal,
	}
}
