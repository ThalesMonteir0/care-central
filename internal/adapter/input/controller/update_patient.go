package controller

import (
	"github.com/ThalesMonteir0/care-central/internal/adapter/input/model/request"
	"github.com/ThalesMonteir0/care-central/internal/application/domain"
	"github.com/ThalesMonteir0/care-central/pkg/rest_err"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (p *patientController) UpdatePatient(ctx *gin.Context) {
	var patientRequestUpdate request.PatientRequestUpdate
	patientID, errConvert := strconv.Atoi(ctx.Param("id"))
	if errConvert != nil {
		errRest := rest_err.NewBadRequestError("unable get patient id")
		ctx.JSON(errRest.Code, errRest.Message)
		return
	}

	if err := ctx.ShouldBindJSON(&patientRequestUpdate); err != nil {
		errRest := rest_err.NewBadRequestError("unable parse json")
		ctx.JSON(errRest.Code, errRest.Message)
		return
	}

	patientDomain := domain.PatientDomain{
		ID:             patientID,
		Name:           patientRequestUpdate.Name,
		ResponsibleCPF: patientRequestUpdate.ResponsibleCPF,
		DateOfBirth:    patientRequestUpdate.DateOfBirth,
	}

	if err := p.service.UpdatePatient(patientDomain); err != nil {
		ctx.JSON(err.Code, err.Message)
		return
	}

	ctx.JSON(http.StatusOK, "patient updated")
}
