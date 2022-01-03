package entity

import "gorm.io/gorm"

// Administrator 系统管理员表（administrator）
type Administrator struct {
	gorm.Model
	SystemAccount string `gorm:"column:system_account;varchar(20);not null"`
	Password      string `gorm:"column:password;type:varchar(20);not null"`
	Wallet        string `gorm:"column:wallet;type:varchar(255)"`
}
