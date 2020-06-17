/*redis Driver接口的redis cache实现*/
package cache

import (
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"github.com/shiniao/gtodo/pkg/log"
	"reflect"
	"strings"
	"time"
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
		exp = DefaultExpireTime
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

// MultiSet 批量存入缓存
func (c *redisCache) MultiSet(ValMap map[string]interface{}, exp time.Duration) error {
	if len(ValMap) == 0 {
		return nil
	}
	paris := make([]interface{}, 0, 2*len(ValMap))
	for key, value := range ValMap {
		buf, err := Marshal(c.encoding, value)
		if err != nil {
			log.Warnf("marshal data err: %+v, value is %+v", err, value)
			continue
		}
		cacheKey, err := BuildCacheKey(c.keyPrefix, key)
		if err != nil {
			log.Warnf("build cache key err: %+v, key is %+v", err, key)
			continue
		}
		paris = append(paris, []byte(cacheKey))
		paris = append(paris, buf)
	}
	if exp == 0 {
		exp = DefaultExpireTime
	}

	// 存入redis
	err := c.client.MSet(paris...).Err()
	if err != nil {
		return errors.Wrapf(err, "redis multi set error")
	}

	// 设置过期时间
	for i := 0; i < len(paris); i = i + 2 {
		switch paris[i].(type) {
		case []byte:
			c.client.Expire(string(paris[i].([]byte)), exp)
		default:
			log.Warnf("redis expire is unsupported key type:&+v", reflect.TypeOf(paris[i]))
		}

	}
	return nil
}

// MultiGet 批量获取缓存
func (c *redisCache) MultiGet(keys []string, val interface{}) error {
	if len(keys) == 0 {
		return nil
	}
	cacheKeys := make([]string, len(keys))
	for i, key := range keys {
		cacheKey, err := BuildCacheKey(c.keyPrefix, key)
		if err != nil {
			return errors.Wrapf(err, "build cache key error: %+v", err)

		}
		cacheKeys[i] = cacheKey
	}
	values, err := c.client.MGet(cacheKeys...).Result()
	if err != nil {
		return errors.Wrapf(err, "redis MGet error, keys is %+v", keys)
	}

	// 通过反射注入到map
	valueMap := reflect.ValueOf(val)
	for i, value := range values {
		if value == nil {
			continue
		}
		object := c.newObject()
		err = Unmarshal(c.encoding, []byte(value.(string)), object)
		if err != nil {
			log.Warnf("unmarshal data error: %+v, key=%s, cacheKey=%s type=%v", err,
				keys[i], cacheKeys[i], reflect.TypeOf(value))
			continue
		}
		valueMap.SetMapIndex(reflect.ValueOf(cacheKeys[i]), reflect.ValueOf(object))
	}
	return nil

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

func (c *redisCache) Incr(key string, step int64) (int64, error) {
	cacheKey, err := BuildCacheKey(c.keyPrefix, key)
	if err != nil {
		return 0, errors.Wrapf(err, "build cache key err, key is %+v", key)
	}

	affectRow, err := c.client.Incr(cacheKey).Result()
	if err != nil {
		return 0, errors.Wrapf(err, "redis incr, keys is %+v", key)
	}
	return affectRow, nil
}

func (c *redisCache) Decr(key string, step int64) (int64, error) {
	cacheKey, err := BuildCacheKey(c.keyPrefix, key)
	if err != nil {
		return 0, errors.Wrapf(err, "build cache key err, key is %+v", key)
	}
	affectRow, err := c.client.DecrBy(cacheKey, step).Result()
	if err != nil {
		return 0, errors.Wrapf(err, "redis incr, keys is %+v", key)
	}
	return affectRow, nil
}

// NewRedisCache 新的redis连接
func NewRedisCache(client *redis.Client, keyPrefix string, encoding Encoding, newObject func() interface{}) Driver {
	return &redisCache{
		client:    client,
		keyPrefix: keyPrefix,
		encoding:  encoding,
		newObject: newObject,
	}
}
