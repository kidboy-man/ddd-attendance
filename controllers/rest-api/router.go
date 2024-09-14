package restapi

import (
	"github.com/gin-gonic/gin"
	restapiv1 "github.com/kidboy-man/ddd-attendance/controllers/rest-api/v1"
)

func Init() {
	router := gin.Default()
	v1Route := router.Group("/api/v1")

	accountController := restapiv1.NewAccountController()

	// Register routes
	v1Route.POST("/account/register", accountController.RegisterEmployee)

	// Start the server
	router.Run(":9000")
}
