package dao

import (
	"sky-take-out-go/config"

	"github.com/go-redis/redis/v8"
)

var Redis *redis.Client

func InitRedis(config config.RedisConfig) {
	Redis = redis.NewClient(&redis.Options{
		Addr:     config.Addr(),
		Password: config.Password,
		DB:       config.DB,
	})
}
