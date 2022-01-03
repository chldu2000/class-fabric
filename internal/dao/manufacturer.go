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

// ManufacturerDaoSet 注入 DI
var ManufacturerDaoSet = wire.NewSet(wire.Struct(new(ManufacturerDao), "*"))

// ManufacturerDao Manufacturer 表相关的数据库操作
type ManufacturerDao struct {
	DB *gorm.DB
}

func (mf *ManufacturerDao) GetByComp(ctx context.Context, name string) (*entity.Manufacturer, error) {
	comp := new(entity.Manufacturer)
	db := mf.DB.Model(&entity.Manufacturer{})
	err := db.Where("company_name = ?", name).First(comp).Error
	if err != nil {
		return nil, err
	}
	return comp, nil
}

func (mf *ManufacturerDao) Query(ctx context.Context, name string) (*entity.Manufacturer, error) {
	comp := new(entity.Manufacturer)
	db := mf.DB.Model(&entity.Manufacturer{})
	err := db.Where("company_name = ?", name).First(comp).Error
	if err != nil {
		return nil, err
	}
	return comp, nil
}

func (cm *ManufacturerDao) Register(ctx context.Context, data schema.RegisterManufacturerReqBodySchema) (*entity.Manufacturer, error) {
	manufacturer := new(entity.Manufacturer)
	db := cm.DB.Model(&entity.Manufacturer{})
	err := copier.Copy(manufacturer, data)
	if err != nil {
		return nil, err
	}
	err = db.Save(manufacturer).Error
	if err != nil {
		return nil, err
	}
	return manufacturer, nil
}

func (cm *ManufacturerDao) Pass(ctx context.Context, data schema.PassReqBodySchema) (*entity.Manufacturer, error) {
	manufacturer := new(entity.Manufacturer)
	//wallet := "测试用钱包"//这里留有接口

	err := fabricUtils.RegisterAndEnroll(strconv.Itoa(int(data.Id)), "Manufacturer")
	if err != nil {
		return nil, err
	}
	wallet, err := gateway.NewFileSystemWallet("wallet")
	if err != nil {
		return nil, err
	}
	if !wallet.Exists(strconv.Itoa(int(data.Id)) + "@Manufacturer") {
		err := fabricUtils.PopulateWallet(wallet, strconv.Itoa(int(data.Id)), "Manufacturer")
		if err != nil {
			return nil, err
		}
	}
	walletLabel := strconv.Itoa(int(data.Id)) + "@Manufacturer"

	db := cm.DB.Model(&entity.Manufacturer{})
	err = db.Where("id = ?", data.Id).Updates(map[string]interface{}{"state": data.State, "wallet": walletLabel}).Find(manufacturer).Error
	if err != nil {
		return nil, err
	}
	return manufacturer, nil
}

func (cm *ManufacturerDao) GetPassing(ctx context.Context) (*[]entity.Manufacturer, error) {
	manufacturers := new([]entity.Manufacturer)
	db := cm.DB.Model(&entity.Manufacturer{})
	err := db.Where("state = ?", 0).Find(manufacturers).Error
	if err != nil {
		return nil, err
	}
	return manufacturers, nil
}
