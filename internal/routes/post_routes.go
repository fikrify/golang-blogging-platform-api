package routes

import (
	"golang-blogging-platform-api/internal/handler"

	"github.com/gin-gonic/gin"
)

// RegisterPostRoutes registers routes for managing posts, including CRUD operations, under the "/posts" route group.
func RegisterPostRoutes(rg *gin.RouterGroup, postHandler *handler.PostHandler) {
	postRoutes := rg.Group("/posts")
	{
		postRoutes.GET("/", postHandler.GetPosts)
		postRoutes.POST("/", postHandler.CreatePost)
		postRoutes.GET("/:id", postHandler.GetPost)
		postRoutes.PUT("/:id", postHandler.UpdatePost)
		postRoutes.DELETE("/:id", postHandler.DeletePost)
	}
}
