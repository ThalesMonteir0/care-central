package controller

import (
	"github.com/ThalesMonteir0/care-central/internal/application/port/input"
	"github.com/gin-gonic/gin"
)

type userController struct {
	service input.UserServiceInterface
}

type UserControllerInterface interface {
	GetUser(ctx *gin.Context)
	CreateUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
	Login(ctx *gin.Context)
}

func NewUserController(service input.UserServiceInterface) UserControllerInterface {
	return &userController{
		service: service,
	}
}
