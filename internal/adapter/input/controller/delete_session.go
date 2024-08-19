package controller

import (
	"github.com/ThalesMonteir0/care-central/internal/application/domain"
	"github.com/ThalesMonteir0/care-central/pkg/rest_err"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (s *sessionController) Delete(ctx *gin.Context) {
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

	sessionDomain := domain.SessionDomain{
		ID:       sessionID,
		ClinicID: clinicID,
	}

	if err := s.service.DeleteSession(sessionDomain); err != nil {
		ctx.JSON(err.Code, err.Message)
		return
	}

	ctx.JSON(http.StatusOK, "session deleted successful")
	return
}
