package controller

import (
	"github.com/ThalesMonteir0/care-central/internal/application/port/input"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type movementController struct {
	service input.MovementServiceInterface
	logger  *zap.Logger
}

type MovementControllerInterface interface {
	FindAll(ctx *gin.Context)
	CreateMovement(ctx *gin.Context)
	DeleteMovement(ctx *gin.Context)
}

func NewMovementController(service input.MovementServiceInterface, logger *zap.Logger) MovementControllerInterface {
	return &movementController{
		service: service,
		logger:  logger,
	}
}
