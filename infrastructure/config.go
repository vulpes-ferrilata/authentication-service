package infrastructure

import (
	"os"
	"time"

	"github.com/pkg/errors"
)

var (
	ErrEnvironmentVariableNotSet error = errors.New("environment variable is not set")
)

func NewConfig() (*Config, error) {
	config := new(Config)

	databaseConfig, err := newDatabaseConfig()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	config.Database = databaseConfig

	redisConfig, err := newRedisConfig()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	config.Redis = redisConfig

	authenticationConfig, err := newAuthenticationConfig()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	config.Authentication = authenticationConfig

	return config, nil
}

type Config struct {
	Database       *DatabaseConfig
	Redis          *RedisConfig
	Authentication *AuthenticationConfig
}

func newDatabaseConfig() (*DatabaseConfig, error) {
	databaseConfig := new(DatabaseConfig)

	dbHost, ok := os.LookupEnv("DB_HOST")
	if !ok {
		return nil, errors.Wrap(ErrEnvironmentVariableNotSet, "DB_HOST")
	}
	databaseConfig.Host = dbHost

	dbPort, ok := os.LookupEnv("DB_PORT")
	if !ok {
		return nil, errors.Wrap(ErrEnvironmentVariableNotSet, "DB_PORT")
	}
	databaseConfig.Port = dbPort

	dbUsername, ok := os.LookupEnv("DB_USERNAME")
	if !ok {
		return nil, errors.Wrap(ErrEnvironmentVariableNotSet, "DB_USERNAME")
	}
	databaseConfig.Username = dbUsername

	dbPassword, ok := os.LookupEnv("DB_PASSWORD")
	if !ok {
		return nil, errors.Wrap(ErrEnvironmentVariableNotSet, "DB_PASSWORD")
	}
	databaseConfig.Password = dbPassword

	dbName, ok := os.LookupEnv("DB_NAME")
	if !ok {
		return nil, errors.Wrap(ErrEnvironmentVariableNotSet, "DB_NAME")
	}
	databaseConfig.Name = dbName

	return databaseConfig, nil
}

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Name     string
}

func newRedisConfig() (*RedisConfig, error) {
	redisConfig := new(RedisConfig)

	redisHost, ok := os.LookupEnv("REDIS_HOST")
	if !ok {
		return nil, errors.Wrap(ErrEnvironmentVariableNotSet, "REDIS_HOST")
	}
	redisConfig.Host = redisHost

	redisPort, ok := os.LookupEnv("REDIS_PORT")
	if !ok {
		return nil, errors.Wrap(ErrEnvironmentVariableNotSet, "REDIS_PORT")
	}
	redisConfig.Port = redisPort

	redisPassword, ok := os.LookupEnv("REDIS_PASSWORD")
	if !ok {
		return nil, errors.Wrap(ErrEnvironmentVariableNotSet, "REDIS_PASSWORD")
	}
	redisConfig.Password = redisPassword

	return redisConfig, nil
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
}

func newAuthenticationConfig() (*AuthenticationConfig, error) {
	authenticationConfig := new(AuthenticationConfig)

	accessTokenSecret, ok := os.LookupEnv("ACCESS_TOKEN_SECRET")
	if !ok {
		return nil, errors.Wrap(ErrEnvironmentVariableNotSet, "ACCESS_TOKEN_SECRET")
	}
	authenticationConfig.AccessTokenSecret = accessTokenSecret

	accessTokenDurationStr, ok := os.LookupEnv("ACCESS_TOKEN_DURATION")
	if !ok {
		return nil, errors.Wrap(ErrEnvironmentVariableNotSet, "ACCESS_TOKEN_DURATION")
	}
	accessTokenDuration, err := time.ParseDuration(accessTokenDurationStr)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	authenticationConfig.AccessTokenDuration = accessTokenDuration

	refreshTokenSecret, ok := os.LookupEnv("REFRESH_TOKEN_SECRET")
	if !ok {
		return nil, errors.Wrap(ErrEnvironmentVariableNotSet, "REFRESH_TOKEN_SECRET")
	}
	authenticationConfig.RefreshTokenSecret = refreshTokenSecret

	refreshTokenDurationStr, ok := os.LookupEnv("REFRESH_TOKEN_DURATION")
	if !ok {
		return nil, errors.Wrap(ErrEnvironmentVariableNotSet, "REFRESH_TOKEN_DURATION")
	}
	refreshTokenDuration, err := time.ParseDuration(refreshTokenDurationStr)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	authenticationConfig.RefreshTokenDuration = refreshTokenDuration

	return authenticationConfig, nil
}

type AuthenticationConfig struct {
	AccessTokenSecret    string
	AccessTokenDuration  time.Duration
	RefreshTokenSecret   string
	RefreshTokenDuration time.Duration
}
