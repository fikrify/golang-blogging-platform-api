package routes

import (
	"golang-blogging-platform-api/internal/handler"
	"golang-blogging-platform-api/internal/repository"
	"golang-blogging-platform-api/internal/service"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	postRepo := repository.NewPostRepository()
	postService := service.NewPostService(postRepo)
	postHandler := handler.NewPostHandler(postService)

	r.GET("/posts", postHandler.GetPosts)
}
