package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/project-sistem-voucher/middleware"
	"github.com/sirupsen/logrus"
)

func SetupRouter(router *gin.Engine) error {

	router.Use(middleware.LogRequestMiddleware(logrus.New()))

	// infraManager := manager.NewInfraManager(config.Cfg)
	// serviceManager := manager.NewRepoManager(infraManager)
	// repoManager := manager.NewServiceManager(serviceManager)

	// v1 := router.Group("/api/v1")
	// {
	// 	sistemVoucher := v1.Group("/sakupay")
	// 	{
	// 	}
	// }

	return router.Run()

}
