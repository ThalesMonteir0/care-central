package controller

import (
	"github.com/ThalesMonteir0/care-central/application/domain"
	"github.com/ThalesMonteir0/care-central/configuration/rest_err"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (uc *userController) DeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		errRest := rest_err.NewInternalServerError("unable get user ID")
		ctx.JSON(errRest.Code, errRest.Message)
		return
	}

	userDomain := domain.UserDomain{ID: id}

	if err := uc.service.DeleteUser(userDomain); err != nil {
		ctx.JSON(err.Code, err.Message)
		return
	}

	ctx.JSON(http.StatusOK, "user deleted successful")
	return
}
