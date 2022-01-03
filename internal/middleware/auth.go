package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"medicineApp/internal/config"
	"medicineApp/internal/dao"
	"medicineApp/pkg/cryptox"
	"medicineApp/pkg/logger"
	"medicineApp/pkg/warpper"
)

func wrapUserAuthContext(c *gin.Context, userID string, authority int) {
	warpper.SetUserID(c, userID)
	warpper.SetUserAuthority(c, authority)
	ctx := c.Request.Context()
	ctx = logger.NewUserIDContext(ctx, userID)
	c.Request = c.Request.WithContext(ctx)
}

func AuthMiddlewareConsumer(consumerDao *dao.ConsumerDao) gin.HandlerFunc {
	if !config.C.JWT.Enable {
		return func(c *gin.Context) {
			c.Next()
		}
	}

	return func(c *gin.Context) {
		ctx := c.Request.Context()

		claims, err := cryptox.ParseToken(warpper.GetToken(c))
		if err != nil {
			warpper.ResError(c, warpper.ErrInvalidToken)
			return
		}

		res, err := consumerDao.Query(ctx, claims.Name)
		if res == nil {
			warpper.ResError(c, warpper.ErrInvalidUser)
			return
		}
		c.Next()
	}
}

func AuthMiddlewareAdmin(adminDao *dao.AdministratorDao) gin.HandlerFunc {
	if !config.C.JWT.Enable {
		return func(c *gin.Context) {
			c.Next()
		}
	}

	return func(c *gin.Context) {
		fmt.Println("权限中间件")
		ctx := c.Request.Context()

		claims, err := cryptox.ParseToken(warpper.GetToken(c))
		if err != nil {
			warpper.ResError(c, warpper.ErrInvalidToken)
			return
		}

		res, err := adminDao.Query(ctx, claims.Name)
		if res == nil {
			warpper.ResError(c, warpper.ErrInvalidUser)
			return
		}
		c.Next()
	}
}

func AuthMiddlewareManufacturer(manufacturerDao *dao.ManufacturerDao) gin.HandlerFunc {
	if !config.C.JWT.Enable {
		return func(c *gin.Context) {
			c.Next()
		}
	}

	return func(c *gin.Context) {
		ctx := c.Request.Context()

		claims, err := cryptox.ParseToken(warpper.GetToken(c))
		if err != nil {
			warpper.ResError(c, warpper.ErrInvalidToken)
			return
		}

		res, err := manufacturerDao.Query(ctx, claims.Name)
		if res == nil {
			warpper.ResError(c, warpper.ErrInvalidUser)
			return
		}
		c.Next()
	}
}

func AuthMiddlewareDistributor(distributorDao *dao.DistributorDao) gin.HandlerFunc {
	if !config.C.JWT.Enable {
		return func(c *gin.Context) {
			c.Next()
		}
	}

	return func(c *gin.Context) {
		ctx := c.Request.Context()

		claims, err := cryptox.ParseToken(warpper.GetToken(c))
		if err != nil {
			warpper.ResError(c, warpper.ErrInvalidToken)
			return
		}

		res, err := distributorDao.Query(ctx, claims.Name)
		if res == nil {
			warpper.ResError(c, warpper.ErrInvalidUser)
			return
		}
		c.Next()
	}
}
