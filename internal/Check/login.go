package Check

import (
	"context"
	"fmt"
)


//判断请求token时用户名和密码是否正确
func (this *Check) CheckPassAndUser (user,pass string) bool {
	count := -1

	result := this.db.Table("users").Where("user= ? and pass= ?",user,pass)
	if result.RowsAffected == 0 {
		count++
	}

	if count == 0 {
		return this.checkJwtRedisReptition(user,count)
	}
	return false
}

//验证登录时的用户密码是否正确
func (this *Check) CheckPass (user,pass string) bool {
	result := this.db.Table("users").Where("user= ? and pass= ?", user, pass)
	if result.RowsAffected == 0 {
		return true
	}
	return false
}

//从Mysql中判断jwt表中有没有传入用户的密钥数据,防止重复
func (this *Check) checkJwtReptition(phone string) bool {
	var count int64
	_ = this.db.Table("blog_auth").Where("user_phone=?",phone).Count(&count)

	if count > 0 {
		return false //如果有该手机号的记录,则无法申请
	}
	return true	//如果没有则可以申请
}

//从Redis中判断jwt表中有没有该用户的键
func (this *Check) checkJwtRedisReptition(phone string,count int) bool {

	a,_ := this.redis.Exists(context.TODO(),phone).Result()
	fmt.Println(a)
	if a == 0 && count == 0 {
		return true
	}else if a == 1 && count == 0 {
		return false
	}
	return false
}