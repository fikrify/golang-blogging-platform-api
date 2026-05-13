package service

import (
	"golang-blogging-platform-api/internal/model"
	"golang-blogging-platform-api/internal/repository"
)

type PostService struct {
	repo *repository.PostRepository
}

func NewPostService(repo *repository.PostRepository) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) GetPosts() ([]model.Post, error) {
	return s.repo.FindAll()
}
