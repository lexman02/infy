package controllers

import (
	"infy/api"
	"infy/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUserProfile retrieves and returns the profile of the currently authenticated user.
func GetUserProfile(c *gin.Context) {
	user := c.MustGet("user").(*models.User)
	if user == nil {
		c.JSON(401, gin.H{"error": "Unauthorized"}) // Responds if no user is found in the context.
		return
	}

	c.JSON(200, gin.H{"profile": user.Profile}) // Successfully returns the user's profile.
}

// GetProfile fetches and returns a user profile based on a username provided as a URL parameter.
func GetProfile(c *gin.Context) {
	profile, err := models.FindUserProfileByUsername(c.Param("username"), c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{"error": "An error occurred"}) // Error handling if the profile cannot be retrieved.
		log.Println(err)
		return
	}

	c.JSON(200, gin.H{"profile": profile}) // Returns the fetched profile.
}

// Follow allows the authenticated user to follow another user specified by their user ID.
func Follow(c *gin.Context) {
	user := c.MustGet("user").(*models.User)
	if user == nil {
		c.JSON(401, gin.H{"error": "Unauthorized"}) // Checks for authenticated user.
		return
	}

	err := user.FollowUser(c.Param("id"), c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{"error": "An error occurred"}) // Handles failure in the follow operation.
		log.Println(err)
		return
	}

	c.JSON(200, gin.H{"success": "User followed"}) // Success response.
}

// Unfollow allows the authenticated user to unfollow another user specified by their user ID.
func Unfollow(c *gin.Context) {
	user := c.MustGet("user").(*models.User)
	if user == nil {
		c.JSON(401, gin.H{"error": "Unauthorized"}) // Checks for authenticated user.
		return
	}

	err := user.UnfollowUser(c.Param("id"), c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{"error": "An error occurred"}) // Handles failure in the unfollow operation.
		log.Println(err)
		return
	}

	c.JSON(200, gin.H{"success": "User unfollowed"}) // Success response.
}

// AddMovieToWatched adds a specified movie to the authenticated user's watched list.
func AddMovieToWatched(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"}) // Checks if the user is authenticated.
		return
	}

	userID := user.(*models.User).ID.Hex() // Extracts userID from the user context.

	var requestBody struct {
		MovieID string `json:"movieId"`
	}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()}) // Validates the JSON body.
		log.Println(err)
		return
	}

	isValid, err := api.IsValidMovieID(requestBody.MovieID) // Validates the movie ID against an external API.
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error validating movie ID"}) // Handles API errors.
		log.Println(err)
		return
	}

	if !isValid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"}) // Handles invalid movie ID.
		log.Println(err)
		return
	}

	err = models.AddMovieToWatchedList(userID, requestBody.MovieID, c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not add movie to watched list"}) // Handles failure in adding to watched list.
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Movie added to watched list"}) // Success response.
}

// AddMovieToWatchlist adds a specified movie to the authenticated user's watchlist.
func AddMovieToWatchlist(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"}) // Checks if the user is authenticated.
		return
	}

	userID := user.(*models.User).ID.Hex() // Extracts userID from the user context.

	var requestBody struct {
		MovieID string `json:"movieId"`
	}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"}) // Validates the JSON body.
		log.Println(err)
		return
	}

	isValid, err := api.IsValidMovieID(requestBody.MovieID) // Validates the movie ID against an external API.
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error validating movie ID"}) // Handles API errors.
		log.Println(err)
		return
	}

	if !isValid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"}) // Handles invalid movie ID.
		log.Println(err)
		return
	}

	err = models.AddMovieToWatchlist(userID, requestBody.MovieID, c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not add movie to watchlist"}) // Handles failure in adding to watchlist.
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Movie added to watchlist"}) // Success response.
}

// RemoveMovieFromWatched removes a specified movie from the authenticated user's watched list.
func RemoveMovieFromWatched(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"}) // Checks if the user is authenticated.
		return
	}

	userID := user.(*models.User).ID.Hex() // Converts ObjectID to string.
	movieID := c.Param("id")               // Assumes the movie ID is passed as a URL parameter.

	err := models.RemoveMovieFromWatchedList(userID, movieID, c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not remove movie from watched list"}) // Handles failure in removing from watched list.
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Movie removed from watched list"}) // Success response.
}

// RemoveMovieFromWatchlist removes a specified movie from the authenticated user's watchlist.
func RemoveMovieFromWatchlist(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"}) // Checks if the user is authenticated.
		return
	}
	userID := user.(*models.User).ID.Hex() // Converts ObjectID to string.
	movieID := c.Param("id")               // Gets the movie ID from the URL parameter.

	err := models.RemoveMovieFromWatchlist(userID, movieID, c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove movie from watchlist"}) // Handles failure in removing from watchlist.
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Movie removed from watchlist successfully"}) // Success response.
}

// AddUserAvatar uploads and sets a user avatar, validating the image type and size.
func AddUserAvatar(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"}) // Checks if the user is authenticated.
		return
	}

	file, err := c.FormFile("avatar")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"}) // Checks if the file was uploaded.
		log.Println(err)
		return
	}

	// Validate file type and size.
	if file.Header.Get("Content-Type") != "image/jpeg" && file.Header.Get("Content-Type") != "image/png" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file type. Only JPEG and PNG images are allowed"})
		return
	}

	if file.Size > 5<<20 { // 5 MB limit
		c.JSON(http.StatusBadRequest, gin.H{"error": "File size should not exceed 5MB"})
		return
	}

	// Set the filename to the user's ID and save the file using the original filetype.
	file.Filename = user.(*models.User).ID.Hex() + file.Filename[len(file.Filename)-4:]

	// Save the file to the uploads/avatars directory.
	filepath := "uploads/avatars/" + file.Filename
	err = c.SaveUploadedFile(file, filepath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save file"})
		log.Println(err)
		return
	}

	// Add avatar to the user's profile.
	err = models.AddAvatar(user.(*models.User), file.Filename, c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not add avatar"})
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Avatar added successfully"})
}
