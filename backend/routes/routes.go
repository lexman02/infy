package routes

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// InitRoutes initializes all the route groups and settings for the application.
func InitRoutes() *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "Cache-Control"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Serve avatar images from the specified directory
	router.Static("/avatars", "./uploads/avatars")

	// Register the route groups
	AuthRoutes(router)
	PostRoutes(router)
	ProfileRoutes(router)
	CommentRoutes(router)
	MovieRoutes(router)
	AdminRoutes(router)

	return router
}
