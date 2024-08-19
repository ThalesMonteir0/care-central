package controller

import (
	"github.com/ThalesMonteir0/care-central/internal/application/domain"
	"github.com/ThalesMonteir0/care-central/pkg/rest_err"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (p *patientController) DeletePatient(ctx *gin.Context) {
	patientID, errConv := strconv.Atoi(ctx.Param("id"))
	if errConv != nil {
		errRest := rest_err.NewInternalServerError("unable convert id")
		ctx.JSON(errRest.Code, errRest.Message)
		return
	}

	patientDomain := domain.PatientDomain{ID: patientID}

	if err := p.service.DeletePatient(patientDomain); err != nil {
		ctx.JSON(err.Code, err.Message)
		return
	}

	ctx.JSON(http.StatusOK, "patient deleted successful")
	return
}
