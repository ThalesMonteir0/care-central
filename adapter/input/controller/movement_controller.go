package controller

import (
	"github.com/ThalesMonteir0/care-central/application/port/input"
	"github.com/gin-gonic/gin"
)

type movementController struct {
	service input.MovementServiceInterface
}

type MovementControllerInterface interface {
	FindAll(ctx *gin.Context)
	CreateMovement(ctx *gin.Context)
	DeleteMovement(ctx *gin.Context)
}

func NewMovementController(service input.MovementServiceInterface) MovementControllerInterface {
	return &movementController{
		service: service,
	}
}
