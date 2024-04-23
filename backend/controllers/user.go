package controllers

import (
	"infy/api"
	"infy/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SearchMovies searches for movies based on a title query parameter.
func SearchMovies(c *gin.Context) {
	query := c.Query("title") // Retrieve the movie title from the query parameters.
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Movie title is required"}) // Validate the presence of the title query.
		return
	}

	results, err := api.SearchMovies(query) // Call to external API to search for movies by title.
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search movies"}) // Handle errors from the movie search API.
		return
	}

	c.JSON(http.StatusOK, results) // Return the search results as a JSON response.
}

// GetTrendingMovies fetches trending movies based on a specified time window.
func GetTrendingMovies(c *gin.Context) {
	timeWindow := c.Param("timeWindow") // Extract timeWindow from the URL parameter.

	trendingMovies, err := api.GetTrendingMovies(timeWindow) // Fetch trending movies from the API.
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch trending movies"}) // Handle potential API errors.
		return
	}

	c.JSON(http.StatusOK, trendingMovies) // Respond with the fetched trending movies data.
}

// GetMovieDetails fetches details for a single movie identified by its ID.
func GetMovieDetails(c *gin.Context) {
	movieID := c.Param("movieID") // Extract the movie ID from URL parameters.

	movieDetails, _, err := api.GetMovieDetails(movieID, false) // Fetch movie details from the API.
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get movie details"}) // Handle errors from the movie details fetch.
		return
	}

	c.JSON(http.StatusOK, movieDetails) // Return the movie details as a JSON response.
}

// GetMovieCast retrieves the cast of a specific movie by its ID.
func GetMovieCast(c *gin.Context) {
	movieID := c.Param("movieID") // Extract the movie ID from URL parameters.

	cast, err := api.GetMovieCast(movieID) // Fetch movie cast from the API.
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch movie cast"}) // Handle errors during the fetch.
		return
	}

	c.JSON(http.StatusOK, cast) // Respond with the movie cast.
}

// GetMovieReviews fetches reviews for a specific movie by its ID.
func GetMovieReviews(c *gin.Context) {
	movieID := c.Param("movieID") // Extract the movie ID from URL parameters.

	reviews, err := api.GetMovieReviews(movieID) // Fetch reviews from the API.
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch movie reviews"}) // Handle errors during the fetch.
		return
	}

	c.JSON(http.StatusOK, reviews) // Respond with the movie reviews.
}

// GetSimilarMovies retrieves movies similar to a specified movie by its ID.
func GetSimilarMovies(c *gin.Context) {
	movieID := c.Param("movieID") // Extract the movie ID from URL parameters.

	similarMovies, err := api.GetSimilarMovies(movieID) // Fetch similar movies from the API.
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch similar movies"}) // Handle errors during the fetch.
		return
	}

	c.JSON(http.StatusOK, similarMovies) // Respond with the similar movies.
}

// GetFollowedUsersWhoWatchedMovie fetches a list of followed users who have watched a specified movie.
func GetFollowedUsersWhoWatchedMovie(c *gin.Context) {
	user, _ := c.Get("user")
	userID := user.(*models.User).ID.Hex() // Extract the user ID from the context.
	movieID := c.Param("movieID")          // Extract the movie ID from URL parameters.

	users, err := models.FindFollowedWhoWatchedMovie(userID, movieID) // Query the database for followed users who watched the movie.
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"}) // Handle errors in the database query.
		return
	}

	c.JSON(http.StatusOK, users) // Respond with the list of users.
}
