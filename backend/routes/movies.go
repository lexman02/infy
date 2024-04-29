package routes

import (
	"infy/controllers"
	"infy/middleware"

	"github.com/gin-gonic/gin"
)

// MovieRoutes sets up routes related to movies.
func MovieRoutes(r *gin.Engine) {
	movies := r.Group("/movies")
	{
		movies.GET("/:movieID", controllers.GetMovieDetails)          // Retrieves details for a specific movie
		movies.GET("/:movieID/cast", controllers.GetMovieCast)        // Retrieves cast information for a specific movie
		movies.GET("/:movieID/reviews", controllers.GetMovieReviews)  // Retrieves reviews for a specific movie
		movies.GET("/:movieID/similar", controllers.GetSimilarMovies) // Retrieves movies similar to a specific movie

		movies.GET("/:movieID/polls", controllers.GetPollsByMovieID)                                  // Retrieves polls related to a specific movie
		movies.POST("/:movieID/polls", middleware.Authorized(), controllers.CreatePoll)               // Creates a poll related to a specific movie
		movies.POST("/:movieID/polls/:pollID/vote", middleware.Authorized(), controllers.AddPollVote) // Adds a vote to a specific poll

		movies.GET("/search", controllers.SearchMovies)                    // Searches for movies based on a query
		movies.GET("/trending/:timeWindow", controllers.GetTrendingMovies) // Retrieves trending movies within a specified time window

		// New actor details route + movie credits route + movie actorID finder + Movie Trailers
		movies.GET("/actor/:actorID", controllers.GetActorDetails)
		movies.GET("/actor/:actorID/movies", controllers.GetActorMovieCredits)
		movies.GET("/:movieID/trailers", controllers.GetMovieTrailers)
	}

	people := r.Group("/people")
	{
		people.GET("/search", controllers.SearchPeople) // Retrieves details for a specific person
	}
}
