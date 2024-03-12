package routes

import (
	"infy/controllers"
	"infy/middleware"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

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

	// Initialize the routes
	AuthRoutes(router)
	PostRoutes(router)

	// Setup routes for movies functionality within the profile route group
	profile := router.Group("/profile")
	profile.Use(middleware.Authorized()) // Use the Authorized middleware
	router.GET("/search/movies", controllers.SearchMovies)
	profile.POST("/add/watched", controllers.AddMovieToWatched)
	profile.POST("/add/watchlist", controllers.AddMovieToWatchlist)
	profile.DELETE("/watched/:id", controllers.RemoveMovieFromWatched)
	profile.DELETE("/watchlist/:id", controllers.RemoveMovieFromWatchlist)

	return router
}
