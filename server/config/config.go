package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"strings"
)

type Config struct {
	Name string
}

// * 初始化配置文件
func Init(cfg string) error {
	c := Config{Name:cfg}
	if err := c.initConfig(); err != nil {
		return err
	}
	// * 热更新
	c.WatchConfig()
	return nil
}

func (c *Config) initConfig() error {
	// * 指定配置文件
	if c.Name != "" {
		// * 设置配置文件名
		viper.SetConfigFile(c.Name)
	}else {
		// * 未指定配置文件，读取默认路径
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
	}
	// * 设置config文件类型
	viper.SetConfigType("yaml")

	// TODO 待理解
	// * 读取环境变量
	viper.AutomaticEnv()
	// * 设置环境变量前缀
	viper.SetEnvPrefix("api")
	replacer := strings.NewReplacer(",", "_")
	viper.SetEnvKeyReplacer(replacer)

	// * 读取配置文件内容，使用viper.get获取配置
	if err:=viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

// * 监控配置文件变化并热加载
func (c *Config) WatchConfig()  {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed: %s", e.Name)
	})
}


