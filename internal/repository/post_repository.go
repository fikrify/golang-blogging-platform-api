package repository

import (
	"golang-blogging-platform-api/database"
	"golang-blogging-platform-api/internal/model"
)

type PostRepository struct{}

func NewPostRepository() *PostRepository {
	return &PostRepository{}
}

func (r *PostRepository) FindAll() ([]model.Post, error) {
	var posts []model.Post

	result := database.DB.Find(&posts)

	return posts, result.Error
}
