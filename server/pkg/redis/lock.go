/*lock 访问缓存市加锁*/

/*
* 为什么需要加锁？防止缓存击穿
* 缓存击穿：大量请求同时查询某个key，但是这个key过期了，导致对数据库的大量请求
* 解决办法：利用分布式锁，只有拿到锁的第一个线程请求数据库，然后缓存
*
 */

package redis

import (
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"strings"
	"time"
)

// Lock
type Lock struct {
	key     string        // redis key
	client  *redis.Client // redis client
	timeout time.Duration // 锁有效期
}

// NewLock 初始化lock
func NewLock(conn *redis.Client, key string, defaultTimeout time.Duration) *Lock {
	return &Lock{
		key:     key,
		client:  conn,
		timeout: defaultTimeout,
	}
}

// Lock 加锁
func (l *Lock) Lock(token string) (bool, error) {
	// SET key value [exp] EX 当key不存在时，才操作
	ok, err := l.client.SetNX(l.GetKey(), token, l.timeout).Result()
	if err == redis.Nil {
		err = nil
	}
	return ok, err
}

// UnLock 解锁
func (l *Lock) UnLock(token string) error {
	// token一致才删除
	script := "if redis.call('get',KEYS[1]) == ARGV[1] then return redis.call('del',KEYS[1]) else return 0 end"
	_, err := l.client.Eval(script, []string{l.GetKey(), token}).Result()
	if err != nil {
		return err
	}
	return nil
}

func (l *Lock) GetKey() string {
	keyPrefix := viper.GetString("name")
	return strings.Join([]string{keyPrefix, "redis:lock"}, "")
}

func (l *Lock) GenToken() string {
	u, _ := uuid.NewRandom()
	return u.String()
}
