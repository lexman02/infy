package controllers

import (
	"infy/api"
	"infy/middleware"
	"infy/models"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetRecommendations returns a list of recommendations for the user
func GetRecommendationsFromWatched(c *gin.Context) {
	// Get the token from the cookie
	token, err := c.Cookie("token")
	if err != nil {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	// Get the user from the token
	user, err := middleware.GetUserFromToken(token, c)
	if err != nil {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	// Get the user's watched list
	watched := user.Profile.Preferences.Watched
	//Get last movie add to watched list
	lastWatched := watched[len(watched)-1]

	//Use api similar movies by id to get recommendations
	recommendations, nil := api.GetSimilarMovies(lastWatched)

	// if nil, return error
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		log.Println(err)
		return
	}

	// return 10 recommendations
	c.JSON(200, recommendations)
}

func GetRecommendationsFromWatchList(c *gin.Context) {
	// Get the token from the cookie
	token, err := c.Cookie("token")
	if err != nil {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	// Get the user from the token
	user, err := middleware.GetUserFromToken(token, c)
	if err != nil {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	// Get the user's watchlist
	watchList := user.Profile.Preferences.WatchList
	//Get last movie add to watched list
	lastWatched := watchList[len(watchList)-1]

	//Use api similar movies by id to get recommendations
	recommendations, nil := api.GetSimilarMovies(lastWatched)

	// if nil, return error
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		log.Println(err)
		return
	}

	// return 10 recommendations
	c.JSON(200, recommendations)
}

func GetRecommendationsFromFollowing(c *gin.Context) {
	// Get the token from the cookie
	token, err := c.Cookie("token")
	if err != nil {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	// Get the user from the token
	user, err := middleware.GetUserFromToken(token, c)
	if err != nil {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	// Get the user's followers list
	following := user.Profile.Preferences.Following
	// Create a map to store recommendations
	recommendations := make(map[string]*api.TMDbSimilarMoviesResponse)
	finalRecommendations := []string{}

	// Loop through the user's following list for at most 5 users
	for i, userID := range following {
		if i > 5 {
			break
		}

		// Get the following user's profile
		followingProfile, err := models.FindUserByID(userID.Hex(), c.Request.Context())
		if err != nil {
			c.JSON(500, gin.H{"error": "Internal Server Error"})
			log.Println(err)
			return
		}

		// Get the user's watched list
		watchedList := followingProfile.Profile.Preferences.Watched
		if len(watchedList) == 0 || watchedList == nil {
			continue
		}
		// Get last movie from watched list
		lastWatched := watchedList[len(watchedList)-1]

		// Get recommendations from last movie
		followingRecommendations, err := api.GetSimilarMovies(lastWatched)
		if err != nil {
			c.JSON(500, gin.H{"error": "Internal Server Error"})
			log.Println(err)
			return
		}
		// Store the recommendations in the map
		recommendations[userID.Hex()] = followingRecommendations

		// Loop through the recommendations and add them to the final recommendations list
		recommendationCount := 0
		for _, recommendation := range recommendations[userID.Hex()].Results {
			recommendationID := strconv.Itoa(recommendation.ID)
			// Check if recommendation is already in finalRecommendations
			if _, exists := recommendations[recommendationID]; !exists {
				finalRecommendations = append(finalRecommendations, recommendationID)
				recommendationCount++
			}

			// Break if we have 2 recommendations and the following list is longer than 5
			if recommendationCount == 2 && len(following) > 5 {
				break
			}
		}
	}

	c.JSON(200, finalRecommendations)
}

func GetRecommendationsFromFollowers(c *gin.Context) {
	// Get the token from the cookie
	token, err := c.Cookie("token")
	if err != nil {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	// Get the user from the token
	user, err := middleware.GetUserFromToken(token, c)
	if err != nil {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	// Get the user's followers list
	followers := user.Profile.Preferences.Followers
	// Create a map to store recommendations
	recommendations := make(map[string]*api.TMDbSimilarMoviesResponse)
	finalRecommendations := []string{}

	// Loop through the user's followers list for at most 5 users
	for i, userID := range followers {
		if i > 5 {
			break
		}

		// Get the follower's profile
		followerProfile, err := models.FindUserByID(userID.Hex(), c.Request.Context())
		if err != nil {
			c.JSON(500, gin.H{"error": "Internal Server Error"})
			log.Println(err)
			return
		}

		// Get the user's watched list
		watchedList := followerProfile.Profile.Preferences.Watched
		if len(watchedList) == 0 || watchedList == nil {
			continue
		}
		// Get last movie from watched list
		lastWatched := watchedList[len(watchedList)-1]

		// Get recommendations from last movie
		followerRecommendations, err := api.GetSimilarMovies(lastWatched)
		if err != nil {
			c.JSON(500, gin.H{"error": "Internal Server Error"})
			log.Println(err)
			return
		}
		// Store the recommendations in the map
		recommendations[userID.Hex()] = followerRecommendations

		// Loop through the recommendations and add them to the final recommendations list
		recommendationCount := 0
		for _, recommendation := range recommendations[userID.Hex()].Results {
			recommendationID := strconv.Itoa(recommendation.ID)
			// Check if recommendation is already in finalRecommendations
			if _, exists := recommendations[recommendationID]; !exists {
				finalRecommendations = append(finalRecommendations, recommendationID)
				recommendationCount++
			}

			// Break if we have 2 recommendations and the followers list is longer than 5
			if recommendationCount == 2 && len(followers) > 5 {
				break
			}
		}
	}

	c.JSON(200, finalRecommendations)
}
