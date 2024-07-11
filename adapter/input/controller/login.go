package controller

import (
	"github.com/ThalesMonteir0/care-central/adapter/input/converter"
	"github.com/ThalesMonteir0/care-central/adapter/input/model/request"
	"github.com/ThalesMonteir0/care-central/application/domain"
	"github.com/ThalesMonteir0/care-central/configuration/rest_err"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (uc *userController) Login(ctx *gin.Context) {
	var userLogin request.UserLogin

	if err := ctx.ShouldBindJSON(&userLogin); err != nil {
		errRest := rest_err.NewBadRequestError("email and password is required")
		ctx.JSON(errRest.Code, errRest.Message)
		return
	}

	userDomain := domain.UserDomain{
		Email:    userLogin.Email,
		Password: userLogin.Password,
	}

	user, clinic, token, err := uc.service.UserLogin(userDomain)
	if err != nil {
		ctx.JSON(err.Code, err.Message)
		return
	}

	ctx.JSON(http.StatusOK, converter.ConvertUserAndClinicDomainToResponse(*user, *clinic, token))
}
