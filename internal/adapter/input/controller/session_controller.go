package controller

import (
	"github.com/ThalesMonteir0/care-central/internal/application/port/input"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type sessionController struct {
	service input.SessionServiceInterface
	logger  *zap.Logger
}

type SessionControllerInterface interface {
	GetSession(ctx *gin.Context)
	CreateSession(ctx *gin.Context)
	UpdateSession(ctx *gin.Context)
	Delete(ctx *gin.Context)
	CreateFormSession(ctx *gin.Context)
}

func NewSessionController(service input.SessionServiceInterface, logger *zap.Logger) SessionControllerInterface {
	return &sessionController{
		service: service,
		logger:  logger,
	}
}
