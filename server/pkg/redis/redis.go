/*
* package redis 是对redis的封装，提供redis的初始化
 */

package redis

import (
	"fmt"
	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis"
	"github.com/shiniao/gtodo/pkg/log"
	"github.com/spf13/viper"
)

// Nil redis 返回为空
const Nil = redis.Nil

var Client *redis.Client

// Init 初始化Redis客户端
func Init() {
	Client = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
		PoolSize: viper.GetInt("redis.pool_size"), // 连接池大小
	})
	fmt.Println("redis addr:", viper.GetString("redis.addr"))
	_, err := Client.Ping().Result()
	if err != nil {
		log.Errorf("[redis] redis ping err: %+v", err)
		panic(err)
	}
}

// InitTestRedis 是测试redis的辅助函数
func InitTestRedis() {
	mr, err := miniredis.Run()

	if err != nil {
		panic(err)
	}

	Client = redis.NewClient(&redis.Options{Addr: mr.Addr()})
	fmt.Println("mini redis addr: ", mr.Addr())
}
