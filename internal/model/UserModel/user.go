package UserModel

import (
	"gorm.io/gorm"
	"goweb/internal/model"
)

type User struct{
	*model.Model
	State uint8 `json:"state"`
	UserPhone string `json:"user_phone" gorm:"type:varchar(15);not null"`
	UserPass  string `json:"user_pass" gorm:"type:varchar(20);not null"`
	UserEmail string `json:"user_email" gorm:"type:varchar(20);not null"`
}

//用户表名
func(this User) tableName() string{
	return "users"
}

//创建用户表
func (this User) Create(db *gorm.DB) error {
	return db.Table(this.tableName()).AutoMigrate(&this)
}

//写入注册信息
func (this User) Insert(db *gorm.DB) error{
	return db.Table("users").Create(&this).Error
}
