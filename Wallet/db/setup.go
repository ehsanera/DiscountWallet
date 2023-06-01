package db

import (
	"gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Configure GORM logger
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // Use standard logger
		logger.Config{
			LogLevel: logger.Info, // Set log level (you can adjust it to your needs)
		},
	)

	database, err := gorm.Open(sqlite.Open("wallet.db"), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&Wallet{}, &Transaction{})
	if err != nil {
		return
	}

	DB = database
}
