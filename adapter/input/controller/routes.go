package controller

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(app *gin.Engine, userController UserControllerInterface, patientController PatientControllerInterface) {
	//users
	app.POST("/user", userController.CreateUser)
	app.GET("/user/:id", userController.GetUser)
	app.DELETE("/user/:id", userController.DeleteUser)
	app.POST("/login", userController.Login)

	//patients
	app.GET("/patient/:clinic_id", patientController.getPatient)
	app.POST("/patient", patientController.CreatePatient)

}
