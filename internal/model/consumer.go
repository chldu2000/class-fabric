package model

import (
	"context"
	"github.com/google/wire"
	"github.com/jinzhu/copier"
	"medicineApp/internal/config"
	"medicineApp/internal/dao"
	"medicineApp/internal/entity"
	"medicineApp/internal/schema"
	"medicineApp/pkg/cryptox"
	"medicineApp/pkg/warpper"
)

// ConsumerModelSet ConsumerModel 注入 DI
var ConsumerModelSet = wire.NewSet(wire.Struct(new(ConsumerModel), "*"))

// ConsumerModel 处理登录的主要逻辑
type ConsumerModel struct {
	ConsumerDao *dao.ConsumerDao
}

func (consumerModel *ConsumerModel) Test(ctx context.Context) ([]schema.TestHandlerResBodySchema, error) {
	res, err := consumerModel.ConsumerDao.Test(ctx)
	if err != nil {
		return nil, err
	}
	length := len(*(res))
	var result []schema.TestHandlerResBodySchema
	for i := 0; i < length; i++ {
		result = append(result, schema.TestHandlerResBodySchema{
			Name:     (*res)[i].Name,
			Phone:    (*res)[i].Phone,
			IdNumber: (*res)[i].IdNumber,
			Password: (*res)[i].Password,
		})
	}
	return result, nil
}

// Login 登录方法
func (cm *ConsumerModel) Login(ctx context.Context, name, password string) (*schema.LoginConsumerResBodySchema, error) {

	consumer, err := cm.ConsumerDao.GetByName(ctx, name)
	if err != nil {
		return nil, err
	}

	if password != consumer.Password {
		return nil, warpper.ErrInvalidPassword
	}

	if consumer.State != 1 {
		return nil, warpper.ErrInvalidState
	}

	token := ""
	if config.C.JWT.Enable {
		token, err = cryptox.GenerateToken(consumer.IdNumber)
		if err != nil {
			return nil, warpper.ErrCanNotGenerateToken
		}
	}

	return &schema.LoginConsumerResBodySchema{
		Name:     consumer.Name,
		IdNumber: consumer.IdNumber,
		Token:    token,
	}, nil
}

// Register 注册方法
func (cm *ConsumerModel) Register(ctx context.Context, data schema.RegisterConsumerReqBodySchema) (*schema.RegisterConsumerResBodySchema, error) {
	res := new(schema.RegisterConsumerResBodySchema)
	consumer, err := cm.ConsumerDao.Register(ctx, data)
	if err != nil {
		return nil, err
	}
	err = copier.Copy(res, consumer)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Pass 注册方法
func (cm *ConsumerModel) Pass(ctx context.Context, data schema.PassReqBodySchema) (*schema.PassConsumerResBodySchema, error) {
	res := new(schema.PassConsumerResBodySchema)
	consumer, err := cm.ConsumerDao.Pass(ctx, data)
	if err != nil {
		return nil, err
	}
	err = copier.Copy(res, consumer)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetPassing 拿到审核列表
func (cm *ConsumerModel) GetPassing(ctx context.Context) (*[]entity.Consumer, error) {
	consumer, err := cm.ConsumerDao.GetPassing(ctx)
	if err != nil {
		return nil, err
	}
	return consumer, nil
}
