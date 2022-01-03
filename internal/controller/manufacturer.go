package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"medicineApp/internal/model"
	"medicineApp/internal/schema"
	"medicineApp/pkg/logger"
	"medicineApp/pkg/warpper"
)

// ManufacturerSet Distributor DI
var ManufacturerSet = wire.NewSet(wire.Struct(new(Manufacturer), "*"))

// Manufacturer 结构体
type Manufacturer struct {
	ManufacturerModel *model.ManufacturerModel
}

// LoginManufacturer  Login 登录方法
func (ad *Manufacturer) LoginManufacturer(c *gin.Context) {
	ctx := c.Request.Context()

	var data schema.LoginCompanyReqBodySchema
	if err := warpper.ParseJSON(c, &data); err != nil {
		warpper.ResError(c, err)
		return
	}

	res, err := ad.ManufacturerModel.Login(ctx, data.CompanyName, data.Password)
	if err != nil {
		warpper.ResError(c, err)
		return
	}

	ctx = logger.NewUserIDContext(ctx, res.CompanyName)
	ctx = logger.NewTagContext(ctx, "__login-manufacturer__")
	logger.WithContext(ctx).Info("登入系统")

	warpper.ResSuccess(c, res)
}

// RegisterManufacturer 注册方法
func (cm *Manufacturer) RegisterManufacturer(c *gin.Context) {
	ctx := c.Request.Context()

	var data schema.RegisterManufacturerReqBodySchema
	if err := warpper.ParseJSON(c, &data); err != nil {
		warpper.ResError(c, err)
		return
	}

	res, err := cm.ManufacturerModel.Register(ctx, data)
	if err != nil {
		warpper.ResError(c, err)
		return
	}
	warpper.ResSuccess(c, res)
}

// PassManufacturer 注册方法
func (cm *Manufacturer) PassManufacturer(c *gin.Context) {
	ctx := c.Request.Context()

	var data schema.PassReqBodySchema
	if err := warpper.ParseJSON(c, &data); err != nil {
		warpper.ResError(c, err)
		return
	}

	res, err := cm.ManufacturerModel.Pass(ctx, data)
	if err != nil {
		warpper.ResError(c, err)
		return
	}
	warpper.ResSuccess(c, res)
}

// GetPassingManufacturer 拿到待审批
func (cm *Manufacturer) GetPassingManufacturer(c *gin.Context) {
	ctx := c.Request.Context()
	res, err := cm.ManufacturerModel.GetPassing(ctx)
	if err != nil {
		warpper.ResError(c, err)
		return
	}
	warpper.ResSuccess(c, res)
}
