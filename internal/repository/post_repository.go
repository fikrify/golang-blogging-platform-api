package repository

import (
	"golang-blogging-platform-api/database"
	"golang-blogging-platform-api/internal/model"
)

type PostRepository struct{}

func NewPostRepository() *PostRepository {
	return &PostRepository{}
}

func (r *PostRepository) FindAll(searchTerm string) ([]model.Post, error) {
	var posts []model.Post

	query := database.DB.Model(&model.Post{})

	if searchTerm != "" {
		query.Where("title LIKE ?", "%"+searchTerm+"%")
	}

	result := database.DB.Find(&posts)

	return posts, result.Error
}
