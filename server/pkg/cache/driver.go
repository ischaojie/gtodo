/*
package cache 封装了不同的缓存client，提供统一的Driver接口，
可以根据需要选择不同的client
*/
package cache

import (
	"time"
)

// Driver 定义cache接口，各cache可以实现该接口
type Driver interface {
	Set(key string, val interface{}, exp time.Duration) error
	Get(key string, val interface{}) error
	MultiSet(ValMap map[string]interface{}, exp time.Duration) error
	MultiGet(keys []string, val interface{}) error
	Del(keys ...string) error
	Incr(key string, step int64) (int64, error)
	Decr(key string, step int64) (int64, error)
}

/*对外提供的cache方法接口*/

// Client 代表缓存客户端
var Client Driver

// Set 数据
func Set(key string, val interface{}, exp time.Duration) error {
	return Client.Set(key, val, exp)
}

// Get 数据
func Get(key string, val interface{}) error {
	return Client.Get(key, val)
}

// MultiSet 批量设置数据
func MultiSet(ValMap map[string]interface{}, exp time.Duration) error {
	return Client.MultiSet(ValMap, exp)
}

// MultiSet 批量获取数据
func MultiGet(keys []string, val interface{}) error {
	return Client.MultiGet(keys, val)
}

// Del 批量删除
func Del(keys ...string) error {
	return Client.Del(keys...)
}

// Incr 自增
func Incr(key string, step int64) (int64, error) {
	return Client.Incr(key, step)
}

// Decr 自减
func Decr(key string, step int64) (int64, error) {
	return Client.Decr(key, step)
}

/*初始化缓存*/

const (
	redisCacheDriver = "redis"
	memCacheDriver   = "memory"
)

// Init 初始化Driver，根据配置选择不同的driver
// func Init() {
// 	cacheDriver := "redis"
// 	cachePrefix := "gtodo"
// 	fmt.Println("get prefixkey1: ", cachePrefix)
// 	encoding := JSONEncoding{}
//
// 	switch cacheDriver {
// 	case redisCacheDriver:
// 		redis.Init()
// 		Client = NewRedisCache(redis.Client, cachePrefix, encoding)
// 	case memCacheDriver:
// 		Client = NewMemoryCache(cachePrefix, encoding)
// 	default:
// 		Client = NewRedisCache(redis.Client, cachePrefix, encoding)
// 	}
//
// }

const (
	// DefaultExpireTime 默认过期时间
	DefaultExpireTime = 60 * time.Second
	// PrefixCacheKey 业务cache key
	PrefixCacheKey = "gtodo"
)
