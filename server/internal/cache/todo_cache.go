package cache

import (
	"fmt"
	"github.com/shiniao/gtodo/internal/model"
	"github.com/shiniao/gtodo/pkg/cache"
	"github.com/shiniao/gtodo/pkg/redis"
)

const (
	// PrefixUserCacheKey cache前缀
	PrefixTodoCacheKey = "todo:cache:%d"
)

func NewTodoCache() *Cache {
	encoding := cache.JSONEncoding{}
	return &Cache{cache: cache.NewRedisCache(redis.Client, cache.PrefixCacheKey, encoding, func() interface{} {
		return &model.TodoModel{}
	})}
}

// GetTodoCache 获取todo cache
func (c *Cache) GetTodoCache(todoID uint64) (todoModel *model.TodoModel, err error) {
	cacheKey := fmt.Sprintf(PrefixTodoCacheKey, todoID)
	err = c.cache.Get(cacheKey, &todoModel)
	if err != nil {
		return todoModel, err
	}
	return todoModel, nil
}

// SetTodoCache 设置todo cache
func (c *Cache) SetTodoCache(todoID uint64, todo *model.TodoModel) error {
	if todo == nil || todo.ID == 0 {
		return nil
	}
	cacheKey := fmt.Sprintf(PrefixTodoCacheKey, todoID)
	err := c.cache.Set(cacheKey, todo, DefaultExpireTime)
	if err != nil {
		return err
	}
	return nil
}

// DelTodoCache 删除todo cache
func (c *Cache) DelTodoCache(todoID uint64) error {
	cacheKey := fmt.Sprintf(PrefixTodoCacheKey, todoID)
	err := c.cache.Del(cacheKey)
	if err != nil {
		return err
	}
	return nil
}
