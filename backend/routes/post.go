package routes

import (
	"infy/controllers"
	"infy/middleware"

	"github.com/gin-gonic/gin"
)

// PostRoutes sets up routes related to posts.
func PostRoutes(r *gin.Engine) {
	post := r.Group("/posts")
	{
		post.GET("/", controllers.GetPosts)                                  // Retrieves all posts
		post.GET("/:id", controllers.GetPost)                                // Retrieves a specific post
		post.POST("/", middleware.Authorized(), controllers.CreatePost)      // Creates a new post
		post.PUT("/:id", middleware.Authorized(), controllers.UpdatePost)    // Updates an existing post
		post.DELETE("/:id", middleware.Authorized(), controllers.DeletePost) // Deletes an existing post

		post.POST("/:id/like", middleware.Authorized(), controllers.LikePost)       // Likes a post
		post.POST("/:id/dislike", middleware.Authorized(), controllers.DislikePost) // Dislikes a post
		post.GET("/:id/report", middleware.Authorized(), controllers.ReportPost)    // Reports a post

		post.GET("/user/:userID", controllers.GetUserPosts)        // Retrieves posts by a specific user
		post.GET("/movie/:movieID", controllers.GetPostsByMovieID) // Retrieves posts related to a specific movie
	}
}
