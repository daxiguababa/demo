package dao

import (
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var Redis *redis.Client

//初始化Redis
func InitRedis() {
	//连接Redis
	redisConfig := redisConfig{
		Host:     viper.GetString("redis.host"),
		Port:     viper.GetString("redis.port"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.DB"),
		PoolSize: viper.GetInt("redis.pool_size"),
	}
	Redis = redisConfig.init()
	if _, err := Redis.Ping().Result(); err != nil {
		logrus.Fatalln("redis连接出错:", err)
		return
	}
	logrus.Println("redis连接成功")
}

type redisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
	PoolSize int
}

//初始化
func (c *redisConfig) init() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     c.Host + ":" + c.Port,
		Password: c.Password,
		DB:       c.DB,
		PoolSize: c.PoolSize,
	})
	return redisClient
}
