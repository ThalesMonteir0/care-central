package controller

import (
	"github.com/ThalesMonteir0/care-central/application/port/input"
	"github.com/gin-gonic/gin"
)

type sessionController struct {
	service input.SessionServiceInterface
}

type SessionControllerInterface interface {
	GetSession(ctx *gin.Context)
	CreateSession(ctx *gin.Context)
	UpdateSession(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

func NewSessionController(service input.SessionServiceInterface) SessionControllerInterface {
	return &sessionController{
		service: service,
	}
}
