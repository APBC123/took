package InitRedis

import (
	"github.com/go-redis/redis/v8"
	"took/server/service/core/internal/config"
)

func InitRedis(c config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     c.CacheRedis.Addr,
		Password: "",
		DB:       0, // use default DB
	})

}
