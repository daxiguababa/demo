#  go-demo #
# 本项目使用gin框架集成rabbitmq、gorm、redis、mongo、viper配置读取等，后续将加入微信开发应用 #

## 注意：使用之前请先配置好conf/config.json文件 ##

## 1.框架 ##
go get -u github.com/gin-gonic/gin

## 2.redis ##
go get -u github.com/go-redis/redis

## 3.rabbitmq 驱动 ##
go get github.com/streadway/amqp

## 4.配置获取 ##
go get -u github.com/spf13/viper

## 5.mongo ##
go get go.mongodb.org/mongo-driver/mongo

## 6.mysql ##
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
