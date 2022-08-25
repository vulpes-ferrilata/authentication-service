package infrastructure

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"github.com/vulpes-ferrilata/authentication-service/infrastructure/config"
)

func NewConfig() (config.Config, error) {
	config := config.Config{}

	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	viper.BindEnv("SERVER_ADDRESS")
	viper.BindEnv("DATABASE_ADDRESS")
	viper.BindEnv("DATABASE_NAME")
	viper.BindEnv("DATABASE_USERNAME")
	viper.BindEnv("DATABASE_PASSWORD")
	viper.BindEnv("REDIS_ADDRESS")
	viper.BindEnv("REDIS_PASSWORD")
	viper.BindEnv("ACCESS_TOKEN_ALGORITHM")
	viper.BindEnv("ACCESS_TOKEN_SECRET_KEY")
	viper.BindEnv("ACCESS_TOKEN_EXPIRATION")
	viper.BindEnv("REFRESH_TOKEN_ALGORITHM")
	viper.BindEnv("REFRESH_TOKEN_SECRET_KEY")
	viper.BindEnv("REFRESH_TOKEN_EXPIRATION")

	if err := viper.ReadInConfig(); err != nil {
		return config, errors.WithStack(err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, errors.WithStack(err)
	}

	return config, nil
}
