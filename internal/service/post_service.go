package service

import (
	"context"
	"golang-blogging-platform-api/internal/model"
)

type PostRepository interface {
	FindAll(ctx context.Context, searchTerm string) ([]model.Post, error)
	FindByID(ctx context.Context, id uint) (model.Post, error)
	Create(ctx context.Context, post *model.Post) error
	Update(ctx context.Context, post *model.Post) error
	Delete(ctx context.Context, id uint) error
}

type PostService struct {
	repo PostRepository
}

// NewPostService initializes a new instance of PostService with the given PostRepository.
func NewPostService(repo PostRepository) *PostService {
	return &PostService{repo: repo}
}

// GetPosts retrieves a list of posts filtered by an optional search term, returning the matching posts or an error.
func (s *PostService) GetPosts(ctx context.Context, searchTerm string) ([]model.Post, error) {
	return s.repo.FindAll(ctx, searchTerm)
}

// GetPost retrieves a post by its unique ID and returns the post or an error if the operation fails.
func (s *PostService) GetPost(ctx context.Context, id uint) (model.Post, error) {
	return s.repo.FindByID(ctx, id)
}

// CreatePost creates a new post in the repository and returns an error if the operation fails.
func (s *PostService) CreatePost(ctx context.Context, post *model.Post) error {
	return s.repo.Create(ctx, post)
}

// UpdatePost updates an existing post in the repository with new data and returns an error if the operation fails.
func (s *PostService) UpdatePost(ctx context.Context, post *model.Post) error {
	return s.repo.Update(ctx, post)
}

// DeletePost deletes a post by its unique ID and returns an error if the deletion fails.
func (s *PostService) DeletePost(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}
