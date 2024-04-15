package controllers

import (
	"infy/api"
	"infy/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SearchMovies(c *gin.Context) {
	query := c.Query("title") // Get the movie title from query parameters
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Movie title is required"})
		return
	}

	results, err := api.SearchMovies(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search movies"})
		return
	}

	c.JSON(http.StatusOK, results)
}

// GetTrendingMovies handles requests for fetching trending movies.
func GetTrendingMovies(c *gin.Context) {
	timeWindow := c.Param("timeWindow") // Get timeWindow from URL parameter.

	// Fetch trending movies from TMDb API.
	trendingMovies, err := api.GetTrendingMovies(timeWindow)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch trending movies"})
		return
	}

	// Respond with the fetched trending movies.
	c.JSON(http.StatusOK, trendingMovies)
}

func GetMovieDetails(c *gin.Context) {
	movieID := c.Param("movieID") // Get the movie ID from the URL parameter

	movieDetails, _, err := api.GetMovieDetails(movieID, false)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get movie details"})
		return
	}

	c.JSON(http.StatusOK, movieDetails)
}

func GetMovieCast(c *gin.Context) {
	movieID := c.Param("movieID")
	cast, err := api.GetMovieCast(movieID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch movie cast"})
		return
	}
	c.JSON(http.StatusOK, cast)
}

func GetMovieReviews(c *gin.Context) {
	movieID := c.Param("movieID")
	reviews, err := api.GetMovieReviews(movieID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch movie reviews"})
		return
	}
	c.JSON(http.StatusOK, reviews)
}

func GetSimilarMovies(c *gin.Context) {
	movieID := c.Param("movieID")
	similarMovies, err := api.GetSimilarMovies(movieID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch similar movies"})
		return
	}
	c.JSON(http.StatusOK, similarMovies)
}

func GetFollowedUsersWhoWatchedMovie(c *gin.Context) {
	user, _ := c.Get("user")
	userID := user.(*models.User).ID.Hex()
	movieID := c.Param("movieID")

	users, err := models.FindFollowedWhoWatchedMovie(userID, movieID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		return
	}

	c.JSON(http.StatusOK, users)
}
