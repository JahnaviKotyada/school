package main

import (
	"log"
	"school/controllers"
	"school/database"
	"school/repositories"
	"school/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database connection
	DB := database.InitDB()
	defer func() {
		db, err := DB.DB()
		if err != nil {
			log.Fatal(err)
		}
		db.Close()
	}()

	// Initialize repositories
	schoolRepo := repositories.NewSchoolRepository(DB)

	// Initialize services
	schoolService := services.NewSchoolService(schoolRepo)

	// Initialize controllers
	schoolController := controllers.NewSchoolController(schoolService)

	// Setup Gin router
	router := gin.Default()
	schoolController.SetupRoutes(router)

	// Run the Gin server
	router.Run(":8080")
}
