package redis

import (
	"context"
	"fmt"
	"time"

	interfaces "course_seckill_clean_architecture/interface"

	"github.com/redis/go-redis/v9"
)

// In this project, redisCache is the only implementation of Cache interface
type redisCache struct {
	cl *redis.Client
}

func NewInstance(host string, port string, pass string, db int) (interfaces.Cache, error) {
	rc := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", host, port),
		Password: pass,
		DB: db,
	})

	return &redisCache{cl: rc}, nil
}

func (c *redisCache) HSet(ctx context.Context, hashName string, key string, value interface{}) error {
	return c.cl.HSet(ctx, hashName, key, value).Err()
}

func (c *redisCache) HGet(ctx context.Context, hashName string, key string) (interface{}, error) {
	if key == "all" {
		return c.cl.HGetAll(ctx, hashName).Result()
	}
	return c.cl.HGet(ctx, hashName, key).Result()
}

func (c *redisCache) HIncrBy(ctx context.Context, hashName string, key string, incr int64) (int64, error) {
	return c.cl.HIncrBy(ctx, hashName, key, incr).Result()
}

func (c *redisCache) SAdd(ctx context.Context, setName string, values ...interface{}) error {
	return c.cl.SAdd(ctx, setName, values...).Err()
}

func (c *redisCache) SIsMember(ctx context.Context, setName string, value interface{}) (bool, error) {
	return c.cl.SIsMember(ctx, setName, value).Result()
}

func (c *redisCache) Expire(ctx context.Context, object string, expiration time.Duration) error {
	return c.cl.Expire(ctx, object, expiration).Err()
}

func (c *redisCache) RunScript(ctx context.Context, script string, keys []string, args ...interface{}) (interface{}, error) {
	s := redis.NewScript(script)
	return s.Run(ctx, c.cl, keys, args...).Result()
}

func (c *redisCache) Del(ctx context.Context, object string) error {
	return c.cl.Del(ctx, object).Err()
}

func (rc *redisCache) Close() error {
	return rc.cl.Close()
}


