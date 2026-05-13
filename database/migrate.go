package database

import (
	"golang-blogging-platform-api/internal/model"
	"log"
)

// RunMigrations performs database migrations for the Post model using the GORM AutoMigrate function.
func RunMigrations() {
	err := DB.AutoMigrate(&model.Post{})
	if err != nil {
		log.Fatal(err)
	}
}
