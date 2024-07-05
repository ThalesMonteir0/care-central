package controller

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(app *gin.Engine, userController UserControllerInterface) {

	app.POST("/user", userController.CreateUser)

}
