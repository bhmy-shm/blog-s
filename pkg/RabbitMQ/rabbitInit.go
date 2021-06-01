package RabbitMQ

import (
	"fmt"
	"github.com/streadway/amqp"
	"goweb/global"
	"log"
)

var RQConn *amqp.Connection

type MQUser struct {
	Name string
	Pass string
	Host string
	Port uint
}


func NewRmqInit() *amqp.Connection{
	mqs := &MQUser{Name: "shm",
		Pass: "123.com",
		Host: global.RabbitMQSetting.Host,
		Port: 5672}
	//mqs := *global.RabbitMQSetting
	dsn := fmt.Sprintf("amqp://%s:%s@%s:%d",mqs.Name,mqs.Pass,mqs.Host,mqs.Port)

	conn,err := amqp.Dial(dsn)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	RQConn = conn
	return RQConn
}

func GetConn() *amqp.Connection{
	return RQConn
}



