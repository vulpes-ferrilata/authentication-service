package infrastructure

import "github.com/go-redis/redis/v8"

func NewRedis(config *Config) (*redis.Client, error) {
	addr := config.Redis.Host + ":" + config.Redis.Port
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: config.Redis.Password,
		DB:       0,
	})
	return rdb, nil
}
