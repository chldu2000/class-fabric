package router

import (
	"github.com/gin-gonic/gin"
	"medicineApp/internal/fabricUtils"
	"medicineApp/internal/middleware"
)

// RegisterAPI 路由列表
func (a *Router) registerAPI(app *gin.Engine) {
	//app.Use(middleware.CORSMiddleware())
	g := app.Group("/api")
	v1 := g.Group("/v1")
	{
		v1.POST("/login", a.LoginAPI.Login)
		v1.GET("/test", middleware.AuthMiddlewareConsumer(a.ConsumerDao), a.ConsumerAPI.Test)
		v1.POST("/login/admin", middleware.CORSMiddleware(), a.AdministratorAPI.LoginAdmin)
		v1.POST("/login/consumer", middleware.CORSMiddleware(), a.ConsumerAPI.LoginConsumer)
		v1.POST("/login/distributor", middleware.CORSMiddleware(), a.DistributorAPI.LoginDistributor)
		v1.POST("/login/manufacturer", middleware.CORSMiddleware(), a.ManufacturerAPI.LoginManufacturer)
		v1.POST("/register/consumer", middleware.CORSMiddleware(), a.ConsumerAPI.RegisterConsumer)
		v1.POST("/register/distributor", middleware.CORSMiddleware(), a.DistributorAPI.RegisterDistributor)
		v1.POST("/register/manufacturer", middleware.CORSMiddleware(), a.ManufacturerAPI.RegisterManufacturer)
		v1.PUT("/pass/manufacturer", middleware.CORSMiddleware(), middleware.AuthMiddlewareAdmin(a.AdministratorDao), a.ManufacturerAPI.PassManufacturer)
		v1.PUT("/pass/consumer", middleware.CORSMiddleware(), middleware.AuthMiddlewareAdmin(a.AdministratorDao), a.ConsumerAPI.PassConsumer)
		v1.PUT("/pass/distributor", middleware.CORSMiddleware(), middleware.AuthMiddlewareAdmin(a.AdministratorDao), a.DistributorAPI.PassDistributor)
		v1.GET("/pass/consumer", middleware.AuthMiddlewareAdmin(a.AdministratorDao), a.ConsumerAPI.GetPassingConsumer)
		v1.GET("/pass/distributor", middleware.AuthMiddlewareAdmin(a.AdministratorDao), a.DistributorAPI.GetPassingDistributor)
		v1.GET("/pass/manufacturer", middleware.AuthMiddlewareAdmin(a.AdministratorDao), a.ManufacturerAPI.GetPassingManufacturer)
	
		v1.POST("/medcine", a.ContractAPI.Add)
		v1.GET("/medicine", a.ContractAPI.Query)
		v1.PUT("/medicine", a.ContractAPI.Update)
		v1.DELETE("/medicine", a.ContractAPI.Delete)

		v1.GET("/medicine/history", fabricUtils.QueryMedicineHistoryByCode)
		v1.PUT("/medicine/trade/propose", fabricUtils.TradePropose)
		v1.PUT("/medicine/trade/receive", fabricUtils.TradeReceive)
	}
}
