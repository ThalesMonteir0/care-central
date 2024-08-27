package controller

import (
	"github.com/ThalesMonteir0/care-central/internal/adapter/input/model/request"
	"github.com/ThalesMonteir0/care-central/internal/application/domain"
	"github.com/ThalesMonteir0/care-central/pkg/rest_err"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (s *sessionController) GetSession(ctx *gin.Context) {
	sessionFilters := getFiltersSession(ctx)
	clinicID, errConv := strconv.Atoi(ctx.Param("clinic_id"))
	if errConv != nil {
		errRest := rest_err.NewInternalServerError("unable get clinic_id")
		ctx.JSON(errRest.Code, errRest.Message)
		return
	}

	sessionDomain := domain.SessionDomain{
		ClinicID: clinicID,
	}

	sessions, err := s.service.GetSessions(sessionDomain, sessionFilters)
	if err != nil {
		ctx.JSON(err.Code, err.Message)
		return
	}

	ctx.JSON(http.StatusOK, sessions)
	return
}

func getFiltersSession(ctx *gin.Context) request.SessionFilters {
	dtSessionInitial := ctx.Query("dt_session_initial")
	dtSessionFinal := ctx.Query("dt_session_final")
	patientID := ctx.Query("patient_id")
	paid := ctx.Query("paid")

	return request.SessionFilters{
		DtSessionInitial: dtSessionInitial,
		DtSessionFinal:   dtSessionFinal,
		PatientID:        patientID,
		Paid:             paid,
	}
}
