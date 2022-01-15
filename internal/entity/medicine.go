package entity

import "time"

type MedicineStatus struct {
	ID            uint          `gorm:"column:id;type:int;primary key;auto_increment"`
	MedicineName  string        `gorm:"column:medicineName"`
	ApprovalNo    string        `gorm:"column:approvalNo"`
	Spacification string        `gorm:"column:spacification"`
	ProduceDate   time.Time     `gorm:"column:produceDate"`
	Producer      uint          `gorm:"column:producer;type:int"`
	Manufacturers Manufacturers `gorm:"foreignKey:Producer;reference:ID"`
	Status        string        `gorm:"column:status"`
	BatchNo       string        `gorm:"column:batchNo"`
	Expiration    string        `gorm:"column:expiration"`
	Num           int           `gorm:"num"`
}

func (MedicineStatus) TableName() string {
	return "medicinestatus"
}
