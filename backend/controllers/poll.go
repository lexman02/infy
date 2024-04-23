package controllers

import (
	"infy/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetPollsByMovieID retrieves polls related to a specific movie identified by movieID.
func GetPollsByMovieID(c *gin.Context) {
	movieID := c.Param("movieID") // Extracting movieID from the URL parameter

	polls, err := models.FindPollsByMovieID(movieID, c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve polls for the movie"})
		return
	}

	c.JSON(http.StatusOK, polls)
}

// CreatePoll processes the incoming request to create a new poll associated with a movie.
func CreatePoll(c *gin.Context) {
	var newPoll struct {
		Question string   `json:"question"`
		Options  []string `json:"options"`
	}

	// Bind JSON payload to struct and handle errors
	if err := c.ShouldBindJSON(&newPoll); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	movieID := c.Param("movieID") // Extracting movieID from the URL parameter

	// Create a new poll instance
	poll := models.NewPoll(newPoll.Question, movieID)

	// Add options to the poll, ensuring no empty options
	for _, option := range newPoll.Options {
		if option == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Option text cannot be empty"})
			return
		}
		poll.AddOption(option)
	}

	// Save the new poll and handle any errors
	err := poll.Save(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create the poll"})
		return
	}

	c.JSON(http.StatusOK, poll)
}

// AddPollVote increments the vote count for a specific option in a poll.
func AddPollVote(c *gin.Context) {
	pollID := c.Param("pollID") // Extracting pollID from the URL parameter
	var voteData struct {
		OptionID string `json:"optionID"`
	}

	// Bind JSON payload to struct and handle errors
	if err := c.ShouldBindJSON(&voteData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Increment the vote count for the specified poll option
	err := models.IncrementPollOptionVote(pollID, voteData.OptionID, c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add vote to the poll option"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Vote added successfully"})
}
