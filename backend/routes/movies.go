package routes

import (
	"infy/controllers"
	"infy/middleware"

	"github.com/gin-gonic/gin"
)

func MovieRoutes(r *gin.Engine) {
	movies := r.Group("/movies")
	{
		// Movie detail routes
		movies.GET("/:movieID", controllers.GetMovieDetails)
		movies.GET("/:movieID/cast", controllers.GetMovieCast)
		movies.GET("/:movieID/reviews", controllers.GetMovieReviews)
		movies.GET("/:movieID/similar", controllers.GetSimilarMovies)

		// Poll routes
		movies.GET("/:movieID/polls", controllers.GetPollsByMovieID)
		movies.POST("/:movieID/polls", middleware.Authorized(), controllers.CreatePoll)
		movies.POST("/:movieID/polls/:pollID/vote", middleware.Authorized(), controllers.AddPollVote)

		// Misc. movie routes
		movies.GET("/search", controllers.SearchMovies)
		movies.GET("/trending/:timeWindow", controllers.GetTrendingMovies)
	}
}
