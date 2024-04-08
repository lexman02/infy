package routes

import (
	"infy/controllers"
	"infy/middleware"

	"github.com/gin-gonic/gin"
)

func PostRoutes(r *gin.Engine) {
	post := r.Group("/posts")
	{
		post.GET("/", controllers.GetPosts)
		post.GET("/:id", controllers.GetPost)
		post.POST("/", middleware.Authorized(), controllers.CreatePost)
		post.PUT("/:id", middleware.Authorized(), controllers.UpdatePost)
		post.DELETE("/:id", middleware.Authorized(), controllers.DeletePost)

		// Like a post
		post.POST("/:id/like", middleware.Authorized(), controllers.LikePost)
		// Dislike a post
		post.POST("/:id/dislike", middleware.Authorized(), controllers.DislikePost)
		post.GET("/user/:userID", controllers.GetUserPosts)
		post.GET("/movie/:movieID", controllers.GetPostsByMovieID)

	}

	// Adding Polls routes under /movies to follow the structure given in the assignment
	movies := r.Group("/movies")
	{
		movies.GET("/:movieID", controllers.GetMovieDetails) // Adjusted for consistency
		movies.GET("/:movieID/polls", controllers.GetPollsByMovieID)
		movies.POST("/:movieID/polls", middleware.Authorized(), controllers.CreatePoll)
		movies.POST("/:movieID/polls/:pollID/vote", middleware.Authorized(), controllers.AddPollVote)
	}
}
