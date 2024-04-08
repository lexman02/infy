package routes

import (
	"infy/controllers"
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
	ProfileRoutes(router)
	CommentRoutes(router)

	router.GET("/search/movies", controllers.SearchMovies)
	router.GET("/trending/:timeWindow", controllers.GetTrendingMovies)

	router.GET("/movies/:movieID/cast", controllers.GetMovieCast)
	router.GET("/movies/:movieID/reviews", controllers.GetMovieReviews)
	router.GET("/movies/:movieID/similar", controllers.GetSimilarMovies)

	return router
}
