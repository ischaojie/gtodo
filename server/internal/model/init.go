package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/shiniao/gtodo/pkg/log"
	"github.com/spf13/viper"
	"time"
)

// Database 定义数据库
type Database struct {
	Default *gorm.DB
	Docker *gorm.DB
}

// DB 全局数据库变量
var DB *Database

// Init 初始化数据库
func (db *Database) Init() {
	DB = &Database{
		Default: getDefaultDB(),
		Docker:  getDockerDB(),
	}
}

// Close 关闭数据库
func (db *Database) Close() {
	err := DB.Default.Close()
	if err != nil {
		log.Warnf("[model] close default db err: %+v", err)
	}
	err = DB.Docker.Close()
	if err != nil {
		log.Warnf("[model] close docker db err: %+v", err)
	}
}

// getDB 返回默认的数据库实例
func getDB() *gorm.DB {
	return DB.Default
}

// getDefaultDB 获取默认数据库配置
func getDefaultDB() *gorm.DB {
	return initDefaultDB()
}

// initDefaultDB 初始化默认数据库
func initDefaultDB() *gorm.DB {
	return openDB(viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.name"))
}

func getDockerDB() *gorm.DB {
	return initDockerDB()
}

func initDockerDB() *gorm.DB {
	return openDB(viper.GetString("docker_db.username"),
		viper.GetString("docker_db.password"),
		viper.GetString("docker_db.addr"),
		viper.GetString("docker_db.name"))
}

// openDB 生成数据库实例
func openDB(username, password, addr, name string) *gorm.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		// "Asia/Shanghai"),
		"Local")

	// * 以config打开mysql数据库
	db, err := gorm.Open("mysql", config)
	if err != nil {
		log.Errorf("Database connection failed. Database name: %s, err: %+v", name, err)
	}
	// * 解决中文字符问题：Error 1366
	db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8 auto_increment=1")

	// 配置数据库
	setupDB(db)

	return db

}

// setupDB 配置数据库参数
func setupDB(db *gorm.DB) {
	db.LogMode(viper.GetBool("gorm.show_log"))
	// 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	db.DB().SetMaxOpenConns(viper.GetInt("gorm.max_open_conn"))
	// 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
	db.DB().SetMaxIdleConns(viper.GetInt("gorm.max_idle_conn"))
	db.DB().SetConnMaxLifetime(time.Minute * viper.GetDuration("gorm.conn_max_lift_time"))
}
