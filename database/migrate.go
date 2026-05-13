package database

import (
	"golang-blogging-platform-api/internal/model"
	"log"

	"gorm.io/gorm"
)

// RunMigrations performs database migrations for the Post model using the GORM AutoMigrate function.
func RunMigrations(db *gorm.DB) {
	err := db.AutoMigrate(&model.Post{})
	if err != nil {
		log.Fatal(err)
	}
}
