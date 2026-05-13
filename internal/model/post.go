package model

import (
	"time"
)

type Post struct {
	ID        uint `gorm:"primaryKey"`
	Title     string
	Content   string
	Category  string
	Tags      []string `gorm:"serializer:json"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
