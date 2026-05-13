package routes

import (
	"golang-blogging-platform-api/internal/handler"
	"golang-blogging-platform-api/internal/repository"
	"golang-blogging-platform-api/internal/service"

	"github.com/gin-gonic/gin"
)

// SetupRoutes defines all the HTTP routes for post-related operations and associates them with their respective handlers.
func SetupRoutes(r *gin.Engine) {
	postRepo := repository.NewPostRepository()
	postService := service.NewPostService(postRepo)
	postHandler := handler.NewPostHandler(postService)

	postRoutes := r.Group("/posts")
	{
		postRoutes.GET("/", postHandler.GetPosts)
		postRoutes.POST("/", postHandler.CreatePost)
		postRoutes.GET("/:id", postHandler.GetPost)
		postRoutes.PUT("/:id", postHandler.UpdatePost)
		postRoutes.DELETE("/:id", postHandler.DeletePost)
	}
}
