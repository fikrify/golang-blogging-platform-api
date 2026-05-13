package repository

import (
	"golang-blogging-platform-api/database"
	"golang-blogging-platform-api/internal/model"
)

type PostRepository struct{}

func NewPostRepository() *PostRepository {
	return &PostRepository{}
}

// FindAll retrieves a list of posts from the repository, optionally filtered by a search term in the title.
func (r *PostRepository) FindAll(searchTerm string) ([]model.Post, error) {
	var posts []model.Post

	query := database.DB.Model(&model.Post{})

	if searchTerm != "" {
		query.Where("title LIKE ?", "%"+searchTerm+"%")
	}

	result := database.DB.Find(&posts)

	return posts, result.Error
}

// FindByID retrieves a single post from the database by its unique ID. Returns the post and an error if the operation fails.
func (r *PostRepository) FindByID(id uint) (model.Post, error) {
	var post model.Post

	result := database.DB.First(&post, id)

	return post, result.Error
}

// Create inserts a new post into the database and returns an error if the operation fails.
func (r *PostRepository) Create(post *model.Post) error {
	result := database.DB.Create(&post)

	return result.Error
}

// Update modifies an existing post in the repository based on the provided post data and returns an error if it fails.
func (r *PostRepository) Update(post *model.Post) error {
	result := database.DB.
		Model(&model.Post{}).
		Where("id = ?", post.ID).
		Updates(post)

	return result.Error
}

// Delete removes a post identified by its unique ID from the database and returns an error if the operation fails.
func (r *PostRepository) Delete(id uint) error {
	result := database.DB.Delete(&model.Post{}, id)

	return result.Error
}
