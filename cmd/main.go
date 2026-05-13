package main

import (
	"golang-blogging-platform-api/database"
	"golang-blogging-platform-api/internal/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDatabase()

	database.RunMigrations()

	router := gin.Default()

	routes.SetupRoutes(router)

	if err := router.Run(); err != nil {
		log.Fatal(err)
	}
}
