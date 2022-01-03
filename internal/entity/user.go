package entity

import (
	"gorm.io/gorm"
)

// User users表模型, 模型采用 gorm 自动迁移生成, 不推荐直接通过 sql 修改表结构
// 各字段含义如下:
//   UID: 内部 id, 充当用户名
//   Password: 用户密码, 前端 md5 一次, 后端 md5 一次, 即存储 md5(md5(password))
//   Status: 用户状态 (暂时不知道作用)
//   Authority: 用户权限, 采用三位二进制判断, 从高到低依次为 leader, hr, admin
//   Name: 姓名
//   Sex: 性别
//   Stuid: 学号
//   College: 学院
//   Major: 专业
//   Class: 班级
//   Dormitory: 寝室
//   Phone: 电话号码
//   QQ: qq 号码
//   Wechat: 微信号
//   Email: 邮箱
//   Grade: 年级
//   Role: 用户角色, 与角色表对应
type User struct {
	gorm.Model
	UID       string `gorm:"column:uid;type=varchar(255);not null"`
	Password  string `gorm:"type:varchar(255);not null"`
	Status    string `gorm:"type:varchar(255);default:'0';not null"`
	Authority int    `gorm:"type:int;default:0;not null"`

	Name      string `gorm:"type=varchar(255);not null"`
	Sex       int    `gorm:"type=tinyint;not null"`
	Stuid     string `gorm:"type=varchar(255);not null"`
	College   string `gorm:"type=varchar(255);not null"`
	Major     string `gorm:"type=varchar(255);not null"`
	Class     string `gorm:"type=varchar(255);not null"`
	Dormitory string `gorm:"type=varchar(255);not null"`
	Phone     string `gorm:"type=varchar(255);not null"`
	QQ        string `gorm:"column:qq;type=varchar(255);not null"`
	Wechat    string `gorm:"type=varchar(255);not null"`
	Email     string `gorm:"type=varchar(255);not null"`
	Grade     string `gorm:"type=varchar(255);not null"`

	RoleID int `gorm:"not null"`
	Role   Role
}
