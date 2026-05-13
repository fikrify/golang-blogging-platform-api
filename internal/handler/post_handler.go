package handler

import (
	"golang-blogging-platform-api/internal/service"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	service *service.PostService
}

func NewPostHandler(service *service.PostService) *PostHandler {
	return &PostHandler{service: service}
}

func (h *PostHandler) GetPosts(c *gin.Context) {
	posts, err := h.service.GetPosts()

	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, posts)
}
