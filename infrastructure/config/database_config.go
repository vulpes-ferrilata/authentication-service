package config

type DatabaseConfig struct {
	Address  string `mapstructure:"address"`
	Name     string `mapstructure:"name"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}
