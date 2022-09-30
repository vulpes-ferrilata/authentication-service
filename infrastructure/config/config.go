package config

type Config struct {
	Server       ServerConfig `mapstructure:"server"`
	Mongo        MongoConfig  `mapstructure:"mongo"`
	Redis        RedisConfig  `mapstructure:"redis"`
	AccessToken  TokenConfig  `mapstructure:"access_token"`
	RefreshToken TokenConfig  `mapstructure:"refresh_token"`
}
