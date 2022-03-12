package database

import (
	"fmt"

	"github.com/dev-hyunsang/my-own-diary/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectionDB() (*gorm.DB, error) {
	// user:password@tcp(location)/dbname
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		config.GetEnv("MYSQL_ACCOUNT"),
		config.GetEnv("MYSQL_PASSWORD"),
		config.GetEnv("MYSQL_LOCATION"),
		config.GetEnv("MYSQL_DATABASE"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return db, err
}
