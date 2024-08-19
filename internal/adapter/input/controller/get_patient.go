package controller

import (
	"github.com/ThalesMonteir0/care-central/internal/application/domain"
	"github.com/ThalesMonteir0/care-central/pkg/rest_err"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (p *patientController) getPatient(ctx *gin.Context) {
	clinicID, errConvert := strconv.Atoi(ctx.Param("clinic_id"))
	if errConvert != nil {
		errRest := rest_err.NewInternalServerError("error converting id to int")
		ctx.JSON(errRest.Code, errRest.Message)
		return
	}

	patientDomain := domain.PatientDomain{ClinicID: clinicID}

	patient, err := p.service.FindPatient(patientDomain)
	if err != nil {
		ctx.JSON(err.Code, err.Message)
		return
	}

	ctx.JSON(http.StatusOK, patient)
	return
}
