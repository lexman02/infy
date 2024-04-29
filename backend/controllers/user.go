package controllers

import (
	"infy/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
