package routes

import (
	"golang-blogging-platform-api/internal/handler"

	"github.com/gin-gonic/gin"
)

// RegisterPostRoutes registers routes for managing posts, including CRUD operations, under the "/posts" route group.
func RegisterPostRoutes(rg *gin.RouterGroup, postHandler *handler.PostHandler) {
	rg.GET("/", postHandler.GetPosts)
	rg.POST("/", postHandler.CreatePost)
	rg.GET("/:id", postHandler.GetPost)
	rg.PUT("/:id", postHandler.UpdatePost)
	rg.DELETE("/:id", postHandler.DeletePost)
}
