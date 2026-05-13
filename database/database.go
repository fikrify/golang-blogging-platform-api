package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// ConnectDatabase initializes a connection to the SQLite database using GORM and assigns it to the global `DB` variable.
// It logs a fatal error and
func ConnectDatabase() {
	db, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	DB = db
}
