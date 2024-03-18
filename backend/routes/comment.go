package routes

import (
	"infy/controllers"
	"infy/middleware"

	"github.com/gin-gonic/gin"
)

func CommentRoutes(r *gin.Engine) {
	comment := r.Group("/comments")
	{
		comment.POST("/", middleware.Authorized(), controllers.CreateComment)
		comment.PUT("/:id", middleware.Authorized(), controllers.UpdateComment)
		comment.DELETE("/:id", middleware.Authorized(), controllers.DeleteComment)
	}
}
