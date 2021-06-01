package Check

import (
	"fmt"
	"goweb/global"
	"goweb/internal/model/UserModel"
	"goweb/pkg/RabbitMQ"
	"goweb/pkg/errcode"
	"regexp"
)

var UserRegisterInfo UserModel.User

func (this *Check) Check(user *UserModel.User) (err error) {

	//如果手机号重复返回错误
	if err = this.isTelephoneExist(user.UserPhone) ; err != nil {
		return err
	}

	//如果有些格式不正确返回错误
	if err = this.isEmail(user.UserEmail) ; err != nil {
		return err
	}

	return errcode.ErrorUserRedirect
}

//验证手机号是否重复
func (this *Check) isTelephoneExist(phone string) error {
	var user UserModel.User

	//如果查询的结果为空，代表没有重复，如果不为空，代表这个手机号已经重复了
	Rows := this.db.Where("user_phone = ?",phone).Find(&user).RowsAffected
	if Rows != 0 {
		return errcode.ErrorUserRegisterRepetition
	}

	//如果手机号为空，则向消息队列发送该手机号
	err := this.SendMQ(phone)
	return err
}

//验证邮箱号
func (this *Check) isEmail(email string) error {

	reEmail := `[\w\.]+@\w+\.[a-z]{2,3}(\.[a-z]{2,3})?`
	re,_ := regexp.MatchString(reEmail,email)
	if re==true{
		return nil
	}else{
		return fmt.Errorf("邮箱错误,请重新填写")
	}
}

//验证验证码是否正确
func (this *Check) IsAuth(auth string) error {
	var k bool
	var mapkey string
	//如果传入的验证码存在于Gmap中则通过注册，然后入库
	global.GMap.Range(func(key, value interface{}) bool {
		mapkey = key.(string)
		if auth == value {
			k = true
			return k
		}
		 k = false
		 return k
	})

	if k == true {
		this.dao.RegisterUser(&UserRegisterInfo)	//将用户的注册信息入库
		global.GMap.Delete(mapkey)	//入库的同时将验证码从Gmap中删掉
		return nil
	}else{
		return fmt.Errorf("验证码错误，请重新发送")
	}
}

//发送消息队列
func (this *Check) SendMQ(phone string) error {

	MQ := RabbitMQ.NewMQ()
	//向接收注册请求的消息队列发送数据
	err := MQ.SendMessage(global.RabbitMQSetting.Exchange,
		global.RabbitMQSetting.Router,
		phone,
	)
	if err == nil {
		return nil
	}else {
		//如果发送消息失败，则向队列进行重发
		err = MQ.SendMessage(global.RabbitMQSetting.Exchange,
			global.RabbitMQSetting.Router,
			phone,
		)
	}
	return err
}
