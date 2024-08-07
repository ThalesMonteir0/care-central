package controller

import (
	"github.com/ThalesMonteir0/care-central/adapter/input/model/request"
	"github.com/ThalesMonteir0/care-central/application/domain"
	"github.com/ThalesMonteir0/care-central/configuration/rest_err"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (m *movementController) CreateMovement(ctx *gin.Context) {
	var movement request.MovementRequest

	if err := ctx.ShouldBindJSON(&movement); err != nil {
		errRest := rest_err.NewBadRequestError("todos os campos são obrigatorios")
		ctx.JSON(errRest.Code, errRest.Message)
		return
	}

	movementDomain := domain.MovementDomain{
		ClinicID:       movement.ClinicID,
		TypeMovementID: movement.MovementTypeID,
		Value:          movement.Value,
		Description:    movement.Description,
	}

	if err := m.service.CreateMovement(movementDomain); err != nil {
		ctx.JSON(err.Code, err.Message)
		return
	}

	ctx.JSON(http.StatusCreated, "Movement created successful")
}
