package JWT

import (
	"context"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"goweb/internal/model"
)

//数据库结构体
type AuthJwt struct{
	UserPhone string `json:"user_phone" form:"user_phone" gorm:"type:varchar(20)"`
	AppKey string `json:"app_key" form:"app_key" gorm:"type:varchar(50)"`
	AppSecret string `json:"app_secret" form:"app_secret" gorm:"type:varchar(50)"`
	*model.Model
}



func(this AuthJwt) tableName() string{
	return "blog_auth"
}

//创建该表
func (this AuthJwt) Create(db *gorm.DB) error {
	return db.Table(this.tableName()).AutoMigrate(&this)
}

//将新用户的jwt密文文件入库
func (this AuthJwt) Insert(db *gorm.DB,jwt *AuthJwt) error {
	return db.Table(this.tableName()).Create(&jwt).Error
}

//比对数据库中的token键和密钥
func (a AuthJwt) Get(db *gorm.DB) (AuthJwt,error){
	var auth AuthJwt
	db = db.Where("app_key = ? AND app_secret = ? AND is_del = ?",a.AppKey,
		a.AppSecret,0)
	err := db.First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		return auth,err
	}
	return auth,nil
}

//Redis密文入库 hash
func (this AuthJwt) JwtHmset(redis *redis.Client,jwt *AuthJwt) error {
	//Hmset写入jwt的hashmap
	datas := map[string]interface{}{
		"key":jwt.AppKey,
		"secret":jwt.AppSecret,
	}
	return redis.HMSet(context.TODO(),jwt.UserPhone,datas).Err()
}
