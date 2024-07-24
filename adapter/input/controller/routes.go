package controller

import (
	"github.com/ThalesMonteir0/care-central/adapter/input/controller/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRoutes(app *gin.Engine,
	userController UserControllerInterface,
	patientController PatientControllerInterface,
	sessionController SessionControllerInterface) {

	//users
	app.POST("/user", middlewares.VerifyTokenMiddleware, userController.CreateUser)
	app.GET("/user/:id", middlewares.VerifyTokenMiddleware, userController.GetUser)
	app.DELETE("/user/:id", middlewares.VerifyTokenMiddleware, userController.DeleteUser)
	app.POST("/login", userController.Login)

	//patients
	app.GET("/patient/:clinic_id", middlewares.VerifyTokenMiddleware, patientController.getPatient)
	app.POST("/patient", middlewares.VerifyTokenMiddleware, patientController.CreatePatient)
	app.PUT("/patient/:id", middlewares.VerifyTokenMiddleware, patientController.UpdatePatient)
	app.DELETE("/patient/:id", middlewares.VerifyTokenMiddleware, patientController.DeletePatient)

	//sessions
	app.POST("/session", middlewares.VerifyTokenMiddleware, sessionController.CreateSession)
	app.GET("/session/:clinic_id", middlewares.VerifyTokenMiddleware, sessionController.GetSession)
	app.DELETE("/session/:clinic_id/:session_id", middlewares.VerifyTokenMiddleware, sessionController.Delete)
	app.PUT("/session/:clinic_id/:session_id", middlewares.VerifyTokenMiddleware, sessionController.UpdateSession)

	//form_session
	app.PUT("/form_session/session/:session_id", middlewares.VerifyTokenMiddleware, sessionController.CreateFormSession)
}
