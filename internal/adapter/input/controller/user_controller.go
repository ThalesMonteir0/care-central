package controller

import (
	"github.com/ThalesMonteir0/care-central/internal/application/port/input"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type userController struct {
	service input.UserServiceInterface
	logger  *zap.Logger
}

type UserControllerInterface interface {
	GetUser(ctx *gin.Context)
	CreateUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
	Login(ctx *gin.Context)
}

func NewUserController(service input.UserServiceInterface, logger *zap.Logger) UserControllerInterface {
	return &userController{
		service: service,
		logger:  logger,
	}
}
