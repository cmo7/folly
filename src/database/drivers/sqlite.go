package drivers

import (
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectSQLite() (*gorm.DB, error) {
	var file string = viper.GetString("database.sqlite.file")
	return gorm.Open(sqlite.Open(file), &gorm.Config{})
}
