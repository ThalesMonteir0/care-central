package controller

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(app *gin.Engine,
	userController UserControllerInterface,
	patientController PatientControllerInterface,
	sessionController SessionControllerInterface) {

	//users
	app.POST("/user", userController.CreateUser)
	app.GET("/user/:id", userController.GetUser)
	app.DELETE("/user/:id", userController.DeleteUser)
	app.POST("/login", userController.Login)

	//patients
	app.GET("/patient/:clinic_id", patientController.getPatient)
	app.POST("/patient", patientController.CreatePatient)
	app.PUT("/patient/:id", patientController.UpdatePatient)
	app.DELETE("/patient/:id", patientController.DeletePatient)

	//sessions
	app.POST("/session", sessionController.CreateSession)
}
