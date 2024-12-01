package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/project-sistem-voucher/api/handler"
	"github.com/project-sistem-voucher/config"
	"github.com/project-sistem-voucher/manager"
	"github.com/project-sistem-voucher/middleware"
	"github.com/sirupsen/logrus"
)

func SetupRouter(router *gin.Engine) error {

	router.Use(middleware.LogRequestMiddleware(logrus.New()))

	infraManager := manager.NewInfraManager(config.Cfg)
	serviceManager := manager.NewRepoManager(infraManager)
	repoManager := manager.NewServiceManager(serviceManager)

	voucherHandler := handler.NewVoucherHandler(repoManager.VoucherService())

	v1 := router.Group("/api/v1")
	{
		sistemVoucher := v1.Group("/management-voucher")
		{
			sistemVoucher.POST("/create", voucherHandler.CreateVoucher)
			sistemVoucher.DELETE("/delete/:id", voucherHandler.DeleteVoucher)
			sistemVoucher.PUT("update/:id", voucherHandler.UpdateVoucher)
			sistemVoucher.GET("list", voucherHandler.GetVouchers)
			sistemVoucher.GET("/redeem", voucherHandler.GetVouchersForRedeem)
		}
	}

	return router.Run()

}
