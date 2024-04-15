package controllers

import (
	"infy/api"
	"infy/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserProfile(c *gin.Context) {
	user := c.MustGet("user").(*models.User)
	if user == nil {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	c.JSON(200, gin.H{"profile": user.Profile})
}

func GetProfile(c *gin.Context) {
	profile, err := models.FindUserProfileByUsername(c.Param("username"), c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{"error": "An error occurred"})
		log.Println(err)
		return
	}

	c.JSON(200, gin.H{"profile": profile})
}

func Follow(c *gin.Context) {
	user := c.MustGet("user").(*models.User)
	if user == nil {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	err := user.FollowUser(c.Param("id"), c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{"error": "An error occurred"})
		log.Println(err)
		return
	}

	c.JSON(200, gin.H{"success": "User followed"})
}

func Unfollow(c *gin.Context) {
	user := c.MustGet("user").(*models.User)
	if user == nil {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	err := user.UnfollowUser(c.Param("id"), c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{"error": "An error occurred"})
		log.Println(err)
		return
	}

	c.JSON(200, gin.H{"success": "User unfollowed"})
}

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
		log.Println(err)
		return
	}

	isValid, err := api.IsValidMovieID(requestBody.MovieID) // First declaration of err
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error validating movie ID"})
		log.Println(err)
		return
	}

	if !isValid {
		// The movie ID is not valid
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		log.Println(err)
		return
	}

	err = models.AddMovieToWatchedList(userID, requestBody.MovieID, c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not add movie to watched list"})
		log.Println(err)
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
		log.Println(err)
		return
	}

	// Validate the movie ID with TMDB API before adding to watchlist
	isValid, err := api.IsValidMovieID(requestBody.MovieID)
	if err != nil {
		// Handle error: Maybe the TMDb API is down or there's a network issue
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error validating movie ID"})
		log.Println(err)
		return
	}

	if !isValid {
		// The movie ID is not valid
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		log.Println(err)
		return
	}

	// Proceed to add the movie to watchlist if the movie ID is valid
	err = models.AddMovieToWatchlist(userID, requestBody.MovieID, c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not add movie to watchlist"})
		log.Println(err)
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

	err := models.RemoveMovieFromWatchedList(userID, movieID, c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not remove movie from watched list"})
		log.Println(err)
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

	err := models.RemoveMovieFromWatchlist(userID, movieID, c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove movie from watchlist"})
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Movie removed from watchlist successfully"})
}
