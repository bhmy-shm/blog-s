package dao

import (
	"goweb/internal/model/UserModel"
)


//创建用户表
func (this Dao) CreateUser(user *UserModel.User) error {
	return user.Create(this.db)
}

//注册一条用户信息
func (this Dao) RegisterUser(user *UserModel.User)error{
	return user.Insert(this.db)
}

