package dao

import (
	"context"
	"github.com/google/wire"
	"gorm.io/gorm"
	"medicineApp/internal/entity"
)

// AdministratorDaoSet 注入 DI
var AdministratorDaoSet = wire.NewSet(wire.Struct(new(AdministratorDao), "*"))

// AdministratorDao Administrator 表相关的数据库操作
type AdministratorDao struct {
	DB *gorm.DB
}

// LoginAdminResBody 需要返回给前端的数据
type LoginAdminResBody struct {
	SystemAccount string `json:"system_account"`
	Token         string `json:"token"`
}

func (ad *AdministratorDao) GetByAccount(ctx context.Context, account string) (*entity.Administrator, error) {
	admin := new(entity.Administrator)
	db := ad.DB.Model(&entity.Administrator{})
	err := db.Where("system_account = ?", account).First(admin).Error
	if err != nil {
		return nil, err
	}
	return admin, nil
}

func (ad *AdministratorDao) Query(ctx context.Context, account string) (*entity.Administrator, error) {
	admin := new(entity.Administrator)
	db := ad.DB.Model(&entity.Administrator{})
	err := db.Where("system_account = ?", account).First(admin).Error
	if err != nil {
		return nil, err
	}
	return admin, nil
}
