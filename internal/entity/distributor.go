package entity

import "gorm.io/gorm"

// Distributor 药品经销商表（distributor）
type Distributor struct {
	gorm.Model
	CompanyName string `gorm:"column:company_name;type:varchar(20);not null"`
	License     string `gorm:"column:licenseID;type:char(20);not null"`
	LegalPerson string `gorm:"column:legal_person;type:char(20);not null"`
	Phone       string `gorm:"column:phone;type:varchar(20);not null"`
	Password    string `gorm:"column:password;type:varchar(20);not null"`
	State       int    `gorm:"column:state;type:int;not null"`
	Ip          string `gorm:"column:Ip;type:varchar(64);not null"`
	GMPID       string `gorm:"column:GMPID;type:varchar(20);not null"`
	Wallet      string `gorm:"column:wallet;type:varchar(255)"`
}
