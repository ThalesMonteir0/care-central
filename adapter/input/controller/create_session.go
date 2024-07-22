package controller

import (
	"github.com/ThalesMonteir0/care-central/adapter/input/model/request"
	"github.com/ThalesMonteir0/care-central/application/domain"
	"github.com/ThalesMonteir0/care-central/configuration/rest_err"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *sessionController) CreateSession(ctx *gin.Context) {
	var sessionRequest request.SessionRequest

	if err := ctx.ShouldBindJSON(&sessionRequest); err != nil {
		errRest := rest_err.NewBadRequestError("all fields required")
		ctx.JSON(errRest.Code, errRest.Message)
		return
	}

	sessionDomain := domain.SessionDomain{
		PatientID:     sessionRequest.PatientID,
		ClinicID:      sessionRequest.ClinicID,
		Paid:          sessionRequest.Paid,
		SessionReport: sessionRequest.SessionReport,
		Obs:           sessionRequest.Obs,
		ValueSession:  sessionRequest.ValueSession,
		DtSession:     sessionRequest.DtSession,
	}

	if err := s.service.CreateSession(sessionDomain); err != nil {
		ctx.JSON(err.Code, err.Message)
		return
	}

	ctx.JSON(http.StatusCreated, "session created successful")
	return
}
