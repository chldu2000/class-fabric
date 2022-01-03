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

// DistributorModelSet DistributorModel 注入 DI
var DistributorModelSet = wire.NewSet(wire.Struct(new(DistributorModel), "*"))

// DistributorModel 处理登录的主要逻辑
type DistributorModel struct {
	DistributorDao *dao.DistributorDao
}

// Login 登录方法
func (cm *DistributorModel) Login(ctx context.Context, name, password string) (*schema.LoginCompanyResBodySchema, error) {

	comp, err := cm.DistributorDao.GetByComp(ctx, name)
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
func (cm *DistributorModel) Register(ctx context.Context, data schema.RegisterDistributorReqBodySchema) (*schema.RegisterDistributorResBodySchema, error) {
	res := new(schema.RegisterDistributorResBodySchema)
	distributor, err := cm.DistributorDao.Register(ctx, data)
	if err != nil {
		return nil, err
	}
	err = copier.Copy(res, distributor)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Pass 注册方法
func (cm *DistributorModel) Pass(ctx context.Context, data schema.PassReqBodySchema) (*schema.PassDistributorResBodySchema, error) {
	res := new(schema.PassDistributorResBodySchema)
	distributor, err := cm.DistributorDao.Pass(ctx, data)
	if err != nil {
		return nil, err
	}
	err = copier.Copy(res, distributor)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetPassing 拿到审核列表
func (cm *DistributorModel) GetPassing(ctx context.Context) (*[]entity.Distributor, error) {
	consumer, err := cm.DistributorDao.GetPassing(ctx)
	if err != nil {
		return nil, err
	}
	return consumer, nil
}
