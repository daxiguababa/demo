package dao

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"io"
	"log"
	"math"
	"os"
	"time"
)

//Mysql连接处理

var DB *gorm.DB

//初始化Mysql数据库
func InitDB() {
	dbConfig := dbConfig{
		Host:        viper.GetString("mysql.host"),
		Port:        viper.GetString("mysql.port"),
		Database:    viper.GetString("mysql.database"),
		User:        viper.GetString("mysql.user"),
		Password:    viper.GetString("mysql.password"),
		TablePrefix: viper.GetString("mysql.prefix"),
		Charset:     viper.GetString("mysql.charset"),
		MaxIdleConn: viper.GetInt("mysql.MaxIdleConns"),
		MaxOpenConn: viper.GetInt("mysql.MaxOpenConns"),
	}
	DB = dbConfig.init()
}

type dbConfig struct {
	Host        string //主机名
	Port        string //端口
	Database    string //数据库
	User        string //用户名
	Password    string //密码
	TablePrefix string //表前缀
	Charset     string //字符集
	MaxIdleConn int    //空闲时最大的连接数
	MaxOpenConn int    //最大的连接数
}

//初始化连接
func (c *dbConfig) init() *gorm.DB {
	file, err := os.OpenFile(fmt.Sprintf("%s/live-api-sql-%s.log", viper.GetString("log.log_path"), time.Now().Format("2006-01-02")), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("打开日志文件失败：", err)
	}
	newLogger := logger.New(
		log.New(io.MultiWriter(os.Stdout, file), "\r\n", log.Ldate|log.Ltime|log.Lshortfile), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      false,       // 禁用彩色打印
		},
	)
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=%s", c.User, c.Password, c.Host, c.Port, c.Database, c.Charset, "Local")
	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   c.TablePrefix,
			SingularTable: true,
		},
		Logger: newLogger,
		//SkipDefaultTransaction: true,
	})
	if err != nil {
		logrus.Fatalln("mysql初始化错误：", err.Error())
		return nil
	}
	sqlDB, err := db.DB()
	if err != nil {
		logrus.Fatalln("连接数据库失败", err.Error())
		return nil
	}
	//空闲时最大的连接数
	sqlDB.SetMaxIdleConns(c.MaxIdleConn)

	//最大的连接数
	sqlDB.SetMaxOpenConns(c.MaxOpenConn)
	logrus.Println("连接数据库成功")
	return db
}

//分页数据
func Pagination(table string, where gin.H, page int, pageSize int) gin.H {
	var count int64
	var total int64
	db := DB.Table(table).Count(&total)
	if len(where) != 0 {
		db.Where(where).Limit(pageSize).Offset(page).Count(&count)
	} else {
		count = total
	}
	return gin.H{
		"pagination": gin.H{
			"total":        total,                                                //总数
			"count":        count,                                                // 当前页的数量
			"per_page":     pageSize,                                             //数量（页）
			"current_page": page,                                                 //当前多少页
			"last_page":    int64(math.Ceil(float64(total) / float64(pageSize))), //最后一页
		},
	}
}
