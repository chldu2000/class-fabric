package entity

import "gorm.io/gorm"

// Consumer 药品消费者表（consumer）
type Consumer struct {
	gorm.Model
	Name     string `gorm:"column:name;type:varchar(20);not null"`
	IdNumber string `gorm:"column:id_number;type:char(18);not null"`
	Phone    string `gorm:"column:phone;type:varchar(20)"`
	Password string `gorm:"column:password;type:varchar(20);not null"`
	Wallet   string `gorm:"column:wallet;type:varchar(255)"`
	State    int    `gorm:"column:state;type:int"`
}
