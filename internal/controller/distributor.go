package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"medicineApp/internal/model"
	"medicineApp/internal/schema"
	"medicineApp/pkg/logger"
	"medicineApp/pkg/warpper"
)

// DistributorSet Distributor DI
var DistributorSet = wire.NewSet(wire.Struct(new(Distributor), "*"))

// Distributor 结构体
type Distributor struct {
	DistributorModel *model.DistributorModel
}

// LoginDistributor  Login 登录方法
func (ad *Distributor) LoginDistributor(c *gin.Context) {
	ctx := c.Request.Context()

	var data schema.LoginCompanyReqBodySchema
	if err := warpper.ParseJSON(c, &data); err != nil {
		warpper.ResError(c, err)
		return
	}

	res, err := ad.DistributorModel.Login(ctx, data.CompanyName, data.Password)
	if err != nil {
		warpper.ResError(c, err)
		return
	}

	ctx = logger.NewUserIDContext(ctx, res.CompanyName)
	ctx = logger.NewTagContext(ctx, "__login-distributor__")
	logger.WithContext(ctx).Info("登入系统")

	warpper.ResSuccess(c, res)
}

// RegisterDistributor Register方法
func (ad *Distributor) RegisterDistributor(c *gin.Context) {
	ctx := c.Request.Context()

	var data schema.RegisterDistributorReqBodySchema
	if err := warpper.ParseJSON(c, &data); err != nil {
		warpper.ResError(c, err)
		return
	}

	res, err := ad.DistributorModel.Register(ctx, data)
	if err != nil {
		warpper.ResError(c, err)
		return
	}

	warpper.ResSuccess(c, res)
}

// PassDistributor 注册方法
func (cm *Distributor) PassDistributor(c *gin.Context) {
	ctx := c.Request.Context()

	var data schema.PassReqBodySchema
	if err := warpper.ParseJSON(c, &data); err != nil {
		warpper.ResError(c, err)
		return
	}

	res, err := cm.DistributorModel.Pass(ctx, data)
	if err != nil {
		warpper.ResError(c, err)
		return
	}
	warpper.ResSuccess(c, res)
}

// GetPassingDistributor 拿到待审批
func (cm *Distributor) GetPassingDistributor(c *gin.Context) {
	ctx := c.Request.Context()
	res, err := cm.DistributorModel.GetPassing(ctx)
	if err != nil {
		warpper.ResError(c, err)
		return
	}
	warpper.ResSuccess(c, res)
}
