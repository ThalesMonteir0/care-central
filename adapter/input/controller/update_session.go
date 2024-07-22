package controller

import (
	"github.com/ThalesMonteir0/care-central/adapter/input/model/request"
	"github.com/ThalesMonteir0/care-central/application/domain"
	"github.com/ThalesMonteir0/care-central/configuration/rest_err"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (s *sessionController) UpdateSession(ctx *gin.Context) {
	var sessionRequest request.SessionUpdateRequest
	clinicID, errConv := strconv.Atoi(ctx.Param("clinic_id"))
	if errConv != nil {
		errRest := rest_err.NewInternalServerError("unable convert clinic_id")
		ctx.JSON(errRest.Code, errRest.Message)
		return
	}

	sessionID, errConv := strconv.Atoi(ctx.Param("session_id"))
	if errConv != nil {
		errRest := rest_err.NewInternalServerError("unable convert session_id")
		ctx.JSON(errRest.Code, errRest.Message)
		return
	}

	if err := ctx.ShouldBindJSON(&sessionRequest); err != nil {
		errRest := rest_err.NewInternalServerError("unable parse body")
		ctx.JSON(errRest.Code, errRest.Message)
		return
	}

	sessionDomain := domain.SessionDomain{
		ID:           sessionID,
		ClinicID:     clinicID,
		Paid:         sessionRequest.Paid,
		ValueSession: sessionRequest.ValueSession,
		DtSession:    sessionRequest.DtSession,
	}

	if err := s.service.UpdateSession(sessionDomain); err != nil {
		ctx.JSON(err.Code, err.Message)
		return
	}

	ctx.JSON(http.StatusOK, "session updated successful")
	return
}
