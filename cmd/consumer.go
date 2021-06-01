package cmd

import (
	"fmt"
	"github.com/streadway/amqp"
	"goweb/global"
	"goweb/pkg/RabbitMQ"
	"math/rand"
	"strings"
	"time"
)

func NewAuthConsumer(c string) error {

	//从管道中消费数据
	MQ := RabbitMQ.NewMQ()

	//设置消费者限流
	//MQ.Channel.Qos(10,0,false)

	//===============将 c 这个传入的参数，作为counsumer的key传入==================
	err := MQ.Counsumer(global.RabbitMQSetting.Queue, c, SendAuth)
	if err != nil {
		return err
	}
	defer MQ.Channel.Close()

	return nil
}

//发送验证码
func SendAuth(msgs <-chan amqp.Delivery,consumer string){
	auth := genValidateCode(6)

	for msg := range msgs{
		fmt.Printf("%s向phone=%s的用户发送短信,验证码:%s\n", consumer, string(msg.Body),auth)

		//同时向Gmap中写入验证码作为临时记录
		expireMap(time.Second*60,string(msg.Body),auth)

		msg.Ack(false)
	}
}


//随机验证码
func genValidateCode(width int) string {
	numeric := [10]byte{0,1,2,3,4,5,6,7,8,9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())
	var num strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&num, "%d", numeric[rand.Intn(r)])
	}
	return num.String()
}

//超时机制的map
func expireMap(expire time.Duration,key,value string){
	global.GMap.Store(key,value)

	//如果写入成功
	if _,ok := global.GMap.Load(key) ; ok {}

	//距离写入时间55秒秒种后将其删除,清空GMap里面的指定数据
	time.AfterFunc(expire, func() {
		global.GMap.Delete(key)
	})
}



