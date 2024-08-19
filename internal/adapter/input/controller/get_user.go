package controller

import (
	"github.com/ThalesMonteir0/care-central/internal/adapter/input/converter"
	"github.com/ThalesMonteir0/care-central/internal/application/domain"
	"github.com/ThalesMonteir0/care-central/pkg/rest_err"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (uc *userController) GetUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		errRest := rest_err.NewInternalServerError("unable get user id")
		ctx.JSON(errRest.Code, errRest.Message)
		return
	}

	userDomain := domain.UserDomain{ID: id}

	user, errService := uc.service.FindUserByID(userDomain)
	if errService != nil {
		errRest := rest_err.NewInternalServerError("unable get user id")
		ctx.JSON(errRest.Code, errRest.Message)
		return
	}

	ctx.JSON(http.StatusOK, converter.ConvertUserDomainToResponse(*user))
	return
}
