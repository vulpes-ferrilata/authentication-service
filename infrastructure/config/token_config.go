package config

type TokenConfig struct {
	Algorithm  string `mapstructure:"algorithm"`
	SecretKey  string `mapstructure:"secret_key"`
	Expiration string `mapstructure:"expiration"`
}
