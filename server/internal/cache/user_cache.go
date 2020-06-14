package cache

import (
	"fmt"
	"github.com/shiniao/gtodo/internal/model"
	"github.com/shiniao/gtodo/pkg/cache"
	"github.com/shiniao/gtodo/pkg/redis"
	"time"
)

const (
	// PrefixUserCacheKey cache前缀
	PrefixUserCacheKey = "user:cache:%d"
	// DefaultExpireTime 默认过期时间
	DefaultExpireTime = time.Hour * 24
)

type Cache struct {
	cache cache.Driver
}

// NewUserCache 选择不同的cache端初始化user cache
func NewUserCache() *Cache {
	encoding := cache.JSONEncoding{}
	cachePrefix := cache.PrefixCacheKey
	return &Cache{cache: cache.NewRedisCache(redis.Client, cachePrefix, encoding, func() interface{} {
		return &model.UserModel{}
	})}
}

// SetUserCache 将用户存入cache
func (c *Cache) SetUserCache(userID uint64, user *model.UserModel) error {
	if user == nil || user.ID == 0 {
		return nil
	}
	cacheKey := fmt.Sprintf(PrefixUserCacheKey, userID)
	err := c.cache.Set(cacheKey, user, DefaultExpireTime)
	if err != nil {
		return err
	}
	return nil

}

// GetUserCache 获取用户缓存
func (c *Cache) GetUserCache(userID uint64) (userModel *model.UserModel, err error) {
	cacheKey := fmt.Sprintf(PrefixUserCacheKey, userID)
	err = c.cache.Get(cacheKey, &userModel)
	if err != nil {
		return userModel, err
	}
	return userModel, nil
}

// DelUserCache 删除用户缓存
func (c *Cache) DelUserCache(userID uint64) error {
	cacheKey := fmt.Sprintf(PrefixUserCacheKey, userID)
	err := c.cache.Del(cacheKey)
	if err != nil {
		return err
	}
	return nil
}
