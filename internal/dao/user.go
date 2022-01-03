package dao

import (
	"context"
	"medicineApp/internal/entity"

	"github.com/google/wire"
	"gorm.io/gorm"
)

// UserDaoSet 注入 DI
var UserDaoSet = wire.NewSet(wire.Struct(new(UserDao), "*"))

// UserDao users 表相关的数据库操作
type UserDao struct {
	DB *gorm.DB
}

// UserQueryParams 查询用户的参数
type UserQueryParams struct {
	UID   string
	Phone string
}

// Query 根据给定条件查询用户
func (u *UserDao) Query(ctx context.Context, params UserQueryParams) (*[]entity.User, error) {
	result := new([]entity.User)
	db := u.DB.Model(&entity.User{})

	if v := params.UID; v != "" {
		db = db.Where("uid = ?", v)
	}

	if v := params.Phone; v != "" {
		db = db.Where("phone = ?", v)
	}

	err := db.Find(result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
