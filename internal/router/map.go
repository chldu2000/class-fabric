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
	}
	// v2部分才是药品生产商部分接口
	v2 := g.Group("/v2")
	{
		v2.GET("/test", middleware.AuthMiddlewareConsumer(a.ConsumerDao), a.ConsumerAPI.Test)
		v2.POST("/login/admin", middleware.CORSMiddleware(), a.AdministratorAPI.LoginAdmin)
		v2.POST("/login/consumer", middleware.CORSMiddleware(), a.ConsumerAPI.LoginConsumer)
		v2.POST("/login/distributor", middleware.CORSMiddleware(), a.DistributorAPI.LoginDistributor)
		v2.POST("/login/manufacturer", middleware.CORSMiddleware(), a.ManufacturerAPI.LoginManufacturer)
		v2.POST("/register/consumer", middleware.CORSMiddleware(), a.ConsumerAPI.RegisterConsumer)
		v2.POST("/register/distributor", middleware.CORSMiddleware(), a.DistributorAPI.RegisterDistributor)
		v2.POST("/register/manufacturer", middleware.CORSMiddleware(), a.ManufacturerAPI.RegisterManufacturer)
		v2.PUT("/pass/manufacturer", middleware.CORSMiddleware(), middleware.AuthMiddlewareAdmin(a.AdministratorDao), a.ManufacturerAPI.PassManufacturer)
		v2.PUT("/pass/consumer", middleware.CORSMiddleware(), middleware.AuthMiddlewareAdmin(a.AdministratorDao), a.ConsumerAPI.PassConsumer)
		v2.PUT("/pass/distributor", middleware.CORSMiddleware(), middleware.AuthMiddlewareAdmin(a.AdministratorDao), a.DistributorAPI.PassDistributor)
		v2.GET("/pass/consumer", middleware.AuthMiddlewareAdmin(a.AdministratorDao), a.ConsumerAPI.GetPassingConsumer)
		v2.GET("/pass/distributor", middleware.AuthMiddlewareAdmin(a.AdministratorDao), a.DistributorAPI.GetPassingDistributor)
		v2.GET("/pass/manufacturer", middleware.AuthMiddlewareAdmin(a.AdministratorDao), a.ManufacturerAPI.GetPassingManufacturer)
	}
	v3 := g.Group("/v3")
	{
		v3.GET("/medicine/query/history", middleware.CORSMiddleware(), fabricUtils.QueryMedicineHistoryByCode)
		v3.PUT("/medicine/trade/propose", fabricUtils.TradePropose)
		v3.PUT("/medicine/trade/receive", fabricUtils.TradeReceive)
	}
}
