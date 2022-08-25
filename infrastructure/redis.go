package infrastructure

import (
	"github.com/go-redis/redis/v8"
	"github.com/vulpes-ferrilata/authentication-service/infrastructure/config"
)

func NewRedis(config config.Config) (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.Redis.Address,
		Password: config.Redis.Password,
		DB:       0,
	})

	return redisClient, nil
}
