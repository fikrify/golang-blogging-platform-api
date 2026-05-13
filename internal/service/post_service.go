package service

import (
	"golang-blogging-platform-api/internal/model"
	"golang-blogging-platform-api/internal/repository"
)

type PostService struct {
	repo *repository.PostRepository
}

// NewPostService initializes a new instance of PostService with the given PostRepository.
func NewPostService(repo *repository.PostRepository) *PostService {
	return &PostService{repo: repo}
}

// GetPosts retrieves a list of posts filtered by an optional search term, returning the matching posts or an error.
func (s *PostService) GetPosts(searchTerm string) ([]model.Post, error) {
	return s.repo.FindAll(searchTerm)
}

// GetPost retrieves a post by its unique ID and returns the post or an error if the operation fails.
func (s *PostService) GetPost(id uint) (model.Post, error) {
	return s.repo.FindByID(id)
}

// CreatePost creates a new post in the repository and returns an error if the operation fails.
func (s *PostService) CreatePost(post *model.Post) error {
	return s.repo.Create(post)
}

// UpdatePost updates an existing post in the repository with new data and returns an error if the operation fails.
func (s *PostService) UpdatePost(post *model.Post) error {
	return s.repo.Update(post)
}

// DeletePost deletes a post by its unique ID and returns an error if the deletion fails.
func (s *PostService) DeletePost(id uint) error {
	return s.repo.Delete(id)
}
