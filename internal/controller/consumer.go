package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"medicineApp/internal/model"
	"medicineApp/internal/schema"
	"medicineApp/pkg/logger"
	"medicineApp/pkg/warpper"
)

// ConsumerSet Consumer DI
var ConsumerSet = wire.NewSet(wire.Struct(new(Consumer), "*"))

// Consumer 结构体
type Consumer struct {
	ConsumerModel *model.ConsumerModel
}

func (consumer *Consumer) Test(c *gin.Context) {
	ctx := c.Request.Context()
	consumerList, err := consumer.ConsumerModel.Test(ctx)
	if err != nil {
		warpper.ResError(c, err)
		return
	}
	warpper.ResSuccess(c, consumerList)
}

// LoginConsumer 登录方法
func (cm *Consumer) LoginConsumer(c *gin.Context) {
	ctx := c.Request.Context()

	var data schema.LoginConsumerReqBodySchema
	if err := warpper.ParseJSON(c, &data); err != nil {
		warpper.ResError(c, err)
		return
	}

	res, err := cm.ConsumerModel.Login(ctx, data.Name, data.Password)
	if err != nil {
		warpper.ResError(c, err)
		return
	}

	ctx = logger.NewUserIDContext(ctx, res.IdNumber)
	ctx = logger.NewTagContext(ctx, "__login-consumer__")
	logger.WithContext(ctx).Info("登入系统")
	warpper.ResSuccess(c, res)
}

// RegisterConsumer 注册方法
func (cm *Consumer) RegisterConsumer(c *gin.Context) {
	ctx := c.Request.Context()

	var data schema.RegisterConsumerReqBodySchema
	if err := warpper.ParseJSON(c, &data); err != nil {
		warpper.ResError(c, err)
		return
	}

	res, err := cm.ConsumerModel.Register(ctx, data)
	if err != nil {
		warpper.ResError(c, err)
		return
	}
	warpper.ResSuccess(c, res)
}

// PassConsumer 注册方法
func (cm *Consumer) PassConsumer(c *gin.Context) {
	ctx := c.Request.Context()

	var data schema.PassReqBodySchema
	if err := warpper.ParseJSON(c, &data); err != nil {
		warpper.ResError(c, err)
		return
	}

	res, err := cm.ConsumerModel.Pass(ctx, data)
	if err != nil {
		warpper.ResError(c, err)
		return
	}
	warpper.ResSuccess(c, res)
}

// GetPassingConsumer 拿到待审批
func (cm *Consumer) GetPassingConsumer(c *gin.Context) {
	ctx := c.Request.Context()
	res, err := cm.ConsumerModel.GetPassing(ctx)
	if err != nil {
		warpper.ResError(c, err)
		return
	}
	warpper.ResSuccess(c, res)
}
