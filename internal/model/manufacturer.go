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

// ManufacturerModelSet ManufacturerModel 注入 DI
var ManufacturerModelSet = wire.NewSet(wire.Struct(new(ManufacturerModel), "*"))

// ManufacturerModel 处理登录的主要逻辑
type ManufacturerModel struct {
	ManufacturerDao *dao.ManufacturerDao
}

// Login 登录方法
func (cm *ManufacturerModel) Login(ctx context.Context, name, password string) (*schema.LoginCompanyResBodySchema, error) {

	comp, err := cm.ManufacturerDao.GetByComp(ctx, name)
	if err != nil {
		return nil, err
	}

	if password != comp.Password {
		return nil, warpper.ErrInvalidPassword
	}

	if comp.State != 1 {
		return nil, warpper.ErrInvalidState
	}

	token := ""
	if config.C.JWT.Enable {
		//这里到时候可以换成公钥
		token, err = cryptox.GenerateToken(comp.CompanyName)
		if err != nil {
			return nil, warpper.ErrCanNotGenerateToken
		}
	}
	return &schema.LoginCompanyResBodySchema{
		CompanyName: comp.CompanyName,
		Token:       token,
	}, nil
}

// Register 注册方法
func (cm *ManufacturerModel) Register(ctx context.Context, data schema.RegisterManufacturerReqBodySchema) (*schema.RegisterManufacturerResBodySchema, error) {
	res := new(schema.RegisterManufacturerResBodySchema)
	manufacturer, err := cm.ManufacturerDao.Register(ctx, data)
	if err != nil {
		return nil, err
	}
	err = copier.Copy(res, manufacturer)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Pass 注册方法
func (cm *ManufacturerModel) Pass(ctx context.Context, data schema.PassReqBodySchema) (*schema.PassManufacturerResBodySchema, error) {
	res := new(schema.PassManufacturerResBodySchema)
	manufacturer, err := cm.ManufacturerDao.Pass(ctx, data)
	if err != nil {
		return nil, err
	}
	err = copier.Copy(res, manufacturer)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetPassing 拿到审核列表
func (cm *ManufacturerModel) GetPassing(ctx context.Context) (*[]entity.Manufacturer, error) {
	consumer, err := cm.ManufacturerDao.GetPassing(ctx)
	if err != nil {
		return nil, err
	}
	return consumer, nil
}
