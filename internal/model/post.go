package model

import "time"

type Post struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title"  binding:"required"`
	Content   string    `json:"content"  binding:"required"`
	Category  string    `json:"category"  binding:"required"`
	Tags      []string  `json:"tags" gorm:"serializer:json"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
