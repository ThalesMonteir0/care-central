package controller

import (
	"github.com/ThalesMonteir0/care-central/internal/adapter/input/model/request"
	"github.com/ThalesMonteir0/care-central/internal/application/domain"
	"github.com/ThalesMonteir0/care-central/internal/application/port/input"
	"github.com/ThalesMonteir0/care-central/pkg/rest_err"
	"github.com/gin-gonic/gin"
	"net/http"
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
	var createPixRequest request.CreatePixRequest

	if err := ctx.ShouldBindJSON(&createPixRequest); err != nil {
		errRest := rest_err.NewBadRequestError("Unable get body request - all fields required")
		ctx.JSON(errRest.Code, errRest.Message)
		return
	}

	pixDomain := domain.CreatePixDomain{
		DebtorsCPF:  createPixRequest.CpfDebtor,
		DebtorsName: createPixRequest.NomeDebtor,
		Value:       createPixRequest.ValueSession,
		SessionID:   createPixRequest.SessionID,
	}

	pixCopiaEcola, err := c.service.CreatePix(pixDomain)
	if err != nil {
		ctx.JSON(err.Code, err.Message)
		return
	}

	ctx.JSON(http.StatusCreated, pixCopiaEcola)
}
