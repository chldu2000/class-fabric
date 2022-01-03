package model

import (
	"context"
	"github.com/google/wire"
	"medicineApp/internal/config"
	"medicineApp/internal/dao"
	"medicineApp/internal/schema"
	"medicineApp/pkg/cryptox"
	"medicineApp/pkg/warpper"
)

// AdministratorModelSet AdministratorModel 注入 DI
var AdministratorModelSet = wire.NewSet(wire.Struct(new(AdministratorModel), "*"))

// AdministratorModel 处理登录的主要逻辑
type AdministratorModel struct {
	AdministratorDao *dao.AdministratorDao
}

// Login 登录方法
func (ad *AdministratorModel) Login(ctx context.Context, account, password string) (*schema.LoginAdminResBodySchema, error) {

	admin, err := ad.AdministratorDao.GetByAccount(ctx, account)
	if err != nil {
		return nil, err
	}

	if password != admin.Password {
		return nil, warpper.ErrInvalidPassword
	}

	token := ""
	if config.C.JWT.Enable {
		token, err = cryptox.GenerateToken(admin.SystemAccount)
		if err != nil {
			return nil, warpper.ErrCanNotGenerateToken
		}
	}

	return &schema.LoginAdminResBodySchema{
		SystemAccount: admin.SystemAccount,
		Token:         token,
	}, nil
}
