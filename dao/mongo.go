package dao

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var Mongo *mongo.Database

type MonCon struct {
}

type MonConData struct {
	Name string `json:"name" bson:"name" form:"name"`
}

func InitMongo() {
	//连接MongoDB
	mongoConfig := MongoConfig{
		Dsn:       viper.GetString("mongodb.dsn"),
		Host:      viper.GetString("mongodb.host"),
		Port:      viper.GetInt("mongodb.port"),
		DataBase:  viper.GetString("mongodb.database"),
		User:      viper.GetString("mongodb.user"),
		Password:  viper.GetString("mongodb.password"),
		DB:        viper.GetString("mongodb.db"),
		PoolLimit: viper.GetUint64("mongodb.poolLimit"),
	}
	Mongo = mongoConfig.init()
}

type MongoConfig struct {
	Dsn       string
	Host      string //主机名
	Port      int    //端口
	User      string //用户名
	Password  string //密码
	DataBase  string //认证数据库
	DB        string //数据库
	PoolLimit uint64 //连接池
}

//初始化MongoDB
func (c *MongoConfig) init() *mongo.Database {
	connString := c.Dsn
	if connString == "" {
		connString = fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", c.User, c.Password, c.Host, c.Port, c.DataBase)
		//connString = fmt.Sprintf("mongodb://%s:%s@%s:%d", c.User, c.Password, c.Host, c.Port)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI(connString)
	clientOptions.SetMaxPoolSize(c.PoolLimit)
	clientOptions.SetMinPoolSize(1)
	clientOptions.SetMaxConnIdleTime(3 * time.Second)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		logrus.Fatalln("MongoDb连接失败", err.Error())
		return nil
	}
	logrus.Println("MongoDb连接成功")
	return client.Database(c.DB)
	//return client
}

func (c MonCon) Con() context.Context {
	return context.TODO()
}
