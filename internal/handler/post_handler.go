package handler

import (
	"golang-blogging-platform-api/internal/model"
	"golang-blogging-platform-api/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	service *service.PostService
}

func NewPostHandler(service *service.PostService) *PostHandler {
	return &PostHandler{service: service}
}

func (h *PostHandler) GetPosts(c *gin.Context) {
	searchTerm := c.Query("term")

	posts, err := h.service.GetPosts(searchTerm)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, posts)
}

// CreatePost handles the HTTP POST request to create a new post with data provided in the request body.
func (h *PostHandler) CreatePost(c *gin.Context) {
	var post model.Post

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := h.service.CreatePost(post)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, post)
}

// GetPost handles the HTTP request to retrieve a single post by its ID from the route parameters.
func (h *PostHandler) GetPost(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid id",
		})
		return
	}

	var post model.Post
	post, err = h.service.GetPost(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, post)
}

// UpdatePost handles the HTTP PUT request to update an existing post's data identified by its ID in the route parameters.
func (h *PostHandler) UpdatePost(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid id",
		})
		return
	}

	var post model.Post

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	post.ID = uint(id)
	err = h.service.UpdatePost(post)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, post)
}

// DeletePost handles the HTTP DELETE request to remove a post identified by its ID from the route parameters.
func (h *PostHandler) DeletePost(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid id",
		})
		return
	}

	err = h.service.DeletePost(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}
