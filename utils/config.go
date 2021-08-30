package utils

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Name string
}

//初始化配置
func InitConfig(fileName string) error {

	c := Config{
		Name: fileName,
	}
	// 初始化配置文件
	if err := c.setConfig(); err != nil {
		return err
	}
	c.watchConfig()
	return nil
}

func (c *Config) setConfig() error {
	viper.AddConfigPath("./conf/")
	if c.Name != "" {
		// 如果指定了配置文件，则解析指定的配置文件
		viper.SetConfigName(c.Name)
	} else {
		// 如果没有指定配置文件，则解析默认的配置文件
		viper.SetConfigName("config")
	}
	// 设置配置文件格式为YAML
	viper.SetConfigType("json")
	// viper解析配置文件
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

// 监听配置文件是否改变,用于热更新
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Fatalf("配置文件修改更新: %s\n", e.Name)
	})
}
