package controllers

import (
	"infy/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPollsByMovieID(c *gin.Context) {
	movieID := c.Param("movieID") // Extracting movieID from the URL parameter

	polls, err := models.FindPollsByMovieID(movieID, c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve polls for the movie"})
		return
	}

	c.JSON(http.StatusOK, polls)
}

func CreatePoll(c *gin.Context) {
	var newPoll struct {
		Question string   `json:"question"`
		Options  []string `json:"options"`
	}

	if err := c.ShouldBindJSON(&newPoll); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	movieID := c.Param("movieID") // Extracting movieID from the URL parameter

	poll := models.NewPoll(newPoll.Question, movieID)

	for _, option := range newPoll.Options {
		if option == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Option text cannot be empty"})
			return
		}

		poll.AddOption(option)
	}

	err := poll.Save(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create the poll"})
		return
	}

	c.JSON(http.StatusOK, poll)
}

func AddPollVote(c *gin.Context) {
	pollID := c.Param("pollID")
	var voteData struct {
		OptionID string `json:"optionID"`
	}
	if err := c.ShouldBindJSON(&voteData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err := models.IncrementPollOptionVote(pollID, voteData.OptionID, c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add vote to the poll option"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Vote added successfully"})
}
