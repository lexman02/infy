package routes

import (
	"infy/controllers"
	"infy/middleware"

	"github.com/gin-gonic/gin"
)

// ProfileRoutes sets up routes for user profiles and related functionalities.
func ProfileRoutes(r *gin.Engine) {
	profile := r.Group("/profile")
	{
		userProfile := profile.Group("/user")
		userProfile.Use(middleware.Authorized())
		{
			userProfile.GET("/", controllers.GetUserProfile)       // Retrieves the logged-in user's profile
			userProfile.POST("/avatar", controllers.AddUserAvatar) // Adds an avatar to the user's profile
		}

		profile.GET("/:username", controllers.GetProfile) // Retrieves a user's profile by username

		profile.POST("/follow/:id", middleware.Authorized(), controllers.Follow)       // Follows another user
		profile.DELETE("/unfollow/:id", middleware.Authorized(), controllers.Unfollow) // Unfollows another user

		movies := profile.Group("/movies")
		movies.Use(middleware.Authorized())
		{
			movies.POST("/add/watched", controllers.AddMovieToWatched)                             // Adds a movie to the user's watched list
			movies.POST("/add/watchlist", controllers.AddMovieToWatchlist)                         // Adds a movie to the user's watchlist
			movies.DELETE("/watched/:id", controllers.RemoveMovieFromWatched)                      // Removes a movie from the watched list
			movies.DELETE("/watchlist/:id", controllers.RemoveMovieFromWatchlist)                  // Removes a movie from the watchlist
			movies.GET("/:movieID/watchedByFollowed", controllers.GetFollowedUsersWhoWatchedMovie) // Gets followed users who watched a specific movie
			movies.GET("/watched/recommendations", controllers.GetRecommendationsFromWatched)
			movies.GET("/watchlist/recommendations", controllers.GetRecommendationsFromWatchList)
			movies.GET("/following/watched", controllers.GetRecommendationsFromFollowing)
			movies.GET("/followers/watched", controllers.GetRecommendationsFromFollowers)
		}
	}
}
