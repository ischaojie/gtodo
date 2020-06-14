/*memory： Driver接口的memory cache 实现*/
package cache

import (
	"sync"
	"time"
)

type memoryCache struct {
	client    *sync.Map
	keyPrefix string
	encoding  Encoding
}

func (c *memoryCache) Set(key string, val interface{}, exp time.Duration) error {

}

func (c *memoryCache) Get(key string, val interface{}) error {
	panic("implement me")
}

func (c *memoryCache) MultiSet(ValMap map[string]interface{}, exp time.Duration) error {
	panic("implement me")
}

func (c *memoryCache) MultiGet(keys []string, val interface{}) error {
	panic("implement me")
}

func (c *memoryCache) Del(keys ...string) error {
	panic("implement me")
}

func (c *memoryCache) Incr(key string, step int64) (int64, error) {
	panic("implement me")
}

func (c *memoryCache) Decr(key string, step int64) (int64, error) {
	panic("implement me")
}


func NewMemoryCache(keyPrefix string, encoding Encoding) Driver {
	return &memoryCache{
		client:    &sync.Map{},
		keyPrefix: keyPrefix,
		encoding:  encoding,
	}
}
