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

// ConsumerDaoSet 注入 DI
var ConsumerDaoSet = wire.NewSet(wire.Struct(new(ConsumerDao), "*"))

// ConsumerDao Consumer 表相关的数据库操作
type ConsumerDao struct {
	DB *gorm.DB
}

// ListRes 需要返回给前端的数据
type ListRes struct {
	Name     string `json:"name"`
	IdNumber string `json:"id_number"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func (con *ConsumerDao) Test(ctx context.Context) (*[]ListRes, error) {
	result := new([]entity.Consumer)
	db := con.DB.Model(&entity.Consumer{})
	res := new([]ListRes)
	err := db.Find(result).Error
	err = copier.Copy(&res, &result)
	if err != nil {
		return nil, err
	}
	return res, err
}

func (cm *ConsumerDao) GetByName(ctx context.Context, name string) (*entity.Consumer, error) {
	consumer := new(entity.Consumer)
	db := cm.DB.Model(&entity.Consumer{})
	err := db.Where("name = ?", name).First(consumer).Error
	if err != nil {
		return nil, err
	}
	return consumer, nil
}

func (cm *ConsumerDao) GetPassing(ctx context.Context) (*[]entity.Consumer, error) {
	consumers := new([]entity.Consumer)
	db := cm.DB.Model(&entity.Consumer{})
	err := db.Where("state = ?", 0).Find(consumers).Error
	if err != nil {
		return nil, err
	}
	return consumers, nil
}

func (cm *ConsumerDao) Query(ctx context.Context, IdNumber string) (*entity.Consumer, error) {
	consumer := new(entity.Consumer)
	db := cm.DB.Model(&entity.Consumer{})
	err := db.Where("id_number = ?", IdNumber).First(consumer).Error
	if err != nil {
		return nil, err
	}
	return consumer, nil
}

func (cm *ConsumerDao) Pass(ctx context.Context, data schema.PassReqBodySchema) (*entity.Consumer, error) {
	consumer := new(entity.Consumer)
	db := cm.DB.Model(&entity.Consumer{})
	//wallet := "测试用钱包"//这里留有接口

	// 这里 data 最好包括用户的密码作为链上的密码
	err := fabricUtils.RegisterAndEnroll(strconv.Itoa(int(data.Id)), "Consumer")
	if err != nil {
		return nil, err
	}
	wallet, err := gateway.NewFileSystemWallet("wallet")
	if err != nil {
		return nil, err
	}
	if !wallet.Exists(strconv.Itoa(int(data.Id)) + "@Consumer") {
		err := fabricUtils.PopulateWallet(wallet, strconv.Itoa(int(data.Id)), "Consumer")
		if err != nil {
			return nil, err
		}
	}
	walletLabel := strconv.Itoa(int(data.Id)) + "@Consumer"

	err = db.Where("id = ?", data.Id).Updates(map[string]interface{}{"state": data.State, "wallet": walletLabel}).Find(consumer).Error
	if err != nil {
		return nil, err
	}
	return consumer, nil
}

func (cm *ConsumerDao) Register(ctx context.Context, data schema.RegisterConsumerReqBodySchema) (*entity.Consumer, error) {
	consumer := new(entity.Consumer)
	db := cm.DB.Model(&entity.Consumer{})
	err := copier.Copy(consumer, data)
	if err != nil {
		return nil, err
	}
	err = db.Save(consumer).Error
	if err != nil {
		return nil, err
	}
	return consumer, nil
}
