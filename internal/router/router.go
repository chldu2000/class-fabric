package router

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"medicineApp/internal/controller"
	"medicineApp/internal/dao"
)

// RouterSet 路由注入
var RouterSet = wire.NewSet(wire.Struct(new(Router), "*"))

// Router 路由管理器
type Router struct {
	ConsumerDao      *dao.ConsumerDao
	AdministratorDao *dao.AdministratorDao
	DistributorDao   *dao.DistributorDao
	ManufacturerDao  *dao.ManufacturerDao
	LoginAPI         *controller.Login
	ConsumerAPI      *controller.Consumer
	AdministratorAPI *controller.Administrator
	ManufacturerAPI  *controller.Manufacturer
	DistributorAPI   *controller.Distributor
	ContractAPI *controller.ContractController
}

// Register 注册路由
func (a *Router) Register(app *gin.Engine) {
	a.registerAPI(app)
}
