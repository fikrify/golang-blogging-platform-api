package database

import (
	"golang-blogging-platform-api/internal/model"
	"log"
)

func RunMigrations() {
	err := DB.AutoMigrate(&model.Post{})
	if err != nil {
		log.Fatal(err)
	}
}
