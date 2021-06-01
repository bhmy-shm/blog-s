package service

import (
	"fmt"
	"goweb/internal/Check"
	"goweb/internal/model/JWT"
	"goweb/internal/model/UserModel"
)

//创建表
type CreateRequest struct{}

//新增用户数据
type InsertRequest struct{
	UserPhone string `json:"user_phone" form:"user_phone" binding:"lte=11"`
	UserPass string `json:"user_pass" form:"user_pass" binding:"gte=6"`
	UserEmail string `json:"user_email" form:"user_email" binding:"required"`
	State uint8 `json:"state" form:"state,default=1" binding:"oneof=0 1"`
}

//确认验证码
type AuthRequest struct{
	Auth string `json:"auth" form:"auth" binding:"required"`
}

//用户登录
type LoginRequest struct{
	UserPhone string `json:"user_phone" form:"user_phone" binding:"lte=11"`
	UserPass string `json:"user_pass" form:"user_pass" binding:"gte=6"`
	AppKey string `json:"app_key" form:"app_key" binding:"required"`
	AppSecret string `json:"app_secret" form:"app_secret" binding:"required"`
}



//创建用户表
func (svc *Service) CreatUser() error{
	user := UserModel.User{}
	return svc.dao.CreateUser(&user)
}
//创建jwt表
func (svc *Service) CreatJwt() error{
	jwt := JWT.AuthJwt{}
	return svc.dao.CreateAuth(&jwt)
}


//创建和注册用户表
func (svc *Service) InsertUser(request *InsertRequest) error {
	user := UserModel.User{
		UserPhone: request.UserPhone,
		UserEmail: request.UserEmail,
		UserPass: request.UserPass,
		State: request.State,
	}
	Check.UserRegisterInfo = user
	return svc.check.Check(&user)
}

//接收用户填写的验证码
func (svc *Service) Authentication(request *AuthRequest) error{
	return svc.check.IsAuth(request.Auth)
}

//申请jwt前判断
func(svc *Service) GetAuthJwt(request *LoginRequest) error {
	if svc.check.CheckPassAndUser(request.UserPhone,request.UserPass) == true {
		return svc.dao.ToRedisSecret(request.AppKey,request.AppSecret,request.UserPhone)
	}else {
		return fmt.Errorf("请不要重复申请token")
	}
}

//用户携带jwt登录验证
func (svc *Service) LoginUser(request *LoginRequest) error{
	if svc.check.CheckPass(request.UserPhone,request.UserPass) {
		return nil
	}
	return fmt.Errorf("用户名密码输入错误")
}

//将JWT写入redis
func (svc *Service) InputJwtRedis(token,phone string) error {

	return svc.dao.ToRedisJwt(token,phone)
}