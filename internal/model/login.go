package model

import (
	"context"
	"errors"
	"medicineApp/internal/entity"
	"regexp"

	"github.com/google/wire"
	"gorm.io/gorm"
	"medicineApp/internal/dao"
	"medicineApp/internal/schema"
	"medicineApp/pkg/cryptox"
	"medicineApp/pkg/logger"
	"medicineApp/pkg/warpper"
)

// LoginModelSet LoginModel 注入 DI
var LoginModelSet = wire.NewSet(wire.Struct(new(LoginModel), "*"))

// LoginModel 处理登录的主要逻辑
type LoginModel struct {
	UserDao *dao.UserDao
}

// 下面的部分用于通过手机号码或者 hrid 验证用户登录, 采用简单工厂模式
type iVerify interface {
	Verify(ctx context.Context, username string) (*[]entity.User, error)
}

type verifyByPhone struct {
	UserDao *dao.UserDao
}

func (v *verifyByPhone) Verify(ctx context.Context, username string) (*[]entity.User, error) {
	return v.UserDao.Query(ctx, dao.UserQueryParams{
		Phone: username,
	})
}

type verifyByHRID struct {
	UserDao *dao.UserDao
}

func (v *verifyByHRID) Verify(ctx context.Context, username string) (*[]entity.User, error) {
	return v.UserDao.Query(ctx, dao.UserQueryParams{
		UID: username,
	})
}

func verifyFactory(ctx context.Context, username string, userDao *dao.UserDao) (iVerify, error) {
	if len(username) == 4 {
		logger.WithContext(ctx).Infof("使用 hrid 登录 - %v", username)
		return &verifyByHRID{
			UserDao: userDao,
		}, nil
	}

	if match, _ := regexp.MatchString("^1[3-9]\\d{9}", username); match {
		logger.WithContext(ctx).Infof("使用电话号码登录 - %v", username)
		return &verifyByPhone{
			UserDao: userDao,
		}, nil
	}

	return nil, warpper.ErrInvalidUserName
}

// Login 登录方法
func (l *LoginModel) Login(ctx context.Context, username, password string) (*schema.LoginResBodySchema, error) {
	verify, err := verifyFactory(ctx, username, l.UserDao)
	if err != nil {
		return nil, err
	}

	res, err := verify.Verify(ctx, username)
	// 区分未找到错误
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, warpper.ErrInvalidUserName
	}
	if err != nil {
		return nil, err
	}

	user := (*res)[0]
	temp := cryptox.MD5(password)
	if temp != user.Password {
		return nil, warpper.ErrInvalidPassword
	}

	return &schema.LoginResBodySchema{
		UID:       user.UID,
		Authority: user.Authority,
	}, nil
}
