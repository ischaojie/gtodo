package cache

import (
	"github.com/shiniao/gtodo/pkg/cache"
	"time"
)

const (
	// DefaultExpireTime 默认过期时间
	DefaultExpireTime = time.Hour * 24
)

type Cache struct {
	cache cache.Driver
}
