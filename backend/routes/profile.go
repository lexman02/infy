package routes

import (
	"infy/controllers"
	"infy/middleware"

	"github.com/gin-gonic/gin"
)

func ProfileRoutes(r *gin.Engine) {
	profile := r.Group("/profile")
	{
		userProfile := profile.Group("/user")
		userProfile.Use(middleware.Authorized())
		{
			userProfile.GET("/", controllers.GetUserProfile)
			//userProfile.PUT("/", controllers.UpdateUserProfile)
		}

		profile.GET("/:username", controllers.GetProfile)

		// Routes for following and unfollowing users
		profile.POST("/follow/:id", middleware.Authorized(), controllers.Follow)
		profile.DELETE("/unfollow/:id", middleware.Authorized(), controllers.Unfollow)

		// Routes for movies functionality within the profile route group
		movies := profile.Group("/movies")
		movies.Use(middleware.Authorized())
		{
			movies.POST("/add/watched", controllers.AddMovieToWatched)
			movies.POST("/add/watchlist", controllers.AddMovieToWatchlist)
			movies.DELETE("/watched/:id", controllers.RemoveMovieFromWatched)
			movies.DELETE("/watchlist/:id", controllers.RemoveMovieFromWatchlist)
			movies.GET("/:movieID/watchedByFollowed", controllers.GetFollowedUsersWhoWatchedMovie)
		}

	}
}
