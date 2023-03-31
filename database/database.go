package database

import (
	"fmt"
	"grpc-tennis/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		config.App.Config.DatabaseHost,
		config.App.Config.DatabaseUser,
		config.App.Config.DatabasePassword,
		config.App.Config.DatabaseName,
		config.App.Config.DatabasePort,
		config.App.Config.DatabaseSslMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	DB = db
	return db
}
