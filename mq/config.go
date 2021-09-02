package mq

import (
	"fmt"
	"github.com/spf13/viper"
)

type RabbitConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Vhost    string
}

func (m RabbitConfig) GetConn() string {
	m.Host = viper.GetString("rabbitmq.host")
	m.Port = viper.GetString("rabbitmq.port")
	m.User = viper.GetString("rabbitmq.user")
	m.Password = viper.GetString("rabbitmq.password")
	m.Vhost = viper.GetString("rabbitmq.vhost")

	conStr := fmt.Sprintf("amqp://%s:%s@%s:%s/%s", m.User, m.Password, m.Host, m.Port, m.Vhost)
	return conStr
}
