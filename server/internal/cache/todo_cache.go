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

func (c *Cache) GetTodoCache(todoID uint64) (todoModel *model.TodoModel, err error) {
	cacheKey := fmt.Sprintf(PrefixTodoCacheKey, todoID)
	err = c.cache.Get(cacheKey, &todoModel)
	if err != nil {
		return todoModel, err
	}
	return todoModel, nil
}
