package database

import (
	"folly/src/database/drivers"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	switch viper.GetString("database.driver") {
	case "mysql":
		db, err = drivers.ConnectMySQL()
	case "sqlite":
		db, err = drivers.ConnectSQLite()
	case "postgres":
		db, err = drivers.ConnectPostgres()
	default:
		panic("Invalid database driver")
	}
	DB = db

	return db, err
}
