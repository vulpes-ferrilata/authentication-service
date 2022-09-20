package infrastructure

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/authentication-service/infrastructure/config"
)

func NewRedis(config config.Config) (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.Redis.Address,
		Password: config.Redis.Password,
		DB:       0,
	})

	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		return nil, errors.WithStack(err)
	}

	return redisClient, nil
}
