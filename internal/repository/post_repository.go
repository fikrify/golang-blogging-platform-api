package repository

import (
	"context"
	"golang-blogging-platform-api/internal/model"

	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db: db}
}

// FindAll retrieves a list of posts from the repository, optionally filtered by a search term in the title.
func (r *PostRepository) FindAll(ctx context.Context, searchTerm string) ([]model.Post, error) {
	var posts []model.Post

	query := r.db.
		WithContext(ctx).
		Model(&model.Post{})

	if searchTerm != "" {
		query.Where("title LIKE ?", "%"+searchTerm+"%")
	}

	err := query.Find(&posts).Error

	return posts, err
}

// FindByID retrieves a single post from the database by its unique ID. Returns the post and an error if the operation fails.
func (r *PostRepository) FindByID(ctx context.Context, id uint) (model.Post, error) {
	var post model.Post

	err := r.db.
		WithContext(ctx).
		First(&post, id).
		Error

	return post, err
}

// Create inserts a new post into the database and returns an error if the operation fails.
func (r *PostRepository) Create(ctx context.Context, post *model.Post) error {
	return r.db.
		WithContext(ctx).
		Create(&post).
		Error
}

// Update modifies an existing post in the repository based on the provided post-data and returns an error if it fails.
func (r *PostRepository) Update(ctx context.Context, post *model.Post) error {
	return r.db.
		WithContext(ctx).
		Model(&model.Post{}).
		Where("id = ?", post.ID).
		Updates(post).
		Error
}

// Delete removes a post identified by its unique ID from the database and returns an error if the operation fails.
func (r *PostRepository) Delete(ctx context.Context, id uint) error {
	return r.db.
		WithContext(ctx).
		Delete(&model.Post{}, id).
		Error
}
