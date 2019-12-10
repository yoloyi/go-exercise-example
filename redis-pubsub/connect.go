package example

import (
	"github.com/go-redis/redis"
)

const (
	RedisAddr = ":6379"
)

// InitRedis 初始化 Redis
func InitRedis() *redis.Client {
	option := &redis.Options{
		Addr: RedisAddr,
	}

	client := redis.NewClient(option)

	return client
}
