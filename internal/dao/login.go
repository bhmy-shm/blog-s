package dao

import (
	"context"
	"goweb/global"
	"goweb/internal/model/JWT"
	"time"
)

//mysql创建数据表
func (this Dao) CreateAuth(jwt *JWT.AuthJwt) error {
	return jwt.Create(this.db)
}

//向数据库写入一个jwt的 令牌key(密钥) 和 serect()
func (this Dao) ToDBJwt(appkey,appSecret,phone string) error {
	var jwt = JWT.AuthJwt{
		AppSecret: appSecret,
		AppKey: appkey,
		UserPhone: phone,
	}
	return jwt.Insert(this.db,&jwt)
}

//向Redis写入jwt的密文键值对
func (this Dao) ToRedisSecret(appkey,appSecret,phone string) error {
	var jwt = JWT.AuthJwt{
		AppSecret: appSecret,
		AppKey: appkey,
		UserPhone: phone,
	}
	return jwt.JwtHmset(this.redis,&jwt)
}

//向Redis的 hash表 追加写入jwt的token编码，并根据jwt的超时时间为这个hash表设置超时时间
func (this Dao) ToRedisJwt(token,phone string) error {
	data := map[string]string{
		"token":token,
	}
	err := this.redis.HMSet(context.TODO(),phone,data).Err()
	if err != nil{
		return err
	}
	err = this.redis.Expire(context.TODO(),phone,global.JWTSetting.Expire*time.Second).Err()
	if err != nil{
		return err
	}
	return nil
}