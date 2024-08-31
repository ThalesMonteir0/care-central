package controller

import (
	"github.com/ThalesMonteir0/care-central/internal/adapter/input/model/request"
	"github.com/ThalesMonteir0/care-central/internal/application/domain"
	"github.com/ThalesMonteir0/care-central/pkg/rest_err"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func (uc *userController) CreateUser(ctx *gin.Context) {
	uc.logger.Info("iniciando handler de createUser")

	var userRequest request.UserRequest

	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		uc.logger.Error("erro ao fazer o shouldBindJson",
			zap.String("err", err.Error()),
		)

		restErr := rest_err.NewBadRequestError("email,password and email is required. Password min 6 characters")
		ctx.JSON(restErr.Code, restErr.Message)
		return
	}

	userDomain := domain.UserDomain{
		Name:     userRequest.Name,
		Password: userRequest.Password,
		Email:    userRequest.Email,
		Active:   true,
		Role:     "ADMIN",
		ClinicID: userRequest.ClinicID,
	}

	if err := uc.service.CreateUser(userDomain); err != nil {
		ctx.JSON(err.Code, err.Message)
		return
	}

	ctx.JSON(http.StatusOK, "create user success")
	return
}
