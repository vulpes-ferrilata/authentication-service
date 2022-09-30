package config

type MongoConfig struct {
	Address      string `mapstructure:"address"`
	DatabaseName string `mapstructure:"database_name"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
}
