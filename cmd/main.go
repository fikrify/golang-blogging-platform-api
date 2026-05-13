package main

import (
	"golang-blogging-platform-api/database"
	"golang-blogging-platform-api/internal/handler"
	"golang-blogging-platform-api/internal/repository"
	"golang-blogging-platform-api/internal/routes"
	"golang-blogging-platform-api/internal/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.Init()
	database.RunMigrations(db)

	postRepo := repository.NewPostRepository(db)
	postService := service.NewPostService(postRepo)
	postHandler := handler.NewPostHandler(postService)

	router := gin.Default()

	postRouter := router.Group("/posts")

	routes.RegisterPostRoutes(postRouter, postHandler)

	if err := router.Run(); err != nil {
		log.Fatal(err)
	}
}
