package controller

import (
	"github.com/ThalesMonteir0/care-central/application/domain"
	"github.com/ThalesMonteir0/care-central/configuration/rest_err"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (m *movementController) DeleteMovement(ctx *gin.Context) {
	clinicID, errConv := strconv.Atoi(ctx.Param("clinic_id"))
	if errConv != nil {
		restErr := rest_err.NewInternalServerError("unable get clinic_id")
		ctx.JSON(restErr.Code, restErr.Message)
		return
	}
	movementID, errConv := strconv.Atoi(ctx.Param("movement_id"))
	if errConv != nil {
		restErr := rest_err.NewInternalServerError("unable get movement_id")
		ctx.JSON(restErr.Code, restErr.Message)
		return
	}

	movementDomain := domain.MovementDomain{
		ID:       movementID,
		ClinicID: clinicID,
	}

	if err := m.service.DeleteMovement(movementDomain); err != nil {
		ctx.JSON(err.Code, err.Message)
		return
	}

	ctx.JSON(http.StatusOK, "movement deleted successful")
	return
}
