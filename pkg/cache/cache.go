package cache

import (
	redisPkg "core-api/pkg/redis"
	"time"

	"github.com/redis/go-redis/v9"
)

func init() {
	NewCache()
}

type Cache interface {
	Get(key string) (string, error)
	Set(key string, value string, expiration time.Duration) error
	Remember(key string, expiration time.Duration, getter func() (string, error)) (string, error)
}

type CacheImpl struct {
	redisDB *redis.Client
}

func NewCache() Cache {
	return &CacheImpl{
		redisDB: redisPkg.Rdb,
	}
}

func (c *CacheImpl) Get(key string) (string, error) {
	val, err := c.redisDB.Get(redisPkg.Ctx, key).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		return "", err
	}
	return val, nil
}

func (c *CacheImpl) Set(key string, value string, expiration time.Duration) error {
	err := c.redisDB.Set(redisPkg.Ctx, key, value, expiration).Err()
	return err
}

func (c *CacheImpl) Remember(key string, expiration time.Duration, getter func() (string, error)) (string, error) {
	cachedValue, err := c.Get(key)
	if err == nil {
		return cachedValue, nil
	}

	fetchedValue, err := getter()
	if err != nil {
		return "", err
	}

	err = c.Set(key, fetchedValue, expiration)
	if err != nil {
		return "", err
	}

	return fetchedValue, nil
}
