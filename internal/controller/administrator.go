package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"medicineApp/internal/model"
	"medicineApp/internal/schema"
	"medicineApp/pkg/logger"
	"medicineApp/pkg/warpper"
)

// AdministratorSet Administrator DI
var AdministratorSet = wire.NewSet(wire.Struct(new(Administrator), "*"))

// Administrator 结构体
type Administrator struct {
	AdministratorModel *model.AdministratorModel
}

// LoginAdmin Login 登录方法
func (ad *Administrator) LoginAdmin(c *gin.Context) {
	ctx := c.Request.Context()

	var data schema.LoginAdminReqBodySchema
	if err := warpper.ParseJSON(c, &data); err != nil {
		warpper.ResError(c, err)
		return
	}

	res, err := ad.AdministratorModel.Login(ctx, data.SystemAccount, data.Password)
	if err != nil {
		warpper.ResError(c, err)
		return
	}

	ctx = logger.NewUserIDContext(ctx, res.SystemAccount)
	ctx = logger.NewTagContext(ctx, "__login-admin__")
	logger.WithContext(ctx).Info("登入系统")

	warpper.ResSuccess(c, res)
}
