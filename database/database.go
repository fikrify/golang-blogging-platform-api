package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Init establishes a connection to the SQLite database and returns a configured *gorm.DB instance.
func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	return db
}
