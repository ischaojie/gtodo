package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"log"
)

type Database struct {
	Self *gorm.DB
}


func (db *Database) Init() {

	// * 初始化数据库
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.dbname"),
		true,
		//"Asia/Shanghai"),
		"Local")
	// * 以config打开mysql数据库
	_db, err := gorm.Open("mysql", config)
	DB = &Database{Self: _db}
	if err != nil {
		log.Printf("Database connection failed. Database name: %s", viper.GetString("db.dbname"))
	}
	// * 解决中文字符问题：Error 1366
	_db = _db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8 auto_increment=1")
	_db.AutoMigrate(&TodoModel{})

}

func (db *Database) Close()  {
	// * 关闭数据库
	db.Self.Close().Error()
}

var DB *Database
