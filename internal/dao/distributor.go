package dao

import (
	"context"
	"github.com/google/wire"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"medicineApp/internal/entity"
	"medicineApp/internal/fabricUtils"
	"medicineApp/internal/schema"
	"strconv"
)

// DistributorDaoSet 注入 DI
var DistributorDaoSet = wire.NewSet(wire.Struct(new(DistributorDao), "*"))

// DistributorDao Distributor 表相关的数据库操作
type DistributorDao struct {
	DB *gorm.DB
}

func (dt *DistributorDao) GetByComp(ctx context.Context, name string) (*entity.Distributor, error) {
	comp := new(entity.Distributor)
	db := dt.DB.Model(&entity.Distributor{})
	err := db.Where("company_name = ?", name).First(comp).Error
	if err != nil {
		return nil, err
	}
	return comp, nil
}
func (dt *DistributorDao) Query(ctx context.Context, name string) (*entity.Distributor, error) {
	comp := new(entity.Distributor)
	db := dt.DB.Model(&entity.Distributor{})
	err := db.Where("company_name = ?", name).First(comp).Error
	if err != nil {
		return nil, err
	}
	return comp, nil
}
func (cm *DistributorDao) Register(ctx context.Context, data schema.RegisterDistributorReqBodySchema) (*entity.Distributor, error) {
	distributor := new(entity.Distributor)
	db := cm.DB.Model(&entity.Distributor{})
	err := copier.Copy(distributor, data)
	if err != nil {
		return nil, err
	}
	err = db.Save(distributor).Error
	if err != nil {
		return nil, err
	}
	return distributor, nil
}

func (cm *DistributorDao) Pass(ctx context.Context, data schema.PassReqBodySchema) (*entity.Distributor, error) {
	distributor := new(entity.Distributor)
	//wallet := "测试用钱包"//这里留有接口

	err := fabricUtils.RegisterAndEnroll(strconv.Itoa(int(data.Id)), "Supplier")
	if err != nil {
		return nil, err
	}
	wallet, err := gateway.NewFileSystemWallet("wallet")
	if err != nil {
		return nil, err
	}
	if !wallet.Exists(strconv.Itoa(int(data.Id)) + "@Supplier") {
		err := fabricUtils.PopulateWallet(wallet, strconv.Itoa(int(data.Id)), "Supplier")
		if err != nil {
			return nil, err
		}
	}
	walletLabel := strconv.Itoa(int(data.Id)) + "@Supplier"

	db := cm.DB.Model(&entity.Distributor{})
	err = db.Where("id = ?", data.Id).Updates(map[string]interface{}{"state": data.State, "wallet": walletLabel}).Find(distributor).Error
	if err != nil {
		return nil, err
	}
	return distributor, nil
}

func (cm *DistributorDao) GetPassing(ctx context.Context) (*[]entity.Distributor, error) {
	distributors := new([]entity.Distributor)
	db := cm.DB.Model(&entity.Distributor{})
	err := db.Where("state = ?", 0).Find(distributors).Error
	if err != nil {
		return nil, err
	}
	return distributors, nil
}
