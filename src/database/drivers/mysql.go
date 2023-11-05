package drivers

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectMySQL() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("database.username"),
		viper.GetString("database.password"),
		viper.GetString("database.host"),
		viper.GetString("database.port"),
		viper.GetString("database.name"),
	)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})

}
