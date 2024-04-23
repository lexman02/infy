package routes

import (
	"infy/controllers"
	"infy/middleware"

	"github.com/gin-gonic/gin"
)

// CommentRoutes sets up routes for comment-related actions.
func CommentRoutes(r *gin.Engine) {
	comment := r.Group("/comments")
	{
		comment.POST("/:commentId/like", controllers.LikeComment)                  // Likes a comment
		comment.POST("/:commentId/dislike", controllers.DislikeComment)            // Dislikes a comment
		comment.POST("/", middleware.Authorized(), controllers.CreateComment)      // Creates a new comment
		comment.PUT("/:id", middleware.Authorized(), controllers.UpdateComment)    // Updates an existing comment
		comment.DELETE("/:id", middleware.Authorized(), controllers.DeleteComment) // Deletes an existing comment
	}
}
