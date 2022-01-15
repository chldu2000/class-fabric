package entity

import (
	"time"
)

type Manufacturers struct {
	ID          uint      `gorm:"column:id;primarykey;auto_increment;type:int;"` // 生产商ID
	RegisterID  string    `gorm:"column:registerID"`                             // 企业注册号
	CompanyName string    `gorm:"column:company_name"`                           // 企业名称
	Password    string    `gorm:"column:password"`                               // 密码
	LegalPerson string    `gorm:"column:legal_person"`                           // 法人姓名
	Phone       string    `gorm:"column:phone"`                                  // 法人联系方式
	LicenseID   string    `gorm:"column:licenseID"`                              // 经营许可证编号
	GMPID       string    `gorm:"column:GMPID"`                                  // 药品GMP证书编号
	State       int       `gorm:"column:state"`                                  // 审核状态 0 1 2 三个状态 对应 注册完成 审核通过 区块链地址分配 三个阶段
	Ip          string    `gorm:"column:Ip"`                                     // Ip
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
	DeletedAt   time.Time `gorm:"column:deleted_at"`
	Wallet      string    `gorm:"column:wallet"`
}

func (Manufacturers) TableName() string {
	return "manufacturers"
}
