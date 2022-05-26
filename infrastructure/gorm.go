package infrastructure

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewGorm(config *Config) (*gorm.DB, error) {
	dsn := config.Database.Username + ":" + config.Database.Password + "@tcp" + "(" + config.Database.Host + ":" + config.Database.Port + ")" + "/" + config.Database.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
	dialector := mysql.Open(dsn)
	return gorm.Open(dialector)
}
