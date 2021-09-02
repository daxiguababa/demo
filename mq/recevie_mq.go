package mq

import (
	"errors"
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

type ReceiveMQ struct {
}

func (m ReceiveMQ) Receive() {

	conStr := fmt.Sprintf(RabbitConfig{}.GetConn())

	// 连接RabbitMQ服务器
	conn, err := amqp.Dial(conStr)
	if err != nil {
		err = errors.New("消费者：未找到rabbitmq")
		failOnError(err, "消费者：未找到rabbitmq")
	}
	defer conn.Close()
	// 创建一个channel
	ch, err := conn.Channel()
	if err != nil {
		err = errors.New("消费者：未找到频道")
		failOnError(err, "消费者：未找到频道")
	}
	defer ch.Close()
	// 监听队列
	q, err := ch.QueueDeclare(
		"hello", // 队列名称
		false,   // 是否持久化
		false,   // 是否自动删除
		false,   // 是否独立
		false, nil,
	)
	if err != nil {
		err = errors.New("消费者：打开队列失败")
		failOnError(err, "消费者：打开队列失败")
	}
	// 消费队列
	msg, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		err = errors.New("消费者：处理消息失败")
		failOnError(err, "消费者：处理消息失败")
	}
	// 申明一个goroutine,一遍程序始终监听
	forever := make(chan bool)

	go func() {
		for d := range msg {
			log.Printf("处理消息: %s\n", d.Body)
			fmt.Printf("处理消息: %s\n", d.Body)
		}
	}()
	//log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	log.Println("等待接受新消息......")
	fmt.Println("等待接受新消息......")
	<-forever
}
