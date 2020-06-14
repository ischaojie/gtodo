/*redis Driver接口的redis cache实现*/
package cache

import (
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"reflect"
	"strings"
	"time"
)

const (
	DefaultExp = 60 * time.Second
)

type redisCache struct {
	client     *redis.Client
	keyPrefix  string
	encoding   Encoding
	DefaultExp time.Duration
	newObject  func() interface{}
}

// BuildCacheKey 构建一个带有前缀的缓存key
func BuildCacheKey(keyPrefix string, key string) (cacheKey string, err error) {
	if key == "" {
		return "", errors.New("[cache] key should not be empty")
	}

	cacheKey, err = strings.Join([]string{keyPrefix, key}, ":"), nil
	return
}

/*实现Driver接口*/

// Set 存入缓存
func (c *redisCache) Set(key string, val interface{}, exp time.Duration) error {
	// 对val编码
	buf, err := Marshal(c.encoding, val)
	if err != nil {
		return errors.Wrapf(err, "marshal data err, value is %+v", val)
	}
	// 对key加前缀
	cacheKey, err := BuildCacheKey(c.keyPrefix, key)
	if err != nil {
		return errors.Wrapf(err, "build cache key err, key is %+v", key)
	}
	// 过期时间
	if exp == 0 {
		exp = DefaultExp
	}
	// 存入cache
	err = c.client.Set(cacheKey, buf, exp).Err()
	if err != nil {
		return errors.Wrapf(err, "redis set error")
	}
	return nil
}

// Get 获取缓存
func (c *redisCache) Get(key string, val interface{}) error {
	// 对key加前缀
	cacheKey, err := BuildCacheKey(c.keyPrefix, key)
	if err != nil {
		return errors.Wrapf(err, "build cache key err, key is %+v", key)
	}
	// 从cache中获取缓存数据
	data, err := c.client.Get(cacheKey).Bytes()
	if err != nil {
		return errors.Wrapf(err, "get data from redis error, key: %+v", cacheKey)
	}

	if string(cacheKey) == "" {
		return nil
	}
	// 	解码数据
	err = Unmarshal(c.encoding, data, val)
	if err != nil {
		return errors.Wrapf(err, "unmarshal data error, key=%s, cacheKey=%s type=%v, json is %+v ",
			key, cacheKey, reflect.TypeOf(val), string(data))
	}
	return nil
}

func (redisCache) MultiSet(ValMap map[string]interface{}, exp time.Duration) error {
	panic("implement me")
}

func (redisCache) MultiGet(keys []string, val interface{}) error {
	panic("implement me")
}

// Del 从缓存中删除数据
func (c *redisCache) Del(keys ...string) error {
	if len(keys) == 0 {
		return nil
	}
	cacheKeys := make([]string, len(keys))
	for i, key := range keys {
		cacheKey, err := BuildCacheKey(c.keyPrefix, key)
		if err != nil {
			return errors.Wrapf(err, "build cache key err, key is %+v", key)
		}
		cacheKeys[i] = cacheKey
	}
	err := c.client.Del(cacheKeys...).Err()
	if err != nil {
		return errors.Wrapf(err, "redis delete data error, keys is %+v", keys)
	}
	return nil
}

func (redisCache) Incr(key string, step int64) (int64, error) {
	panic("implement me")
}

func (redisCache) Decr(key string, step int64) (int64, error) {
	panic("implement me")
}

func NewRedisCache(client *redis.Client, keyPrefix string, encoding Encoding, newObject func() interface{}) Driver {
	return &redisCache{
		client:    client,
		keyPrefix: keyPrefix,
		encoding:  encoding,
		newObject: newObject,
	}
}
