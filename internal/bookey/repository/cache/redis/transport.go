package redis

import (
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

type Redis struct {
	Client *redis.Client
	Cache  *cache.Cache
}

type RedisTransport interface {
}

func NewRedis(client *redis.Client, cache *cache.Cache) RedisTransport {
	return &Redis{
		Client: client,
		Cache:  cache,
	}
}
