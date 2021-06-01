package goCron

import (
	"github.com/robfig/cron/v3"
)

func JwtDelCron(){
	MyCron := cron.New(cron.WithSeconds())

	MyCron.AddFunc("0 0/10 * * * *",delJwt)
}

//每10分钟进行一次jwt删除监测
func delJwt(){


}

//执行监测的步骤
func delJwtStep(){

	//1.将数据库中的创建时间设置成变量 @a,取整数

	//2.时间戳计算,当前时间-创建时间 / 60 拿到分钟数

	//3.如果>12分钟,则将数据库中的记录删掉

	//4.删掉变量@a
}