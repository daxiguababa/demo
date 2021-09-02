package mq

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"log"
	"time"
)

type SendMQ struct {
}

func (m SendMQ) Send() (error, string) {
	// 连接RabbitMQ服务器
	conn, err := amqp.Dial(RabbitConfig{}.GetConn())
	if err != nil {
		err = errors.New("生产者：未找到rabbitmq")
		failOnError(err, "生产者：未找到rabbitmq")
	}
	defer conn.Close()
	// 创建一个channel
	ch, err := conn.Channel()
	if err != nil {
		err = errors.New("生产者：未找到频道")
		failOnError(err, "生产者：未找到频道")
	}
	defer ch.Close()

	// 声明一个队列
	q, err := ch.QueueDeclare(
		viper.GetString("rabbitmq.queen"), // 队列名称
		false,                             // 是否持久化
		false,                             // 是否自动删除
		false,                             // 是否独立
		false, nil,
	)
	if err != nil {
		err = errors.New("生产者：打开队列失败")
		failOnError(err, "生产者：打开队列失败")
	}
	// 发送消息到队列中
	msg := fmt.Sprintf("当前时间：%d", time.Now().UnixNano())
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
	if err != nil {
		err = errors.New("生产者：发布消息失败")
		failOnError(err, "生产者：发布消息失败")
	}
	return err, msg
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
