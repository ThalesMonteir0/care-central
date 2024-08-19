package controller

import (
	"github.com/ThalesMonteir0/care-central/internal/application/port/input"
	"github.com/gin-gonic/gin"
)

type createPixController struct {
	service input.CreatePixService
}

type CreatePixController interface {
	CreatePix(ctx *gin.Context)
}

func NewCreatePixController(service input.CreatePixService) CreatePixController {
	return &createPixController{
		service: service,
	}
}

func (c *createPixController) CreatePix(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}
