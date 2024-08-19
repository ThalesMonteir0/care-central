package controller

import (
	"github.com/ThalesMonteir0/care-central/internal/adapter/input/model/request"
	"github.com/ThalesMonteir0/care-central/internal/application/domain"
	"github.com/ThalesMonteir0/care-central/pkg/rest_err"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (s *sessionController) CreateFormSession(ctx *gin.Context) {
	var formSessionRequest request.FormSessionRequest
	sessionID, errConv := strconv.Atoi(ctx.Param("session_id"))
	if errConv != nil {
		errRest := rest_err.NewInternalServerError("unable get session_id")
		ctx.JSON(errRest.Code, errRest.Message)
		return
	}

	if err := ctx.ShouldBindJSON(&formSessionRequest); err != nil {
		errRest := rest_err.NewBadRequestError("unable get body")
		ctx.JSON(errRest.Code, errRest.Message)
		return
	}

	sessionDomain := domain.SessionDomain{
		ID:            sessionID,
		SessionReport: formSessionRequest.SessionReport,
		Obs:           formSessionRequest.Obs,
	}

	if err := s.service.CreateFormSession(sessionDomain); err != nil {
		ctx.JSON(err.Code, err.Message)
		return
	}

	ctx.JSON(http.StatusOK, "form session and obs updated successful")
}
