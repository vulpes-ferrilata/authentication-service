package config

type Config struct {
	Server       ServerConfig   `mapstructure:"server"`
	Database     DatabaseConfig `mapstructure:"database"`
	Redis        RedisConfig    `mapstructure:"redis"`
	AccessToken  TokenConfig    `mapstructure:"access_token"`
	RefreshToken TokenConfig    `mapstructure:"refresh_token"`
}
