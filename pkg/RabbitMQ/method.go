package RabbitMQ

import (
	"fmt"
	"github.com/streadway/amqp"
	"goweb/global"
	"log"
	"strings"
)

type RegiterMQ interface{

}

type MQ struct {
	Channel *amqp.Channel	//AMQP全局消息通道
	notifyConfirm chan amqp.Confirmation
	notifyReturn chan amqp.Return
}

//创建消息通道
func NewMQ() *MQ {
	c,err := NewRmqInit().Channel()
	if err != nil {
		log.Println("xxx",err)
	}
	return &MQ{Channel: c}
}

//生成交换器，路由键
func MQExchangeInit()error{
	MQ := NewMQ()
	if MQ == nil {
		return fmt.Errorf("mq init is nil ")
	}
	defer MQ.Channel.Close()

	//创建交换器
	err := MQ.Channel.ExchangeDeclare(global.RabbitMQSetting.Exchange,"direct", false, false, false, false, nil)
	if err != nil {
		return fmt.Errorf("Exchange Error :%v", err)
	}
	//根据交换器生产队列
	err = MQ.decQueueAndBind(global.RabbitMQSetting.Queue,	//队列
		global.RabbitMQSetting.Router,		//路由键
		global.RabbitMQSetting.Exchange,	//交换器
	)
	if err != nil {
		return fmt.Errorf("Queue Bind error:%v", err)
	}
	return nil
}

//发送消息队列
func (this *MQ) SendMessage(exchange string, key string, message string) error {

	//发送队列到交换器
	err := this.Channel.Publish(exchange, key, true, false, amqp.Publishing{
		ContentType: "text/palin",
		Body:        []byte(message),
	})
	return err
}

//循环生成队列
func (this *MQ) decQueueAndBind(queues string, key string, exchange string) error {
	//分隔队列名称
	qList := strings.Split(queues, ",")
	//1.循环创建多个队列
	for _, queue := range qList {
		q, err := this.Channel.QueueDeclare(queue, false, false, false, false, nil)
		if err != nil {
			return err
		}

		//2.每创建一个队列就绑定一个路由键
		err = this.Channel.QueueBind(q.Name, key, exchange, false, nil)
		if err != nil {
			return err
		}
	}
	return nil
}

//消费者消费队列
func (this *MQ) Counsumer(queue string, key string, callback func(<-chan amqp.Delivery,string)) error {
	msgs, err := this.Channel.Consume(queue, key, false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}
	callback(msgs,key)
	return err
}



