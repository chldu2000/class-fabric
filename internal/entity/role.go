package entity

import (
	"gorm.io/gorm"
)

// Role 角色信息表
type Role struct {
	gorm.Model
	Name string `gorm:"type:varchar(255);not null"`
}
