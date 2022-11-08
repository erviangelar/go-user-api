package helper

import (
	redis "github.com/go-redis/redis/v8"

	"github.com/erviangelar/go-user-api/common/cache"
	"github.com/erviangelar/go-user-api/common/config"
)

func New(conf *config.Configurations) (*cache.RedisClient, error) {
	if len(conf.Cache.Host) == 0 {
		return nil, nil
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.Cache.Host,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	redisCache := cache.New(rdb)

	return redisCache, nil
}
