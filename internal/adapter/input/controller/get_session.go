package controller

import (
	"github.com/ThalesMonteir0/care-central/internal/application/domain"
	"github.com/ThalesMonteir0/care-central/pkg/rest_err"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (s *sessionController) GetSession(ctx *gin.Context) {
	clinicID, errConv := strconv.Atoi(ctx.Param("clinic_id"))
	if errConv != nil {
		errRest := rest_err.NewInternalServerError("unable get clinic_id")
		ctx.JSON(errRest.Code, errRest.Message)
		return
	}

	sessionDomain := domain.SessionDomain{
		ClinicID: clinicID,
	}

	sessions, err := s.service.GetSessions(sessionDomain)
	if err != nil {
		ctx.JSON(err.Code, err.Message)
		return
	}

	ctx.JSON(http.StatusOK, sessions)
	return
}
