package dao

import (
	"context"
	"github.com/google/wire"
	"gorm.io/gorm"
	"medicineApp/internal/entity"
	"medicineApp/internal/schema"
	"time"
)

var MedicineDaoSet = wire.NewSet(wire.Struct(new(MedicineDao), "*"))

type MedicineDao struct {
	DB *gorm.DB
}

func (m *MedicineDao) Update(ctx context.Context, params schema.Medicine) error {
	db := m.DB.Model(&entity.MedicineStatus{})
	db = db.Where("id = ?", params.MedicineId)
	if v := params.MedicineName; v != "" {
		db = db.Updates(map[string]interface{}{"medicineName": v})
	}
	if v := params.ApprovalNo; v != "" {
		db = db.Updates(map[string]interface{}{"approvalNo": v})
	}
	if v := params.Spacification; v != "" {
		db = db.Updates(map[string]interface{}{"spacification": v})
	}
	if v := params.ProduceDate; v != "" {
		tmpTime, err := time.Parse("2006-01-02", v)
		if err != nil {
			return err
		}
		db = db.Updates(map[string]interface{}{"produceDate": tmpTime})
	}
	if v := params.Expiration; v != "" {
		db = db.Updates(map[string]interface{}{"expiration": v})
	}
	if v := params.Producer; v != 0 {
		db = db.Updates(map[string]interface{}{"producer": v})
	}
	if v := params.Status; v != "" {
		db = db.Updates(map[string]interface{}{"status": v})
	}
	if v := params.Num; v != 0 {
		db = db.Updates(map[string]interface{}{"num": v})
	}
	return nil
}

func (m *MedicineDao) Query(ctx context.Context, params schema.QueryMedicine) (*[]entity.MedicineStatus, int, error) {
	db := m.DB.Model(&entity.MedicineStatus{})
	res := new([]entity.MedicineStatus)

	db = db.Joins("Manufacturers")

	if v := params.Name; v != "" {
		db = db.Where("medicineName LIKE ?", "%"+v+"%")
	}
	if v := params.Approval; v != "" {
		db = db.Where("approvalNo LIKE ?", "%"+v+"%")
	}
	if v := params.Producer; v != "" {
		db = db.Where("manufacturers.company_name LIKE ?", "%"+v+"%")
	}
	db = db.Scopes(Paginate(&QueryOption{
		Offset:   params.Offset,
		PageSize: params.PageSize,
	}))

	err := db.Find(res).Error
	if err != nil {
		return nil, 0, err
	}

	count, err := m.GetCount(ctx)
	if err != nil {
		return nil, 0, err
	}

	return res, count, nil
}

func (m *MedicineDao) GetCount(ctx context.Context) (int, error) {
	db := m.DB.Model(&entity.MedicineStatus{})
	res := new([]entity.MedicineStatus)
	var count int64
	err := db.Find(res).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

func (m *MedicineDao) Delete(ctx context.Context, params schema.Medicine) error {
	db := m.DB.Model(&entity.MedicineStatus{})
	err := db.Delete("id = ?", params.MedicineId).Error
	if err != nil {
		return err
	}
	return nil
}

func (m *MedicineDao) Add(ctx context.Context, params schema.Medicine) (*entity.MedicineStatus, error) {
	db := m.DB.Model(&entity.MedicineStatus{})
	tmpTime, err := time.Parse("2006-01-02", params.ProduceDate)
	if err != nil {
		return nil, err
	}
	tmp := &entity.MedicineStatus{
		MedicineName:  params.MedicineName,
		ApprovalNo:    params.ApprovalNo,
		Spacification: params.Spacification,
		ProduceDate:   tmpTime,
		Producer:      uint(params.Producer),
		Status:        "申请",
		BatchNo:       params.BatchNo,
		Expiration:    params.Expiration,
		Num:           params.Num,
	}
	if params.Status != "" {
		tmp.Status = params.Status
	}
	err = db.Create(tmp).Error
	if err != nil {
		return nil, err
	}
	return tmp, nil
}
