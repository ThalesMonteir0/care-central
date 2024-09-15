package controller

import (
	"github.com/ThalesMonteir0/care-central/internal/adapter/input/controller/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRoutes(app *gin.Engine,
	userController UserControllerInterface,
	patientController PatientControllerInterface,
	sessionController SessionControllerInterface,
	movementController MovementControllerInterface,
	createPixCtrl CreatePixController,
) {

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

	//movements
	app.GET("/movements/:clinic_id", middlewares.VerifyTokenMiddleware, movementController.FindAll)
	app.POST("movements/clinic/:clinic_id/", middlewares.VerifyTokenMiddleware, movementController.CreateMovement)
	app.DELETE("/movements/clinic/:clinic_id/movement/:movement_id", middlewares.VerifyTokenMiddleware, movementController.DeleteMovement)

	//PIX
	app.POST("/create_pix", middlewares.VerifyTokenMiddleware, createPixCtrl.CreatePix)

	app.GET("/health-check", HealthCheck)
}
