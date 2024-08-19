package controller

import (
	"github.com/ThalesMonteir0/care-central/internal/adapter/input/model/request"
	"github.com/ThalesMonteir0/care-central/internal/application/domain"
	"github.com/ThalesMonteir0/care-central/pkg/rest_err"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (p *patientController) CreatePatient(ctx *gin.Context) {
	var patientRequest request.PatientRequest

	if err := ctx.ShouldBindJSON(&patientRequest); err != nil {
		errRest := rest_err.NewBadRequestError("all fields required, clinic_id not can null or 0")
		ctx.JSON(errRest.Code, errRest.Message)
		return
	}

	patientDomain := domain.PatientDomain{
		Name:           patientRequest.Name,
		ResponsibleCPF: patientRequest.ResponsibleCPF,
		DateOfBirth:    patientRequest.DateOfBirth,
		ClinicID:       patientRequest.ClinicID,
	}

	if err := p.service.CreatePatient(patientDomain); err != nil {
		ctx.JSON(err.Code, err.Message)
		return
	}

	ctx.JSON(http.StatusOK, "patient create successful")
}
