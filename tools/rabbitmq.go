package tools

import (
	"Go-RabbitMQ/config"
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

/**
* @Author: super
* @Date: 2020-08-13 07:57
* @Description:
**/

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel

	//队列名称
	QueueName string
	//交换机
	Exchange string
	//key
	Key string
	//连接信息
	Mqurl string
}

// 创建RabbitMQ实例
func NewRabbitMQ(queuqName string,
	exchange string, key string) *RabbitMQ {
	rabbitmq := &RabbitMQ{
		QueueName: queuqName,
		Exchange:  exchange,
		Key:       key,
		Mqurl:     config.MQURL,
	}
	//创建rabbitmq连接
	var err error
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.FailOnErr(err, "创建连接错误")
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.FailOnErr(err, "获取channel失败")
	return rabbitmq
}

// 断开channel和connection的连接释放资源
func (r *RabbitMQ) Destory() {
	r.channel.Close()
	r.conn.Close()
}

//自定义错误处理函数
func (r *RabbitMQ) FailOnErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%s", message, err)
		panic(fmt.Sprintf("%s:%s", message, err))
	}
}
