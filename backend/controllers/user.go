package controllers

import (
	"infy/api"
	"infy/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddMovieToWatched adds a movie to the user's watched list
func AddMovieToWatched(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userID := user.(*models.User).ID.Hex()

	var requestBody struct {
		MovieID string `json:"movieId"`
	}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}

	isValid, err := api.IsValidMovieID(requestBody.MovieID) // First declaration of err
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error validating movie ID"})
		return
	}

	if !isValid {
		// The movie ID is not valid
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		return
	}

	err = models.AddMovieToWatchedList(userID, requestBody.MovieID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not add movie to watched list"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Movie added to watched list"})
}

func AddMovieToWatchlist(c *gin.Context) {
	user, _ := c.Get("user")
	userID := user.(*models.User).ID.Hex() // Convert ObjectID to string

	var requestBody struct {
		MovieID string `json:"movieId"`
	}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Validate the movie ID with TMDB API before adding to watchlist
	isValid, err := api.IsValidMovieID(requestBody.MovieID)
	if err != nil {
		// Handle error: Maybe the TMDb API is down or there's a network issue
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error validating movie ID"})
		return
	}

	if !isValid {
		// The movie ID is not valid
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		return
	}

	// Proceed to add the movie to watchlist if the movie ID is valid
	err = models.AddMovieToWatchlist(userID, requestBody.MovieID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not add movie to watchlist"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Movie added to watchlist"})
}

// RemoveMovieFromWatched removes a movie from the user's watched list
func RemoveMovieFromWatched(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userID := user.(*models.User).ID.Hex() // Converts ObjectID to string

	movieID := c.Param("id") // Assuming the movie ID is passed as a URL parameter

	err := models.RemoveMovieFromWatchedList(userID, movieID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not remove movie from watched list"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Movie removed from watched list"})
}

func RemoveMovieFromWatchlist(c *gin.Context) {

	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}
	userID := user.(*models.User).ID.Hex() // Converting ObjectID to string

	movieID := c.Param("id") // Getting the movie ID from the URL parameter

	err := models.RemoveMovieFromWatchlist(userID, movieID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove movie from watchlist"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Movie removed from watchlist successfully"})
}

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
	movieID := c.Param("id") // Get the movie ID from the URL parameter

	movieDetails, err := api.GetMovieDetails(movieID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get movie details"})
		return
	}

	c.JSON(http.StatusOK, movieDetails)
}

func GetMovieCast(c *gin.Context) {
	movieID := c.Param("id")
	cast, err := api.GetMovieCast(movieID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch movie cast"})
		return
	}
	c.JSON(http.StatusOK, cast)
}

func GetMovieReviews(c *gin.Context) {
	movieID := c.Param("id")
	reviews, err := api.GetMovieReviews(movieID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch movie reviews"})
		return
	}
	c.JSON(http.StatusOK, reviews)
}

func GetSimilarMovies(c *gin.Context) {
	movieID := c.Param("id")
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
